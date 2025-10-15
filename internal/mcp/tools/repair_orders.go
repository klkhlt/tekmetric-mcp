package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/beetlebugorg/tekmetric-mcp/internal/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterRepairOrderTools registers all repair order-related tools
func (r *Registry) RegisterRepairOrderTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("repair_orders",
			mcp.WithDescription("Search and filter repair orders, or get specific RO by ID. Search by RO#, customer name, or vehicle info (make/model/VIN). Supports filtering by date range, status, customer ID, vehicle ID. Returns RO details including jobs, parts, labor, and totals. **IMPORTANT: Default returns 10 results. For broad queries like 'all repair orders' or 'current repair orders', ALWAYS add filters (status, date range, customer) to narrow results. For analytics, use date ranges and increase limit parameter.**"),
			mcp.WithNumber("id",
				mcp.Description("Get specific repair order by ID"),
			),
			mcp.WithString("search",
				mcp.Description("Search by RO#, customer name, or vehicle info (e.g., 'Ford', 'Smith', '12345')"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to configured shop)"),
			),
			mcp.WithString("start_date",
				mcp.Description("Filter by created after date (YYYY-MM-DD)"),
			),
			mcp.WithString("end_date",
				mcp.Description("Filter by created before date (YYYY-MM-DD)"),
			),
			mcp.WithString("status",
				mcp.Description("Filter by status: estimate, wip, complete, saved, posted, ar, deleted"),
			),
			mcp.WithNumber("customer_id",
				mcp.Description("Filter by customer ID"),
			),
			mcp.WithNumber("vehicle_id",
				mcp.Description("Filter by vehicle ID"),
			),
			mcp.WithNumber("limit",
				mcp.Description("Maximum results (default 10, max 25). Keep queries focused with filters."),
			),
		),
		r.handleRepairOrders,
	)

	r.logger.Debug("registered repair order tools")
}

// handleRepairOrders handles repair order search and retrieval
func (r *Registry) handleRepairOrders(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	// If ID is provided, get specific repair order
	if id, ok := parseFloatArg(arguments, "id"); ok {
		repairOrder, err := r.client.GetRepairOrder(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get repair order: %v", err)), nil
		}
		return formatJSON(repairOrder)
	}

	// Get shop ID
	shopID := r.config.Tekmetric.DefaultShopID
	if shop, ok := parseFloatArg(arguments, "shop"); ok {
		shopID = shop
	}

	// Default to 10 results to avoid overwhelming Claude's context
	limit := 10
	if lim, ok := parseFloatArg(arguments, "limit"); ok {
		limit = lim
	}

	// Cap at 25 total results to prevent context overload
	maxResults := 25
	if limit > maxResults {
		limit = maxResults
	}

	// Calculate pages needed (API max is 100 per page)
	pageSize := 100
	if limit < pageSize {
		pageSize = limit
	}

	// Build query params for search/filter
	params := tekmetric.RepairOrderQueryParams{
		Shop: shopID,
		Page: 0,
		Size: pageSize,
	}

	// Use the native search parameter (searches RO#, customer name, vehicle info)
	if search, ok := parseStringArg(arguments, "search"); ok {
		params.Search = search
	}

	if startDate, ok := parseDateArg(arguments, "start_date"); ok {
		params.Start = startDate
	}
	if endDate, ok := parseDateArg(arguments, "end_date"); ok {
		params.End = endDate
	}
	if status, ok := parseStringArg(arguments, "status"); ok {
		// Convert status names to IDs: estimate=1, wip=2, complete=3, saved=4, posted=5, ar=6, deleted=7
		statusMap := map[string]int{
			"estimate": 1, "wip": 2, "complete": 3, "saved": 4,
			"posted": 5, "ar": 6, "deleted": 7,
		}
		for _, s := range strings.Split(status, ",") {
			s = strings.TrimSpace(strings.ToLower(s))
			if statusID, ok := statusMap[s]; ok {
				params.RepairOrderStatusIds = append(params.RepairOrderStatusIds, statusID)
			}
		}
	}
	if customerID, ok := parseFloatArg(arguments, "customer_id"); ok {
		params.CustomerID = customerID
	}
	if vehicleID, ok := parseFloatArg(arguments, "vehicle_id"); ok {
		params.VehicleID = vehicleID
	}

	// Fetch first page
	repairOrders, err := r.client.GetRepairOrdersWithParams(ctx, params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get repair orders: %v", err)), nil
	}

	allResults := repairOrders.Content
	totalAvailable := repairOrders.TotalElements

	// Fetch additional pages if needed (up to maxResults)
	pagesNeeded := (limit + pageSize - 1) / pageSize // Ceiling division
	for page := 1; page < pagesNeeded && len(allResults) < limit && len(allResults) < totalAvailable; page++ {
		params.Page = page
		nextPage, err := r.client.GetRepairOrdersWithParams(ctx, params)
		if err != nil {
			r.logger.Warn("failed to fetch additional page", "page", page, "error", err)
			break
		}
		allResults = append(allResults, nextPage.Content...)
		if len(allResults) >= limit {
			allResults = allResults[:limit]
			break
		}
	}

	// Create response with clear warning message
	response := map[string]interface{}{
		"data":          allResults,
		"totalElements": totalAvailable,
		"returned":      len(allResults),
	}

	// Add prominent warning if results were truncated
	if totalAvailable > maxResults {
		response["WARNING"] = fmt.Sprintf("⚠️ SHOWING ONLY %d OF %d TOTAL REPAIR ORDERS ⚠️", len(allResults), totalAvailable)
		response["message"] = "Results are limited. To see more results, add filters like date range (start_date/end_date), status, or customer_id to narrow your search."
		response["truncated"] = true
	} else if totalAvailable > len(allResults) {
		response["WARNING"] = fmt.Sprintf("⚠️ SHOWING %d OF %d RESULTS ⚠️", len(allResults), totalAvailable)
		response["message"] = "Increase the 'limit' parameter to see more results."
		response["truncated"] = true
	}

	return formatJSON(response)
}

