package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/beetlebugorg/tekmetric-mcp/pkg/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterCustomerTools registers all customer-related tools
func (r *Registry) RegisterCustomerTools(s *server.MCPServer) {
	// Search/list customers - consolidated tool
	s.AddTool(
		mcp.NewTool("customers",
			mcp.WithDescription("Search for customers by name, email, phone, or get customer by ID. Returns customer details including contact info and vehicles."),
			mcp.WithNumber("id",
				mcp.Description("Get specific customer by ID"),
			),
			mcp.WithString("query",
				mcp.Description("Search customers by name, email, or phone"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to configured shop)"),
			),
			mcp.WithNumber("limit",
				mcp.Description("Maximum results to return for search (default 10)"),
			),
		),
		r.handleCustomers,
	)

	r.logger.Debug("registered customer tools")
}

// handleCustomers handles customer search and retrieval
func (r *Registry) handleCustomers(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	// If ID is provided, get specific customer
	if id, ok := parseFloatArg(arguments, "id"); ok {
		customer, err := r.client.GetCustomer(ctx, id)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get customer: %v", err)), nil
		}
		return formatJSON(customer)
	}

	// Get shop ID
	shopID := r.config.Tekmetric.DefaultShopID
	if shop, ok := parseFloatArg(arguments, "shop"); ok {
		shopID = shop
	}

	// If query is provided, search customers
	if query, ok := parseStringArg(arguments, "query"); ok {
		limit := 10
		if lim, ok := parseFloatArg(arguments, "limit"); ok {
			limit = lim
		}

		// Use API's native search instead of client-side filtering
		customers, err := r.client.SearchCustomers(ctx, shopID, query, 0, limit)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to search customers: %v", err)), nil
		}

		response := map[string]interface{}{
			"query":         query,
			"returned":      len(customers.Content),
			"totalElements": customers.TotalElements,
			"results":       customers.Content,
		}

		// Warn if there are more results available
		if customers.TotalElements > limit {
			response["WARNING"] = fmt.Sprintf("⚠️ SHOWING %d OF %d MATCHING CUSTOMERS ⚠️", len(customers.Content), customers.TotalElements)
			response["message"] = "Use a more specific search query or increase the 'limit' parameter."
		}

		return formatJSON(response)
	}

	// No ID or query - return error suggesting what to provide
	return mcp.NewToolResultError("Please provide either 'id' to get a specific customer or 'query' to search customers"), nil
}

// formatCustomerSummary creates a formatted summary of a customer
func (r *Registry) formatCustomerSummary(c *tekmetric.Customer) (*mcp.CallToolResult, error) {
	var summary strings.Builder

	// Header
	summary.WriteString(fmt.Sprintf("%s %s", c.FirstName, c.LastName))
	if c.CustomerType != nil {
		summary.WriteString(fmt.Sprintf(" (%s)", c.CustomerType.Name))
	}
	summary.WriteString(fmt.Sprintf("\nCustomer ID: %d\n\n", c.ID))

	// Contact Information
	if c.Email != "" {
		summary.WriteString(fmt.Sprintf("Email: %s\n", c.Email))
	}

	if len(c.Phone) > 0 {
		for _, phone := range c.Phone {
			phoneType := phone.Type
			if phoneType == "" {
				phoneType = "Phone"
			}
			primary := ""
			if phone.Primary {
				primary = " (Primary)"
			}
			summary.WriteString(fmt.Sprintf("%s: %s%s\n", phoneType, phone.Number, primary))
		}
	}

	if c.Address != nil && (c.Address.Address1 != "" || c.Address.City != "") {
		summary.WriteString("\nAddress:\n")
		if c.Address.Address1 != "" {
			summary.WriteString(fmt.Sprintf("  %s\n", c.Address.Address1))
		}
		if c.Address.Address2 != "" {
			summary.WriteString(fmt.Sprintf("  %s\n", c.Address.Address2))
		}
		if c.Address.City != "" {
			cityLine := fmt.Sprintf("  %s", c.Address.City)
			if c.Address.State != "" {
				cityLine += fmt.Sprintf(", %s", c.Address.State)
			}
			if c.Address.Zip != "" {
				cityLine += fmt.Sprintf(" %s", c.Address.Zip)
			}
			summary.WriteString(cityLine + "\n")
		}
	}

	// Account Information
	if c.EligibleForAccountsReceivable || c.CreditLimit > 0 || c.OkForMarketing {
		summary.WriteString("\n")
		if c.EligibleForAccountsReceivable {
			summary.WriteString("Accounts Receivable: Yes\n")
		}
		if c.CreditLimit > 0 {
			summary.WriteString(fmt.Sprintf("Credit Limit: $%.2f\n", c.CreditLimit))
		}
		if c.OkForMarketing {
			summary.WriteString("Marketing: Yes\n")
		}
	}

	// Notes
	if c.Notes != "" {
		summary.WriteString(fmt.Sprintf("\nNotes: %s\n", c.Notes))
	}

	// Metadata
	summary.WriteString(fmt.Sprintf("\nCustomer Since: %s", c.CreatedDate.Format("January 2, 2006")))

	return formatRichResult(summary.String(), c)
}
