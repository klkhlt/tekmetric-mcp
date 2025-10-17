package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

// CustomerQueryParams holds query parameters for customer searches
type CustomerQueryParams struct {
	Shop                          int    `url:"shop,omitempty"`
	Page                          int    `url:"page,omitempty"`
	Size                          int    `url:"size,omitempty"`
	Search                        string `url:"search,omitempty"`                        // Search by name, email, phone
	Email                         string `url:"email,omitempty"`                         // Filter by email
	Phone                         string `url:"phone,omitempty"`                         // Filter by phone
	EligibleForAccountsReceivable *bool  `url:"eligibleForAccountsReceivable,omitempty"` // Filter by AR eligibility
	OkForMarketing                *bool  `url:"okForMarketing,omitempty"`                // Filter by marketing permission
	UpdatedDateStart              string `url:"updatedDateStart,omitempty"`              // Filter by updated date
	UpdatedDateEnd                string `url:"updatedDateEnd,omitempty"`                // Filter by updated date
	DeletedDateStart              string `url:"deletedDateStart,omitempty"`              // Filter by deleted date
	DeletedDateEnd                string `url:"deletedDateEnd,omitempty"`                // Filter by deleted date
	CustomerTypeID                int    `url:"customerTypeId,omitempty"`                // 1=Customer, 2=Business
	Sort                          string `url:"sort,omitempty"`                          // lastName, firstName, email (can be comma-separated)
	SortDirection                 string `url:"sortDirection,omitempty"`                 // ASC, DESC
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

// GetCustomersWithParams returns customers with advanced filtering
func (c *Client) GetCustomersWithParams(ctx context.Context, params CustomerQueryParams) (*PaginatedResponse[Customer], error) {
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
	if params.Search != "" {
		query.Add("search", params.Search)
	}
	if params.Email != "" {
		query.Add("email", params.Email)
	}
	if params.Phone != "" {
		query.Add("phone", params.Phone)
	}
	if params.EligibleForAccountsReceivable != nil {
		query.Add("eligibleForAccountsReceivable", fmt.Sprintf("%t", *params.EligibleForAccountsReceivable))
	}
	if params.OkForMarketing != nil {
		query.Add("okForMarketing", fmt.Sprintf("%t", *params.OkForMarketing))
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
	if params.CustomerTypeID > 0 {
		query.Add("customerTypeId", fmt.Sprintf("%d", params.CustomerTypeID))
	}
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/customers?" + query.Encode()
	var resp PaginatedResponse[Customer]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
