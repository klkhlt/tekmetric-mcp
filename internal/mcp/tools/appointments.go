package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterAppointmentTools registers all appointment-related tools
func (r *Registry) RegisterAppointmentTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("appointments",
			mcp.WithDescription("Get appointment details by ID. Appointments are scheduled service visits."),
			mcp.WithNumber("id",
				mcp.Required(),
				mcp.Description("Appointment ID"),
			),
		),
		r.handleAppointments,
	)

	r.logger.Debug("registered appointment tools")
}

// handleAppointments gets a specific appointment by ID
func (r *Registry) handleAppointments(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	id, errResult := requireFloatArg(arguments, "id")
	if errResult != nil {
		return errResult, nil
	}

	appointment, err := r.client.GetAppointment(ctx, id)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get appointment: %v", err)), nil
	}

	return formatJSON(appointment)
}
