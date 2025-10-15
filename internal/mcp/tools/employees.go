package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterEmployeeTools registers all employee-related tools
func (r *Registry) RegisterEmployeeTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("employees",
			mcp.WithDescription("Get employee/technician details by ID."),
			mcp.WithNumber("id",
				mcp.Required(),
				mcp.Description("Employee ID"),
			),
		),
		r.handleEmployees,
	)

	r.logger.Debug("registered employee tools")
}

// handleEmployees gets a specific employee by ID
func (r *Registry) handleEmployees(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	id, errResult := requireFloatArg(arguments, "id")
	if errResult != nil {
		return errResult, nil
	}

	employee, err := r.client.GetEmployee(ctx, id)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get employee: %v", err)), nil
	}

	return formatJSON(employee)
}
