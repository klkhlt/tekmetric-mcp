package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

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
