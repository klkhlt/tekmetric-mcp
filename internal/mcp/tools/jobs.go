package tools

import (
	"context"
	"fmt"

	"github.com/beetlebugorg/tekmetric-mcp/internal/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterJobTools registers all job-related tools
func (r *Registry) RegisterJobTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("jobs",
			mcp.WithDescription("Search and filter jobs (work items/services on repair orders), or get a specific job by ID. Supports filtering by repair order, employee, status, and more."),
			mcp.WithNumber("id",
				mcp.Description("Get specific job by ID"),
			),
			mcp.WithString("search",
				mcp.Description("Search jobs by name or description"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to TEKMETRIC_DEFAULT_SHOP_ID)"),
			),
			mcp.WithNumber("repair_order_id",
				mcp.Description("Filter by repair order ID"),
			),
			mcp.WithNumber("employee_id",
				mcp.Description("Filter by assigned employee/technician ID"),
			),
			mcp.WithString("status",
				mcp.Description("Filter by job status"),
			),
			mcp.WithString("sort",
				mcp.Description("Property to sort results by (e.g., createdDate, name)"),
			),
			mcp.WithString("sort_direction",
				mcp.Description("Sort direction (ASC or DESC)"),
			),
			mcp.WithNumber("limit",
				mcp.Description("Maximum number of results to return (max: 100, default: 20)"),
			),
			mcp.WithNumber("page",
				mcp.Description("Page number for pagination (default: 0)"),
			),
		),
		r.handleJobs,
	)

	r.logger.Debug("registered job tools")
}

// handleJobs searches jobs or gets a specific job by ID
func (r *Registry) handleJobs(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	// If ID is provided, get specific job
	if id, ok := parseFloatArg(arguments, "id"); ok {
		job, err := r.client.GetJob(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get job: %v", err)), nil
		}
		return formatJSON(job)
	}

	// Otherwise, search with filters
	params := tekmetric.JobQueryParams{
		Shop: r.config.Tekmetric.DefaultShopID,
		Page: 0,
		Size: 20,
	}

	// Parse optional parameters
	if shop, ok := parseFloatArg(arguments, "shop"); ok {
		params.Shop = shop
	}
	if repairOrderID, ok := parseFloatArg(arguments, "repair_order_id"); ok {
		params.RepairOrderID = repairOrderID
	}
	if employeeID, ok := parseFloatArg(arguments, "employee_id"); ok {
		params.EmployeeID = employeeID
	}
	if search, ok := parseStringArg(arguments, "search"); ok {
		params.Search = search
	}
	if status, ok := parseStringArg(arguments, "status"); ok {
		params.Status = status
	}
	if sort, ok := parseStringArg(arguments, "sort"); ok {
		params.Sort = sort
	}
	if sortDirection, ok := parseStringArg(arguments, "sort_direction"); ok {
		params.SortDirection = sortDirection
	}
	if limit, ok := parseFloatArg(arguments, "limit"); ok {
		params.Size = limit
		if params.Size > 100 {
			params.Size = 100
		}
	}
	if page, ok := parseFloatArg(arguments, "page"); ok {
		params.Page = page
	}

	resp, err := r.client.GetJobsWithParams(ctx, params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to search jobs: %v", err)), nil
	}

	return formatJSON(resp)
}
