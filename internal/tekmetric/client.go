// Package tekmetric provides a client for the Tekmetric shop management API.
// It handles OAuth2 authentication, rate limiting, and provides methods for
// accessing shops, customers, vehicles, repair orders, and other resources.
package tekmetric

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/beetlebugorg/tekmetric-mcp/internal/config"
	"github.com/beetlebugorg/tekmetric-mcp/pkg/ratelimit"
)

// temporaryError represents a temporary error that should be retried.
// This includes rate limit errors (429) and server errors (5xx).
type temporaryError struct {
	statusCode int
	message    string
}

func (e *temporaryError) Error() string {
	return e.message
}

// Temporary returns true indicating this error is temporary and should be retried.
func (e *temporaryError) Temporary() bool {
	return true
}

// Client is a Tekmetric API client that handles authentication and API requests.
// It manages OAuth2 tokens, implements rate limiting, and provides a clean
// interface to the Tekmetric REST API.
//
// The client automatically:
//   - Obtains and refreshes OAuth2 access tokens
//   - Retries failed requests with exponential backoff
//   - Adds proper authentication headers
//   - Handles JSON encoding/decoding
type Client struct {
	baseURL      string                 // API base URL (sandbox or production)
	clientID     string                 // OAuth2 client ID
	clientSecret string                 // OAuth2 client secret
	httpClient   *http.Client           // HTTP client with timeout
	accessToken  string                 // Current OAuth2 access token
	tokenExpiry  time.Time              // Token expiration time
	shopIDs      []string               // Shop IDs from token scope
	rateLimiter  *ratelimit.RateLimiter // Rate limiter for API requests
	logger       *slog.Logger           // Structured logger
}

// NewClient creates a new Tekmetric API client.
// The client is ready to use but not yet authenticated.
// Call Authenticate() before making API requests.
//
// Parameters:
//   - cfg: Tekmetric API configuration (credentials, base URL, timeouts)
//   - logger: Structured logger for client operations
//
// Returns:
//   - *Client: Configured API client ready for authentication
func NewClient(cfg *config.TekmetricConfig, logger *slog.Logger) *Client {
	return &Client{
		baseURL:      cfg.BaseURL,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.TimeoutSeconds) * time.Second,
		},
		rateLimiter: ratelimit.New(cfg.MaxRetries, cfg.MaxBackoffSec),
		logger:      logger,
	}
}

// Authenticate obtains an access token from the Tekmetric API
func (c *Client) Authenticate(ctx context.Context) error {
	c.logger.Info("authenticating with Tekmetric API")

	// Create Basic Auth header
	auth := base64.StdEncoding.EncodeToString([]byte(c.clientID + ":" + c.clientSecret))

	// Prepare request body
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/api/v1/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("User-Agent", "tekmetric-mcp (https://github.com/beetlebugorg/tekmetric-mcp)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send auth request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("authentication failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	c.accessToken = tokenResp.AccessToken
	c.shopIDs = strings.Fields(tokenResp.Scope)    // Space-separated shop IDs
	c.tokenExpiry = time.Now().Add(24 * time.Hour) // Assume 24h expiry if not specified

	c.logger.Info("authentication successful", "shop_ids", c.shopIDs)
	return nil
}

// ensureAuthenticated checks if we have a valid token and authenticates if needed
func (c *Client) ensureAuthenticated(ctx context.Context) error {
	if c.accessToken == "" || time.Now().After(c.tokenExpiry) {
		return c.Authenticate(ctx)
	}
	return nil
}

// doRequest performs an HTTP request with authentication and rate limiting
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	if err := c.ensureAuthenticated(ctx); err != nil {
		return err
	}

	return c.rateLimiter.Do(func() error {
		var reqBody io.Reader
		if body != nil {
			jsonData, err := json.Marshal(body)
			if err != nil {
				return fmt.Errorf("failed to marshal request body: %w", err)
			}
			reqBody = bytes.NewReader(jsonData)
		}

		req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+c.accessToken)
		req.Header.Set("User-Agent", "tekmetric-mcp (https://github.com/beetlebugorg/tekmetric-mcp)")
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		c.logger.Debug("making API request", "method", method, "path", path)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		// Check for non-success status codes
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			// Rate limit (429) and server errors (5xx) are temporary - should retry
			if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
				return &temporaryError{
					statusCode: resp.StatusCode,
					message:    fmt.Sprintf("temporary error with status %d: %s", resp.StatusCode, string(responseBody)),
				}
			}
			// Client errors (4xx except 429) are permanent - don't retry
			return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(responseBody))
		}

		if result != nil {
			if err := json.Unmarshal(responseBody, result); err != nil {
				return fmt.Errorf("failed to decode response: %w", err)
			}
		}

		return nil
	})
}

