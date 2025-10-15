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

// removeNullsAndEmpty recursively removes null, empty strings, empty slices, and zero values from maps
func removeNullsAndEmpty(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, val := range v {
			cleaned := removeNullsAndEmpty(val)
			// Skip null, empty strings, empty slices, and false booleans
			if cleaned == nil {
				continue
			}
			if str, ok := cleaned.(string); ok && str == "" {
				continue
			}
			if slice, ok := cleaned.([]interface{}); ok && len(slice) == 0 {
				continue
			}
			// Keep the value
			result[key] = cleaned
		}
		// Return nil if map is empty after cleaning
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		result := make([]interface{}, 0, len(v))
		for _, item := range v {
			cleaned := removeNullsAndEmpty(item)
			if cleaned != nil {
				result = append(result, cleaned)
			}
		}
		return result
	default:
		return v
	}
}

// cleanJSON converts data to JSON and back to remove omitempty fields, then filters nulls
func cleanJSON(data interface{}) (interface{}, error) {
	// First marshal to JSON to apply omitempty tags
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal back to interface{} to get a generic structure
	var generic interface{}
	if err := json.Unmarshal(jsonData, &generic); err != nil {
		return nil, err
	}

	// Remove nulls and empty values
	return removeNullsAndEmpty(generic), nil
}

// formatJSON marshals data to indented JSON and returns a tool result
func formatJSON(data interface{}) (*mcp.CallToolResult, error) {
	// Clean the data first
	cleaned, err := cleanJSON(data)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to clean data: %v", err)), nil
	}

	jsonData, err := json.MarshalIndent(cleaned, "", "  ")
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
	// Clean the data first
	cleaned, err := cleanJSON(data)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to clean data: %v", err)), nil
	}

	jsonData, err := json.MarshalIndent(cleaned, "", "  ")
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

// PaginatedResult wraps paginated data with metadata
type PaginatedResult[T any] struct {
	Data          []T    `json:"data"`
	TotalElements int    `json:"totalElements"`
	Returned      int    `json:"returned"`
	Truncated     bool   `json:"truncated,omitempty"`
	Message       string `json:"message,omitempty"`
}

// hasFinancialData checks if data contains financial fields
func hasFinancialData(resourceType string) bool {
	financialTypes := map[string]bool{
		"REPAIR ORDERS": true,
		"JOBS":          true,
	}
	return financialTypes[resourceType]
}

// formatPaginatedResultWithWarning creates a response with prominent truncation warnings
func formatPaginatedResultWithWarning[T any](data []T, totalElements int, returned int, maxResults int, resourceType string) (*mcp.CallToolResult, error) {
	response := map[string]interface{}{
		"data":          data,
		"totalElements": totalElements,
		"returned":      returned,
	}

	// ALWAYS add financial warning for financial data types
	if hasFinancialData(resourceType) {
		response["FINANCIAL_WARNING"] = "🚨 NOT FOR FINANCIAL REPORTING - Use Tekmetric's built-in reports 🚨"
	}

	// Add prominent warning if results were truncated
	if totalElements > maxResults {
		response["WARNING"] = fmt.Sprintf("⚠️ SHOWING ONLY %d OF %d %s ⚠️", returned, totalElements, resourceType)
		response["truncated"] = true
	} else if totalElements > returned {
		response["WARNING"] = fmt.Sprintf("⚠️ SHOWING %d OF %d %s ⚠️", returned, totalElements, resourceType)
		response["truncated"] = true
	}

	// formatJSON will handle the cleaning
	return formatJSON(response)
}
