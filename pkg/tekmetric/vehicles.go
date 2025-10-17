package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

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

// GetVehiclesWithParams returns vehicles with advanced filtering
func (c *Client) GetVehiclesWithParams(ctx context.Context, params VehicleQueryParams) (*PaginatedResponse[Vehicle], error) {
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
	if params.Search != "" {
		query.Add("search", params.Search)
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
	if params.Sort != "" {
		query.Add("sort", params.Sort)
	}
	if params.SortDirection != "" {
		query.Add("sortDirection", params.SortDirection)
	}

	path := "/api/v1/vehicles?" + query.Encode()
	var resp PaginatedResponse[Vehicle]
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