// GetShops returns all shops accessible by the current token
func (c *Client) GetShops(ctx context.Context) ([]Shop, error) {
	var shops []Shop
	if err := c.doRequest(ctx, "GET", "/api/v1/shops", nil, &shops); err != nil {
		return nil, err
	}
	return shops, nil
}

// GetShop returns a specific shop by ID
func (c *Client) GetShop(ctx context.Context, id int) (*Shop, error) {
	var shop Shop
	path := fmt.Sprintf("/api/v1/shops/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &shop); err != nil {
		return nil, err
	}
	return &shop, nil
}

// GetCustomers returns a paginated list of customers
func (c *Client) GetCustomers(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Customer], error) {
	path := fmt.Sprintf("/api/v1/customers?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Customer]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCustomer returns a specific customer by ID
func (c *Client) GetCustomer(ctx context.Context, id int) (*Customer, error) {
	var customer Customer
	path := fmt.Sprintf("/api/v1/customers/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetVehicles returns a paginated list of vehicles
func (c *Client) GetVehicles(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Vehicle], error) {
	path := fmt.Sprintf("/api/v1/vehicles?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Vehicle]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVehicle returns a specific vehicle by ID
func (c *Client) GetVehicle(ctx context.Context, id int) (*Vehicle, error) {
	var vehicle Vehicle
	path := fmt.Sprintf("/api/v1/vehicles/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &vehicle); err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// RepairOrderQueryParams holds query parameters for repair order searches
type RepairOrderQueryParams struct {
	Shop                 int    `url:"shop,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Size                 int    `url:"size,omitempty"`
	Start                string `url:"start,omitempty"`            // Date format: YYYY-MM-DD
	End                  string `url:"end,omitempty"`              // Date format: YYYY-MM-DD
	PostedDateStart      string `url:"postedDateStart,omitempty"`  // Date format: YYYY-MM-DD
	PostedDateEnd        string `url:"postedDateEnd,omitempty"`    // Date format: YYYY-MM-DD
	UpdatedDateStart     string `url:"updatedDateStart,omitempty"` // Date format: YYYY-MM-DD
	UpdatedDateEnd       string `url:"updatedDateEnd,omitempty"`   // Date format: YYYY-MM-DD
	DeletedDateStart     string `url:"deletedDateStart,omitempty"` // Date format: YYYY-MM-DD
	DeletedDateEnd       string `url:"deletedDateEnd,omitempty"`   // Date format: YYYY-MM-DD
	RepairOrderNumber    int    `url:"repairOrderNumber,omitempty"`
	RepairOrderStatusIds []int  `url:"repairOrderStatusId,omitempty"` // 1-Estimate, 2-WIP, 3-Complete, 4-Saved, 5-Posted, 6-AR, 7-Deleted
	CustomerID           int    `url:"customerId,omitempty"`
	VehicleID            int    `url:"vehicleId,omitempty"`
	Search               string `url:"search,omitempty"`        // Search by RO#, customer name, vehicle info
	Sort                 string `url:"sort,omitempty"`          // createdDate, repairOrderNumber, customer.firstName, customer.lastName
	SortDirection        string `url:"sortDirection,omitempty"` // ASC, DESC
}

// CustomerQueryParams holds query parameters for customer searches
type CustomerQueryParams struct {
	Shop          int    `url:"shop,omitempty"`
	Page          int    `url:"page,omitempty"`
	Size          int    `url:"size,omitempty"`
	Search        string `url:"search,omitempty"`        // Search by name, email, phone
	Email         string `url:"email,omitempty"`         // Filter by email
	Phone         string `url:"phone,omitempty"`         // Filter by phone
	Sort          string `url:"sort,omitempty"`          // firstName, lastName, email, createdDate
	SortDirection string `url:"sortDirection,omitempty"` // ASC, DESC
}

// VehicleQueryParams holds query parameters for vehicle searches
type VehicleQueryParams struct {
	Shop          int    `url:"shop,omitempty"`
	Page          int    `url:"page,omitempty"`
	Size          int    `url:"size,omitempty"`
	Search        string `url:"search,omitempty"`        // Search by VIN, license plate, make/model
	VIN           string `url:"vin,omitempty"`           // Filter by VIN
	LicensePlate  string `url:"licensePlate,omitempty"`  // Filter by license plate
	Year          int    `url:"year,omitempty"`          // Filter by year
	Make          string `url:"make,omitempty"`          // Filter by make
	Model         string `url:"model,omitempty"`         // Filter by model
	CustomerID    int    `url:"customerId,omitempty"`    // Filter by customer
	Sort          string `url:"sort,omitempty"`          // year, make, model, createdDate
	SortDirection string `url:"sortDirection,omitempty"` // ASC, DESC
}

// AppointmentQueryParams holds query parameters for appointment searches
type AppointmentQueryParams struct {
	Shop          int    `url:"shop,omitempty"`
	Page          int    `url:"page,omitempty"`
	Size          int    `url:"size,omitempty"`
	StartDate     string `url:"startDate,omitempty"`     // Date format: YYYY-MM-DD
	EndDate       string `url:"endDate,omitempty"`       // Date format: YYYY-MM-DD
	CustomerID    int    `url:"customerId,omitempty"`    // Filter by customer
	VehicleID     int    `url:"vehicleId,omitempty"`     // Filter by vehicle
	Status        string `url:"status,omitempty"`        // Filter by status
	Search        string `url:"search,omitempty"`        // Search by customer name, vehicle info
	Sort          string `url:"sort,omitempty"`          // scheduledTime, createdDate, customer.lastName
	SortDirection string `url:"sortDirection,omitempty"` // ASC, DESC
}

// JobQueryParams holds query parameters for job searches
type JobQueryParams struct {
	Shop          int    `url:"shop,omitempty"`
	Page          int    `url:"page,omitempty"`
	Size          int    `url:"size,omitempty"`
	RepairOrderID int    `url:"repairOrderId,omitempty"` // Filter by repair order
	EmployeeID    int    `url:"employeeId,omitempty"`    // Filter by assigned employee
	Status        string `url:"status,omitempty"`        // Filter by status
	Search        string `url:"search,omitempty"`        // Search by name, description
	Sort          string `url:"sort,omitempty"`          // createdDate, name
	SortDirection string `url:"sortDirection,omitempty"` // ASC, DESC
}

// EmployeeQueryParams holds query parameters for employee searches
type EmployeeQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	Search           string `url:"search,omitempty"`           // Search by name, email
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Date format: YYYY-MM-DD
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Date format: YYYY-MM-DD
	Active           *bool  `url:"active,omitempty"`           // Filter by active status
	Role             string `url:"role,omitempty"`             // Filter by role
	Sort             string `url:"sort,omitempty"`             // firstName, lastName, email
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// InventoryQueryParams holds query parameters for inventory searches
type InventoryQueryParams struct {
	Shop          int    `url:"shop,omitempty"`
	Page          int    `url:"page,omitempty"`
	Size          int    `url:"size,omitempty"`
	Search        string `url:"search,omitempty"`        // Search by part name, number
	PartNumber    string `url:"partNumber,omitempty"`    // Filter by part number
	LowStock      *bool  `url:"lowStock,omitempty"`      // Filter by low stock status
	Sort          string `url:"sort,omitempty"`          // name, partNumber, quantity
	SortDirection string `url:"sortDirection,omitempty"` // ASC, DESC
}

// GetRepairOrders returns a paginated list of repair orders
func (c *Client) GetRepairOrders(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[RepairOrder], error) {
	params := RepairOrderQueryParams{
		Shop: shopID,
		Page: page,
		Size: size,
	}
	return c.GetRepairOrdersWithParams(ctx, params)
}

// GetRepairOrdersWithParams returns repair orders with advanced filtering
func (c *Client) GetRepairOrdersWithParams(ctx context.Context, params RepairOrderQueryParams) (*PaginatedResponse[RepairOrder], error) {
	// Build query string
	query := url.Values{}
	if params.Shop > 0 {
		query.Add("shop", fmt.Sprintf("%d", params.Shop))
	}
	query.Add("page", fmt.Sprintf("%d", params.Page))
	if params.Size > 0 {
		query.Add("size", fmt.Sprintf("%d", params.Size))
	} else {
		query.Add("size", "100")
	}
	if params.Start != "" {
		query.Add("start", params.Start)
	}
	if params.End != "" {
		query.Add("end", params.End)
	}
	if params.PostedDateStart != "" {
		query.Add("postedDateStart", params.PostedDateStart)
	}
	if params.PostedDateEnd != "" {
		query.Add("postedDateEnd", params.PostedDateEnd)
	}
	if params.UpdatedDateStart != "" {
		query.Add("updatedDateStart", params.UpdatedDateStart)
	}
	if params.UpdatedDateEnd != "" {
		query.Add("updatedDateEnd", params.UpdatedDateEnd)
	}
	if params.DeletedDateStart != "" {
		query.Add("deletedDateStart", params.DeletedDateStart)
	}
	if params.DeletedDateEnd != "" {
		query.Add("deletedDateEnd", params.DeletedDateEnd)
	}
	if params.RepairOrderNumber > 0 {
		query.Add("repairOrderNumber", fmt.Sprintf("%d", params.RepairOrderNumber))
	}
	for _, statusID := range params.RepairOrderStatusIds {
		query.Add("repairOrderStatusId", fmt.Sprintf("%d", statusID))
	}
	if params.CustomerID > 0 {
		query.Add("customerId", fmt.Sprintf("%d", params.CustomerID))
	}
	if params.VehicleID > 0 {
		query.Add("vehicleId", fmt.Sprintf("%d", params.VehicleID))
	}
	if params.Search != "" {
		query.Add("search", params.Search)
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/repair-orders?" + query.Encode()
	var resp PaginatedResponse[RepairOrder]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRepairOrder returns a specific repair order by ID
func (c *Client) GetRepairOrder(ctx context.Context, id int) (*RepairOrder, error) {
	var ro RepairOrder
	path := fmt.Sprintf("/api/v1/repair-orders/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &ro); err != nil {
		return nil, err
	}
	return &ro, nil
}

// GetJobs returns a paginated list of jobs
func (c *Client) GetJobs(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Job], error) {
	path := fmt.Sprintf("/api/v1/jobs?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Job]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJob returns a specific job by ID
func (c *Client) GetJob(ctx context.Context, id int) (*Job, error) {
	var job Job
	path := fmt.Sprintf("/api/v1/jobs/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// GetJobsWithParams returns jobs with advanced filtering
func (c *Client) GetJobsWithParams(ctx context.Context, params JobQueryParams) (*PaginatedResponse[Job], error) {
	query := url.Values{}
	if params.Shop > 0 {
		query.Add("shop", fmt.Sprintf("%d", params.Shop))
	}
	query.Add("page", fmt.Sprintf("%d", params.Page))
	if params.Size > 0 {
		query.Add("size", fmt.Sprintf("%d", params.Size))
	} else {
		query.Add("size", "100")
	}
	if params.RepairOrderID > 0 {
		query.Add("repairOrderId", fmt.Sprintf("%d", params.RepairOrderID))
	}
	if params.EmployeeID > 0 {
		query.Add("employeeId", fmt.Sprintf("%d", params.EmployeeID))
	}
	if params.Status != "" {
		query.Add("status", params.Status)
	}
	if params.Search != "" {
		query.Add("search", params.Search)
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/jobs?" + query.Encode()
	var resp PaginatedResponse[Job]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAppointments returns a paginated list of appointments
func (c *Client) GetAppointments(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Appointment], error) {
	path := fmt.Sprintf("/api/v1/appointments?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Appointment]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAppointment returns a specific appointment by ID
func (c *Client) GetAppointment(ctx context.Context, id int) (*Appointment, error) {
	var appointment Appointment
	path := fmt.Sprintf("/api/v1/appointments/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &appointment); err != nil {
		return nil, err
	}
	return &appointment, nil
}

// GetAppointmentsWithParams returns appointments with advanced filtering
func (c *Client) GetAppointmentsWithParams(ctx context.Context, params AppointmentQueryParams) (*PaginatedResponse[Appointment], error) {
	query := url.Values{}
	if params.Shop > 0 {
		query.Add("shop", fmt.Sprintf("%d", params.Shop))
	}
	query.Add("page", fmt.Sprintf("%d", params.Page))
	if params.Size > 0 {
		query.Add("size", fmt.Sprintf("%d", params.Size))
	} else {
		query.Add("size", "100")
	}
	if params.StartDate != "" {
		query.Add("start", params.StartDate)
	}
	if params.EndDate != "" {
		query.Add("end", params.EndDate)
	}
	if params.CustomerID > 0 {
		query.Add("customerId", fmt.Sprintf("%d", params.CustomerID))
	}
	if params.VehicleID > 0 {
		query.Add("vehicleId", fmt.Sprintf("%d", params.VehicleID))
	}
	if params.Status != "" {
		query.Add("status", params.Status)
	}
	if params.Search != "" {
		query.Add("search", params.Search)
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/appointments?" + query.Encode()
	var resp PaginatedResponse[Appointment]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEmployees returns a paginated list of employees
func (c *Client) GetEmployees(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Employee], error) {
	path := fmt.Sprintf("/api/v1/employees?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Employee]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEmployee returns a specific employee by ID
func (c *Client) GetEmployee(ctx context.Context, id int) (*Employee, error) {
	var employee Employee
	path := fmt.Sprintf("/api/v1/employees/%d", id)
	if err := c.doRequest(ctx, "GET", path, nil, &employee); err != nil {
		return nil, err
	}
	return &employee, nil
}

// GetEmployeesWithParams returns employees with advanced filtering
func (c *Client) GetEmployeesWithParams(ctx context.Context, params EmployeeQueryParams) (*PaginatedResponse[Employee], error) {
	query := url.Values{}
	// Shop parameter is optional but recommended
	if params.Shop > 0 {
		query.Add("shop", fmt.Sprintf("%d", params.Shop))
	}
	query.Add("page", fmt.Sprintf("%d", params.Page))
	if params.Size > 0 {
		query.Add("size", fmt.Sprintf("%d", params.Size))
	} else {
		query.Add("size", "100")
	}
	if params.Search != "" {
		query.Add("search", params.Search)
	}
	if params.UpdatedDateStart != "" {
		query.Add("updatedDateStart", params.UpdatedDateStart)
	}
	if params.UpdatedDateEnd != "" {
		query.Add("updatedDateEnd", params.UpdatedDateEnd)
	}
	if params.Active != nil {
		query.Add("active", fmt.Sprintf("%t", *params.Active))
	}
	if params.Role != "" {
		query.Add("role", params.Role)
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/employees?" + query.Encode()
	var resp PaginatedResponse[Employee]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInventory returns a paginated list of inventory parts
func (c *Client) GetInventory(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[InventoryPart], error) {
	path := fmt.Sprintf("/api/v1/inventory?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[InventoryPart]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCannedJobs returns a paginated list of canned jobs
func (c *Client) GetCannedJobs(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[CannedJob], error) {
	path := fmt.Sprintf("/api/v1/canned-jobs?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[CannedJob]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

