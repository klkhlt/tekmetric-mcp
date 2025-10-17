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
			mcp.WithDescription("Search for customers by name, email, phone, or get customer by ID. Supports advanced filtering by customer type, AR eligibility, and date ranges."),
			mcp.WithNumber("id",
				mcp.Description("Get specific customer by ID"),
			),
			mcp.WithString("search",
				mcp.Description("Search customers by name, email, or phone"),
			),
			mcp.WithNumber("shop",
				mcp.Description("Shop ID (defaults to configured shop)"),
			),
			mcp.WithNumber("customer_type",
				mcp.Description("Customer type: 1=Customer, 2=Business"),
			),
			mcp.WithBoolean("ar_eligible",
				mcp.Description("Filter by accounts receivable eligibility"),
			),
			mcp.WithBoolean("ok_for_marketing",
				mcp.Description("Filter by marketing permission"),
			),
			mcp.WithString("updated_date_start",
				mcp.Description("Filter by updated date start (YYYY-MM-DD)"),
			),
			mcp.WithString("updated_date_end",
				mcp.Description("Filter by updated date end (YYYY-MM-DD)"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field: lastName, firstName, email (comma-separated)"),
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

	// Build query params
	params := tekmetric.CustomerQueryParams{
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
	if customerType, ok := parseFloatArg(arguments, "customer_type"); ok {
		params.CustomerTypeID = customerType
	}
	if arEligible, ok := parseBoolArg(arguments, "ar_eligible"); ok {
		params.EligibleForAccountsReceivable = &arEligible
	}
	if okMarketing, ok := parseBoolArg(arguments, "ok_for_marketing"); ok {
		params.OkForMarketing = &okMarketing
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

	resp, err := r.client.GetCustomersWithParams(ctx, params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to search customers: %v", err)), nil
	}

	return formatPaginatedResultWithWarning(
		resp.Content,
		resp.TotalElements,
		len(resp.Content),
		25,
		"CUSTOMERS",
	)
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
