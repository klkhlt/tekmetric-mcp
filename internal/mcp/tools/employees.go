package tools

import (
	"context"
	"fmt"

	"github.com/beetlebugorg/tekmetric-mcp/internal/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterEmployeeTools registers all employee-related tools
func (r *Registry) RegisterEmployeeTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("employees",
			mcp.WithDescription("Search and filter employees/technicians, or get a specific employee by ID. Supports filtering by active status, role, and more."),
			mcp.WithNumber("id",
				mcp.Description("Get specific employee by ID"),
			),
			mcp.WithString("search",
				mcp.Description("Search employees by name or email"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to TEKMETRIC_DEFAULT_SHOP_ID)"),
			),
			mcp.WithBoolean("active",
				mcp.Description("Filter by active status (true for active employees only, false for inactive)"),
			),
			mcp.WithString("role",
				mcp.Description("Filter by employee role (e.g., technician, service advisor, manager)"),
			),
			mcp.WithString("sort",
				mcp.Description("Property to sort results by (e.g., firstName, lastName, email)"),
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
		r.handleEmployees,
	)

	r.logger.Debug("registered employee tools")
}

// handleEmployees searches employees or gets a specific employee by ID
func (r *Registry) handleEmployees(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	// If ID is provided, get specific employee
	if id, ok := parseFloatArg(arguments, "id"); ok {
		employee, err := r.client.GetEmployee(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get employee: %v", err)), nil
		}
		return formatJSON(employee)
	}

	// Otherwise, search with filters
	// Default to 10 results to avoid overwhelming context
	params := tekmetric.EmployeeQueryParams{
		Shop: r.config.Tekmetric.DefaultShopID,
		Page: 0,
		Size: 10,
	}

	// Parse optional parameters
	if shop, ok := parseFloatArg(arguments, "shop"); ok {
		params.Shop = shop
	}
	if search, ok := parseStringArg(arguments, "search"); ok {
		params.Search = search
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

	resp, err := r.client.GetEmployeesWithParams(ctx, params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to search employees: %v", err)), nil
	}

	return formatPaginatedResultWithWarning(
		resp.Content,
		resp.TotalElements,
		len(resp.Content),
		25,
		"EMPLOYEES",
	)
}
