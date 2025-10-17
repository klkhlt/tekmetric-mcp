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
			mcp.WithDescription("Search for vehicles by VIN, license plate, make/model, or get vehicle by ID. Returns vehicle details including make, model, year, and service history."),
			mcp.WithNumber("id",
				mcp.Description("Get specific vehicle by ID"),
			),
			mcp.WithString("query",
				mcp.Description("Search vehicles by VIN, license plate, or make/model"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to configured shop)"),
			),
			mcp.WithNumber("limit",
				mcp.Description("Maximum results to return for search (default 10)"),
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

	// Get shop ID
	shopID := r.config.Tekmetric.DefaultShopID
	if shop, ok := parseFloatArg(arguments, "shop"); ok {
		shopID = shop
	}

	// If query is provided, search vehicles
	if query, ok := parseStringArg(arguments, "query"); ok {
		limit := 10
		if lim, ok := parseFloatArg(arguments, "limit"); ok {
			limit = lim
		}

		// Use API's native search instead of client-side filtering
		vehicles, err := r.client.SearchVehicles(ctx, shopID, query, 0, limit)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to search vehicles: %v", err)), nil
		}

		response := map[string]interface{}{
			"query":         query,
			"returned":      len(vehicles.Content),
			"totalElements": vehicles.TotalElements,
			"results":       vehicles.Content,
		}

		// Warn if there are more results available
		if vehicles.TotalElements > limit {
			response["WARNING"] = fmt.Sprintf("⚠️ SHOWING %d OF %d MATCHING VEHICLES ⚠️", len(vehicles.Content), vehicles.TotalElements)
			response["message"] = "Use a more specific search query or increase the 'limit' parameter."
		}

		return formatJSON(response)
	}

	// No ID or query - return error suggesting what to provide
	return mcp.NewToolResultError("Please provide either 'id' to get a specific vehicle or 'query' to search vehicles"), nil
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
