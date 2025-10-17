package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/beetlebugorg/tekmetric-mcp/pkg/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterVehicleTools registers all vehicle-related tools
func (r *Registry) RegisterVehicleTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("vehicles",
			mcp.WithDescription("Search for vehicles by VIN, license plate, make/model, or get vehicle by ID. Supports filtering by customer and date ranges."),
			mcp.WithNumber("id",
				mcp.Description("Get specific vehicle by ID"),
			),
			mcp.WithString("search",
				mcp.Description("Search vehicles by VIN, license plate, year, make, or model"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to configured shop)"),
			),
			mcp.WithNumber("customer_id",
				mcp.Description("Filter vehicles by customer ID"),
			),
			mcp.WithString("updated_date_start",
				mcp.Description("Filter by updated date start (YYYY-MM-DD)"),
			),
			mcp.WithString("updated_date_end",
				mcp.Description("Filter by updated date end (YYYY-MM-DD)"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field (e.g., year, make, model)"),
			),
			mcp.WithString("sort_direction",
				mcp.Description("Sort direction: ASC or DESC"),
			),
			mcp.WithNumber("limit",
				mcp.Description("Maximum results to return (max: 100, default: 10)"),
			),
			mcp.WithNumber("page",
				mcp.Description("Page number for pagination (default: 0)"),
			),
		),
		r.handleVehicles,
	)

	r.logger.Debug("registered vehicle tools")
}

// handleVehicles handles vehicle search and retrieval
func (r *Registry) handleVehicles(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	// If ID is provided, get specific vehicle
	if id, ok := parseFloatArg(arguments, "id"); ok {
		vehicle, err := r.client.GetVehicle(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get vehicle: %v", err)), nil
		}
		return formatJSON(vehicle)
	}

	// Build query params
	params := tekmetric.VehicleQueryParams{
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
	if customerID, ok := parseFloatArg(arguments, "customer_id"); ok {
		params.CustomerID = customerID
	}
	if updatedStart, ok := parseStringArg(arguments, "updated_date_start"); ok {
		params.UpdatedDateStart = updatedStart
	}
	if updatedEnd, ok := parseStringArg(arguments, "updated_date_end"); ok {
		params.UpdatedDateEnd = updatedEnd
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

	resp, err := r.client.GetVehiclesWithParams(ctx, params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to search vehicles: %v", err)), nil
	}

	return formatPaginatedResultWithWarning(
		resp.Content,
		resp.TotalElements,
		len(resp.Content),
		25,
		"VEHICLES",
	)
}

// formatVehicleSummary creates a formatted summary of a vehicle
func (r *Registry) formatVehicleSummary(v *tekmetric.Vehicle) (*mcp.CallToolResult, error) {
	var summary strings.Builder

	// Header
	vehicleName := fmt.Sprintf("%d %s %s", v.Year, v.Make, v.Model)
	if v.SubModel != "" {
		vehicleName += fmt.Sprintf(" %s", v.SubModel)
	}
	if v.Color != "" {
		vehicleName += fmt.Sprintf(" (%s)", v.Color)
	}
	summary.WriteString(vehicleName + "\n")
	summary.WriteString(fmt.Sprintf("Vehicle ID: %d\n\n", v.ID))

	// Identification
	if v.VIN != "" {
		summary.WriteString(fmt.Sprintf("VIN: %s\n", v.VIN))
	}
	if v.LicensePlate != "" {
		summary.WriteString(fmt.Sprintf("License Plate: %s\n", v.LicensePlate))
	}
	if v.UnitNumber != "" {
		summary.WriteString(fmt.Sprintf("Unit Number: %s\n", v.UnitNumber))
	}

	// Mileage
	if v.Mileage > 0 {
		summary.WriteString(fmt.Sprintf("Current Mileage: %.0f miles\n", v.Mileage))
	}

	// Technical Specifications
	if v.Engine != "" || v.Transmission != "" || v.DriveType != "" {
		summary.WriteString("\n")
		if v.Engine != "" {
			summary.WriteString(fmt.Sprintf("Engine: %s\n", v.Engine))
		}
		if v.Transmission != "" {
			summary.WriteString(fmt.Sprintf("Transmission: %s\n", v.Transmission))
		}
		if v.DriveType != "" {
			summary.WriteString(fmt.Sprintf("Drive Type: %s\n", v.DriveType))
		}
	}

	// Notes
	if v.Notes != "" {
		summary.WriteString(fmt.Sprintf("\nNotes: %s\n", v.Notes))
	}

	// Metadata
	summary.WriteString(fmt.Sprintf("\nAdded: %s", v.CreatedDate.Format("January 2, 2006")))

	return formatRichResult(summary.String(), v)
}
