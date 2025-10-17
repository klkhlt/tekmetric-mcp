package tekmetric

import (
	"context"
	"fmt"
	"net/url"
)

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
