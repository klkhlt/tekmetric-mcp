package tools

import (
	"github.com/beetlebugorg/tekmetric-mcp/internal/config"
	"github.com/beetlebugorg/tekmetric-mcp/internal/tekmetric"
	"github.com/mark3labs/mcp-go/server"
	"log/slog"
)

// Registry holds all tools and provides registration methods
type Registry struct {
	client *tekmetric.Client
	config *config.Config
	logger *slog.Logger
}

// NewRegistry creates a new tool registry
func NewRegistry(client *tekmetric.Client, cfg *config.Config, logger *slog.Logger) *Registry {
	return &Registry{
		client: client,
		config: cfg,
		logger: logger,
	}
}

// RegisterAll registers all tools with the MCP server
func (r *Registry) RegisterAll(s *server.MCPServer) {
	r.RegisterShopTools(s)
	r.RegisterCustomerTools(s)
	r.RegisterVehicleTools(s)
	r.RegisterRepairOrderTools(s)
	r.RegisterJobTools(s)
	r.RegisterAppointmentTools(s)
	r.RegisterEmployeeTools(s)
	r.RegisterInventoryTools(s)

	r.logger.Info("registered all MCP tools")
}