// formatRepairOrderSummary creates a rich formatted summary of a repair order
func (r *Registry) formatRepairOrderSummary(ro *tekmetric.RepairOrder) (*mcp.CallToolResult, error) {
	var summary strings.Builder

	// Header
	summary.WriteString(fmt.Sprintf("# Repair Order #%d\n\n", ro.RepairOrderNumber))

	// Status
	summary.WriteString(fmt.Sprintf("**Status:** %s\n", ro.RepairOrderStatus.Name))
	if ro.CompletedDate != nil {
		summary.WriteString(fmt.Sprintf("**Completed:** %s\n", ro.CompletedDate.Format("2006-01-02 15:04")))
	}
	if ro.PostedDate != nil {
		summary.WriteString(fmt.Sprintf("**Posted:** %s\n", ro.PostedDate.Format("2006-01-02 15:04")))
	}
	summary.WriteString("\n")

	// Financial Summary
	summary.WriteString("## Financial Summary\n\n")
	summary.WriteString(fmt.Sprintf("- **Labor:** %s\n", formatCurrency(ro.LaborSales)))
	summary.WriteString(fmt.Sprintf("- **Parts:** %s\n", formatCurrency(ro.PartsSales)))
	summary.WriteString(fmt.Sprintf("- **Sublet:** %s\n", formatCurrency(ro.SubletSales)))
	summary.WriteString(fmt.Sprintf("- **Fees:** %s\n", formatCurrency(ro.FeeTotal)))
	summary.WriteString(fmt.Sprintf("- **Discounts:** -%s\n", formatCurrency(ro.DiscountTotal)))
	summary.WriteString(fmt.Sprintf("- **Taxes:** %s\n", formatCurrency(ro.Taxes)))
	summary.WriteString(fmt.Sprintf("- **Total:** %s\n", formatCurrency(ro.TotalSales)))
	summary.WriteString(fmt.Sprintf("- **Amount Paid:** %s\n", formatCurrency(ro.AmountPaid)))

	balance := ro.TotalSales - ro.AmountPaid
	if balance > 0 {
		summary.WriteString(fmt.Sprintf("- **Balance Due:** %s\n", formatCurrency(balance)))
	}
	summary.WriteString("\n")

	// Jobs
	if len(ro.Jobs) > 0 {
		summary.WriteString(fmt.Sprintf("## Jobs (%d)\n\n", len(ro.Jobs)))
		for i, job := range ro.Jobs {
			summary.WriteString(fmt.Sprintf("%d. **%s**\n", i+1, job.Name))
			if job.Note != "" {
				summary.WriteString(fmt.Sprintf("   - %s\n", job.Note))
			}
			if job.JobCategoryName != "" {
				summary.WriteString(fmt.Sprintf("   - Category: %s\n", job.JobCategoryName))
			}
			summary.WriteString(fmt.Sprintf("   - Labor: %s | Parts: %s | Subtotal: %s\n",
				formatCurrency(job.LaborTotal),
				formatCurrency(job.PartsTotal),
				formatCurrency(job.Subtotal)))
			if job.Authorized {
				summary.WriteString("   - ✓ Authorized\n")
			}
		}
		summary.WriteString("\n")
	}

	// Customer Concerns
	if len(ro.CustomerConcerns) > 0 {
		summary.WriteString("## Customer Concerns\n\n")
		for _, concern := range ro.CustomerConcerns {
			summary.WriteString(fmt.Sprintf("- %s\n", concern.Concern))
		}
		summary.WriteString("\n")
	}

	// Vehicle Info
	if ro.MilesIn != nil {
		summary.WriteString(fmt.Sprintf("**Mileage In:** %.0f\n", *ro.MilesIn))
	}
	if ro.MilesOut != nil {
		summary.WriteString(fmt.Sprintf("**Mileage Out:** %.0f\n", *ro.MilesOut))
	}

	return formatRichResult(summary.String(), ro)
}
