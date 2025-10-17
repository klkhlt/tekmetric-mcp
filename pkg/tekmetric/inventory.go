package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

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
