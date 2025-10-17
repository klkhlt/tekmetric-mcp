// Package tekmetric provides a client for the Tekmetric shop management API.
// It handles OAuth2 authentication, rate limiting, and provides methods for
// accessing shops, customers, vehicles, repair orders, and other resources.
package tekmetric

import (
	"bytes"
	"context"
	"crypto/tls"
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
	"github.com/beetlebugorg/tekmetric-mcp/pkg/retry"
	"golang.org/x/time/rate"
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

// validateSortParams validates sort field and direction parameters
func validateSortParams(sort, sortDirection string, validSorts []string) error {
	// Validate sort direction
	if sortDirection != "" {
		upper := strings.ToUpper(sortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", sortDirection)
		}
	}

	// Validate sort field
	if sort != "" {
		valid := false
		for _, validSort := range validSorts {
			if sort == validSort {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid sort field '%s'", sort)
		}
	}

	return nil
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
	baseURL       string                 // API base URL (sandbox or production)
	clientID      string                 // OAuth2 client ID
	clientSecret  string                 // OAuth2 client secret
	httpClient    *http.Client           // HTTP client with timeout
	accessToken   string         // Current OAuth2 access token
	tokenExpiry   time.Time      // Token expiration time
	shopIDs       []string       // Shop IDs from token scope
	retryer       *retry.Retryer // Retry logic with exponential backoff
	globalLimiter *rate.Limiter  // Global rate limiter (requests per second)
	logger        *slog.Logger   // Structured logger
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
	// Create HTTP transport with secure TLS configuration
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12, // Enforce TLS 1.2 minimum
			MaxVersion: 0,                 // Allow highest available version
		},
	}

	return &Client{
		baseURL:      cfg.BaseURL,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		httpClient: &http.Client{
			Timeout:   time.Duration(cfg.TimeoutSeconds) * time.Second,
			Transport: transport,
		},
		retryer:       retry.New(cfg.MaxRetries, cfg.MaxBackoffSec),
		globalLimiter: rate.NewLimiter(rate.Limit(10), 10), // 10 requests/sec with burst of 10
		logger:        logger,
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
		c.logger.Debug("authentication failed", "status", resp.StatusCode, "body", string(body))
		return fmt.Errorf("authentication failed with status %d", resp.StatusCode)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	c.accessToken = tokenResp.AccessToken
	c.shopIDs = strings.Fields(tokenResp.Scope) // Space-separated shop IDs

	// Use expires_in from response if provided, otherwise assume 24h
	if tokenResp.ExpiresIn > 0 {
		c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
		c.logger.Info("authentication successful", "shop_count", len(c.shopIDs), "expires_in", tokenResp.ExpiresIn)
	} else {
		c.tokenExpiry = time.Now().Add(24 * time.Hour) // Fallback to 24h
		c.logger.Info("authentication successful", "shop_count", len(c.shopIDs), "expires_in", "24h (default)")
	}

	return nil
}

// ensureAuthenticated checks if we have a valid token and authenticates if needed
func (c *Client) ensureAuthenticated(ctx context.Context) error {
	if c.accessToken == "" || time.Now().After(c.tokenExpiry) {
		return c.Authenticate(ctx)
	}
	return nil
}

// isAuthorizedShop checks if the client is authorized to access the specified shop.
// Authorization is determined by the shop IDs in the OAuth token scope.
func (c *Client) isAuthorizedShop(shopID int) error {
	// Skip validation if shopID is 0 (not specified)
	if shopID == 0 {
		return nil
	}

	shopIDStr := fmt.Sprintf("%d", shopID)
	for _, authorizedID := range c.shopIDs {
		if authorizedID == shopIDStr {
			return nil
		}
	}
	return fmt.Errorf("unauthorized access to shop %d: not in token scope", shopID)
}

// doRequest performs an HTTP request with authentication and rate limiting
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	if err := c.ensureAuthenticated(ctx); err != nil {
		return err
	}

	// Wait for global rate limiter before making request
	if err := c.globalLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter wait failed: %w", err)
	}

	return c.retryer.Do(func() error {
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

		// Limit response body to prevent memory exhaustion (10MB max)
		maxBodySize := int64(10 * 1024 * 1024)
		limitedBody := io.LimitReader(resp.Body, maxBodySize)
		responseBody, err := io.ReadAll(limitedBody)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		// Check if we hit the size limit
		if int64(len(responseBody)) == maxBodySize {
			c.logger.Warn("response body may have been truncated", "path", path, "max_size", maxBodySize)
		}

		// Check for non-success status codes
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			// Log detailed error information
			c.logger.Debug("API request failed",
				"method", method,
				"path", path,
				"status", resp.StatusCode,
				"body", string(responseBody))

			// Rate limit (429) and server errors (5xx) are temporary - should retry
			if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
				return &temporaryError{
					statusCode: resp.StatusCode,
					message:    fmt.Sprintf("temporary error with status %d", resp.StatusCode),
				}
			}
			// Client errors (4xx except 429) are permanent - don't retry
			return fmt.Errorf("API request failed with status %d", resp.StatusCode)
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
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/api/v1/customers?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Customer]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchCustomers searches customers using the API's native search
func (c *Client) SearchCustomers(ctx context.Context, shopID int, query string, page int, size int) (*PaginatedResponse[Customer], error) {
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	query = url.QueryEscape(query)
	path := fmt.Sprintf("/api/v1/customers?shop=%d&search=%s&page=%d&size=%d", shopID, query, page, size)
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
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/api/v1/vehicles?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[Vehicle]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchVehicles searches vehicles using the API's native search
func (c *Client) SearchVehicles(ctx context.Context, shopID int, query string, page int, size int) (*PaginatedResponse[Vehicle], error) {
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	query = url.QueryEscape(query)
	path := fmt.Sprintf("/api/v1/vehicles?shop=%d&search=%s&page=%d&size=%d", shopID, query, page, size)
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

// Validate validates the RepairOrderQueryParams
func (p *RepairOrderQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort field - based on Tekmetric API documentation
	if p.Sort != "" {
		validSorts := map[string]bool{
			"createdDate":        true,
			"repairOrderNumber":  true,
			"customer.firstName": true,
			"customer.lastName":  true,
		}
		if !validSorts[p.Sort] {
			return fmt.Errorf("invalid sort field '%s': supported fields are createdDate, repairOrderNumber, customer.firstName, customer.lastName", p.Sort)
		}
	}

	// Validate repair order status IDs
	for _, statusID := range p.RepairOrderStatusIds {
		if statusID < 1 || statusID > 7 {
			return fmt.Errorf("invalid repairOrderStatusId '%d': must be 1-7 (1=Estimate, 2=WIP, 3=Complete, 4=Saved, 5=Posted, 6=AR, 7=Deleted)", statusID)
		}
	}

	return nil
}

// CustomerQueryParams holds query parameters for customer searches
type CustomerQueryParams struct {
	Shop                           int    `url:"shop,omitempty"`
	Page                           int    `url:"page,omitempty"`
	Size                           int    `url:"size,omitempty"`
	Search                         string `url:"search,omitempty"`                         // Search by name, email, phone
	Email                          string `url:"email,omitempty"`                          // Filter by email
	Phone                          string `url:"phone,omitempty"`                          // Filter by phone
	EligibleForAccountsReceivable  *bool  `url:"eligibleForAccountsReceivable,omitempty"`  // Filter by AR eligibility
	OkForMarketing                 *bool  `url:"okForMarketing,omitempty"`                 // Filter by marketing permission
	UpdatedDateStart               string `url:"updatedDateStart,omitempty"`               // Filter by updated date
	UpdatedDateEnd                 string `url:"updatedDateEnd,omitempty"`                 // Filter by updated date
	DeletedDateStart               string `url:"deletedDateStart,omitempty"`               // Filter by deleted date
	DeletedDateEnd                 string `url:"deletedDateEnd,omitempty"`                 // Filter by deleted date
	CustomerTypeID                 int    `url:"customerTypeId,omitempty"`                 // 1=Customer, 2=Business
	Sort                           string `url:"sort,omitempty"`                           // lastName, firstName, email (can be comma-separated)
	SortDirection                  string `url:"sortDirection,omitempty"`                  // ASC, DESC
}

// Validate validates the CustomerQueryParams
func (p *CustomerQueryParams) Validate() error {
	// Validate customer type ID
	if p.CustomerTypeID != 0 && p.CustomerTypeID != 1 && p.CustomerTypeID != 2 {
		return fmt.Errorf("invalid customerTypeId '%d': must be 1 (Customer) or 2 (Business)", p.CustomerTypeID)
	}

	// Validate sort - can be comma-separated list
	if p.Sort != "" {
		sortFields := strings.Split(p.Sort, ",")
		validSorts := map[string]bool{
			"lastName":  true,
			"firstName": true,
			"email":     true,
		}
		for _, field := range sortFields {
			trimmed := strings.TrimSpace(field)
			if !validSorts[trimmed] {
				return fmt.Errorf("invalid sort field '%s': supported fields are lastName, firstName, email", trimmed)
			}
		}
	}

	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	return nil
}

// VehicleQueryParams holds query parameters for vehicle searches
type VehicleQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	CustomerID       int    `url:"customerId,omitempty"`       // Filter by customer
	Search           string `url:"search,omitempty"`           // Search by year, make, model
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	DeletedDateStart string `url:"deletedDateStart,omitempty"` // Filter by deleted date
	DeletedDateEnd   string `url:"deletedDateEnd,omitempty"`   // Filter by deleted date
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// Validate validates the VehicleQueryParams
func (p *VehicleQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for vehicles
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// AppointmentQueryParams holds query parameters for appointment searches
type AppointmentQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	CustomerID       int    `url:"customerId,omitempty"`       // Filter by customer
	VehicleID        int    `url:"vehicleId,omitempty"`        // Filter by vehicle
	Start            string `url:"start,omitempty"`            // Start date filter
	End              string `url:"end,omitempty"`              // End date filter
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	IncludeDeleted   *bool  `url:"includeDeleted,omitempty"`   // Include deleted appointments (default: true)
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// Validate validates the AppointmentQueryParams
func (p *AppointmentQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for appointments
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// JobQueryParams holds query parameters for job searches
type JobQueryParams struct {
	Shop                 int    `url:"shop,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Size                 int    `url:"size,omitempty"`
	VehicleID            int    `url:"vehicleId,omitempty"`            // Filter by vehicle ID
	RepairOrderID        int    `url:"repairOrderId,omitempty"`        // Filter by repair order
	CustomerID           int    `url:"customerId,omitempty"`           // Filter by customer ID
	Authorized           *bool  `url:"authorized,omitempty"`           // Filter by authorized jobs
	AuthorizedDateStart  string `url:"authorizedDateStart,omitempty"`  // Filter by authorization date
	AuthorizedDateEnd    string `url:"authorizedDateEnd,omitempty"`    // Filter by authorization date
	UpdatedDateStart     string `url:"updatedDateStart,omitempty"`     // Filter by updated date
	UpdatedDateEnd       string `url:"updatedDateEnd,omitempty"`       // Filter by updated date
	RepairOrderStatusIds []int  `url:"repairOrderStatusId,omitempty"`  // 1-6 (no Deleted status for jobs)
	Sort                 string `url:"sort,omitempty"`                 // authorizedDate
	SortDirection        string `url:"sortDirection,omitempty"`        // ASC, DESC
}

// Validate validates the JobQueryParams
func (p *JobQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort field - based on Tekmetric API documentation
	if p.Sort != "" && p.Sort != "authorizedDate" {
		return fmt.Errorf("invalid sort field '%s': only 'authorizedDate' is supported", p.Sort)
	}

	// Validate repair order status IDs (jobs don't support status 7 - Deleted)
	for _, statusID := range p.RepairOrderStatusIds {
		if statusID < 1 || statusID > 6 {
			return fmt.Errorf("invalid repairOrderStatusId '%d': must be 1-6 (1=Estimate, 2=WIP, 3=Complete, 4=Saved, 5=Posted, 6=AR)", statusID)
		}
	}

	return nil
}

// EmployeeQueryParams holds query parameters for employee searches
type EmployeeQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	Search           string `url:"search,omitempty"`           // Search by name
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// Validate validates the EmployeeQueryParams
func (p *EmployeeQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for employees
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// InventoryQueryParams holds query parameters for inventory searches
type InventoryQueryParams struct {
	Shop          int      `url:"shop"`                        // Required: Shop ID
	PartTypeID    int      `url:"partTypeId"`                  // Required: 1=Part, 2=Tire, 5=Battery
	Page          int      `url:"page,omitempty"`
	Size          int      `url:"size,omitempty"`
	PartNumbers   []string `url:"partNumbers,omitempty"`       // Exact match on part numbers
	Width         string   `url:"width,omitempty"`             // Tire width (tires only)
	Ratio         float64  `url:"ratio,omitempty"`             // Tire ratio (tires only)
	Diameter      float64  `url:"diameter,omitempty"`          // Tire diameter (tires only)
	TireSize      string   `url:"tireSize,omitempty"`          // Tire size: width+ratio+diameter (tires only)
	Sort          string   `url:"sort,omitempty"`              // id, name, brand, partNumber (comma-separated)
	SortDirection string   `url:"sortDirection,omitempty"`     // ASC, DESC
}

// Validate validates the InventoryQueryParams
func (p *InventoryQueryParams) Validate() error {
	// Validate required fields
	if p.Shop == 0 {
		return fmt.Errorf("shop is required for inventory queries")
	}
	if p.PartTypeID == 0 {
		return fmt.Errorf("partTypeId is required for inventory queries")
	}

	// Validate part type ID
	if p.PartTypeID != 1 && p.PartTypeID != 2 && p.PartTypeID != 5 {
		return fmt.Errorf("invalid partTypeId '%d': must be 1 (Part), 2 (Tire), or 5 (Battery)", p.PartTypeID)
	}

	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort fields - can be comma-separated
	if p.Sort != "" {
		sortFields := strings.Split(p.Sort, ",")
		validSorts := map[string]bool{
			"id":         true,
			"name":       true,
			"brand":      true,
			"partNumber": true,
		}
		for _, field := range sortFields {
			trimmed := strings.TrimSpace(field)
			if !validSorts[trimmed] {
				return fmt.Errorf("invalid sort field '%s': supported fields are id, name, brand, partNumber", trimmed)
			}
		}
	}

	return nil
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
	if err := c.isAuthorizedShop(params.Shop); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
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
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
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
	if err := c.isAuthorizedShop(params.Shop); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
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
	if params.VehicleID > 0 {
		query.Add("vehicleId", fmt.Sprintf("%d", params.VehicleID))
	}
	if params.RepairOrderID > 0 {
		query.Add("repairOrderId", fmt.Sprintf("%d", params.RepairOrderID))
	}
	if params.CustomerID > 0 {
		query.Add("customerId", fmt.Sprintf("%d", params.CustomerID))
	}
	if params.Authorized != nil {
		query.Add("authorized", fmt.Sprintf("%t", *params.Authorized))
	}
	if params.AuthorizedDateStart != "" {
		query.Add("authorizedDateStart", params.AuthorizedDateStart)
	}
	if params.AuthorizedDateEnd != "" {
		query.Add("authorizedDateEnd", params.AuthorizedDateEnd)
	}
	if params.UpdatedDateStart != "" {
		query.Add("updatedDateStart", params.UpdatedDateStart)
	}
	if params.UpdatedDateEnd != "" {
		query.Add("updatedDateEnd", params.UpdatedDateEnd)
	}
	for _, statusID := range params.RepairOrderStatusIds {
		query.Add("repairOrderStatusId", fmt.Sprintf("%d", statusID))
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
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
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
	if err := c.isAuthorizedShop(params.Shop); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
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
	if params.CustomerID > 0 {
		query.Add("customerId", fmt.Sprintf("%d", params.CustomerID))
	}
	if params.VehicleID > 0 {
		query.Add("vehicleId", fmt.Sprintf("%d", params.VehicleID))
	}
	if params.Start != "" {
		query.Add("start", params.Start)
	}
	if params.End != "" {
		query.Add("end", params.End)
	}
	if params.UpdatedDateStart != "" {
		query.Add("updatedDateStart", params.UpdatedDateStart)
	}
	if params.UpdatedDateEnd != "" {
		query.Add("updatedDateEnd", params.UpdatedDateEnd)
	}
	if params.IncludeDeleted != nil {
		query.Add("includeDeleted", fmt.Sprintf("%t", *params.IncludeDeleted))
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
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
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
	if err := c.isAuthorizedShop(params.Shop); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
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
// Note: partTypeId is REQUIRED by the Tekmetric API (1=Part, 2=Tire, 5=Battery)
func (c *Client) GetInventory(ctx context.Context, shopID int, partTypeID int, page int, size int) (*PaginatedResponse[InventoryPart], error) {
	params := InventoryQueryParams{
		Shop:       shopID,
		PartTypeID: partTypeID,
		Page:       page,
		Size:       size,
	}
	return c.GetInventoryWithParams(ctx, params)
}

// GetInventoryWithParams returns inventory parts with advanced filtering
func (c *Client) GetInventoryWithParams(ctx context.Context, params InventoryQueryParams) (*PaginatedResponse[InventoryPart], error) {
	if err := c.isAuthorizedShop(params.Shop); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add("shop", fmt.Sprintf("%d", params.Shop))
	query.Add("partTypeId", fmt.Sprintf("%d", params.PartTypeID))
	query.Add("page", fmt.Sprintf("%d", params.Page))
	if params.Size > 0 {
		query.Add("size", fmt.Sprintf("%d", params.Size))
	} else {
		query.Add("size", "100")
	}
	for _, partNum := range params.PartNumbers {
		query.Add("partNumbers", partNum)
	}
	if params.Width != "" {
		query.Add("width", params.Width)
	}
	if params.Ratio != 0 {
		query.Add("ratio", fmt.Sprintf("%f", params.Ratio))
	}
	if params.Diameter != 0 {
		query.Add("diameter", fmt.Sprintf("%f", params.Diameter))
	}
	if params.TireSize != "" {
		query.Add("tireSize", params.TireSize)
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/inventory?" + query.Encode()
	var resp PaginatedResponse[InventoryPart]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCannedJobs returns a paginated list of canned jobs
func (c *Client) GetCannedJobs(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[CannedJob], error) {
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/api/v1/canned-jobs?shop=%d&page=%d&size=%d", shopID, page, size)
	var resp PaginatedResponse[CannedJob]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

