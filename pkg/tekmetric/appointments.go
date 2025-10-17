package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

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
	IncludeDeleted   *bool  `url:"includeDeleted,omitempty"`   // Include deleted appointments (default: false)
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// GetAppointments returns a paginated list of appointments (excludes deleted by default)
func (c *Client) GetAppointments(ctx context.Context, shopID int, page int, size int) (*PaginatedResponse[Appointment], error) {
	if err := c.isAuthorizedShop(shopID); err != nil {
		return nil, err
	}
	includeDeleted := false
	path := fmt.Sprintf("/api/v1/appointments?shop=%d&page=%d&size=%d&includeDeleted=%t", shopID, page, size, includeDeleted)
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
	// Default to excluding deleted appointments
	if params.IncludeDeleted != nil {
		query.Add("includeDeleted", fmt.Sprintf("%t", *params.IncludeDeleted))
	} else {
		query.Add("includeDeleted", "false")
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
