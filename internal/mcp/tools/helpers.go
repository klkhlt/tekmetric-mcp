package tools

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/beetlebugorg/tekmetric-mcp/internal/tekmetric"
	"github.com/mark3labs/mcp-go/mcp"
)

// paginationParams holds common pagination parameters
type paginationParams struct {
	ShopID int
	Page   int
	Size   int
}

// parsePaginationArgs extracts common pagination arguments from tool arguments
func (r *Registry) parsePaginationArgs(arguments map[string]interface{}) paginationParams {
	params := paginationParams{
		ShopID: r.config.Tekmetric.DefaultShopID,
		Page:   0,
		Size:   100,
	}

	if shop, ok := arguments["shop"].(float64); ok {
		params.ShopID = int(shop)
	}
	if page, ok := arguments["page"].(float64); ok {
		params.Page = int(page)
	}
	if size, ok := arguments["size"].(float64); ok {
		params.Size = int(size)
		if params.Size > 100 {
			params.Size = 100
		}
	}

	return params
}

// parseFloatArg safely extracts a float64 argument and converts to int
func parseFloatArg(arguments map[string]interface{}, key string) (int, bool) {
	if val, ok := arguments[key].(float64); ok {
		return int(val), true
	}
	return 0, false
}

// parseStringArg safely extracts a string argument
func parseStringArg(arguments map[string]interface{}, key string) (string, bool) {
	if val, ok := arguments[key].(string); ok && val != "" {
		return val, true
	}
	return "", false
}

// formatJSON marshals data to indented JSON and returns a tool result
func formatJSON(data interface{}) (*mcp.CallToolResult, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format data: %v", err)), nil
	}
	return mcp.NewToolResultText(string(jsonData)), nil
}

// requireFloatArg extracts a required float64 argument and returns an error if missing
func requireFloatArg(arguments map[string]interface{}, key string) (int, *mcp.CallToolResult) {
	if val, ok := arguments[key].(float64); ok {
		return int(val), nil
	}
	return 0, mcp.NewToolResultError(fmt.Sprintf("%s parameter is required", key))
}

// requireStringArg extracts a required string argument and returns an error if missing
func requireStringArg(arguments map[string]interface{}, key string) (string, *mcp.CallToolResult) {
	if val, ok := arguments[key].(string); ok && val != "" {
		return val, nil
	}
	return "", mcp.NewToolResultError(fmt.Sprintf("%s parameter is required", key))
}

// parseDateArg parses a date string (YYYY-MM-DD) and returns it in ISO8601/RFC3339 format with timezone
// The Tekmetric API expects dates in ZonedDateTime format
func parseDateArg(arguments map[string]interface{}, key string) (string, bool) {
	dateStr, ok := parseStringArg(arguments, key)
	if !ok {
		return "", false
	}

	// Try to parse as YYYY-MM-DD
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// If parsing fails, return the original string
		return dateStr, true
	}

	// Convert to start of day in local timezone with RFC3339 format
	return t.Format(time.RFC3339), true
}

// formatRichResult creates a tool result with both formatted text and JSON data
// This provides a better user experience in Claude Desktop
func formatRichResult(summary string, data interface{}) (*mcp.CallToolResult, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format data: %v", err)), nil
	}

	// Combine summary with JSON data
	fullText := fmt.Sprintf("%s\n\n```json\n%s\n```", summary, string(jsonData))
	return mcp.NewToolResultText(fullText), nil
}

// formatCurrency converts Currency to dollar string for display
func formatCurrency(cents tekmetric.Currency) string {
	dollars := float64(cents) / 100.0
	return fmt.Sprintf("$%.2f", dollars)
}
