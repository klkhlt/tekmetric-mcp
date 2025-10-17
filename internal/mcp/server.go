// Package mcp provides the Model Context Protocol server implementation for Tekmetric.
// It handles MCP server initialization, authentication, and tool registration.
package mcp

import (
	"context"
	"log/slog"

	"github.com/beetlebugorg/tekmetric-mcp/internal/config"
	"github.com/beetlebugorg/tekmetric-mcp/internal/mcp/analysis"
	"github.com/beetlebugorg/tekmetric-mcp/internal/mcp/tools"
	"github.com/beetlebugorg/tekmetric-mcp/pkg/tekmetric"
	"github.com/mark3labs/mcp-go/server"
)

// Server represents the MCP server for Tekmetric.
// It wraps an MCP server instance and provides integration with the Tekmetric API.
type Server struct {
	server *server.MCPServer  // The underlying MCP server
	client *tekmetric.Client  // Authenticated Tekmetric API client
	config *config.Config     // Server configuration
	logger *slog.Logger       // Structured logger
}

// NewServer creates a new MCP server instance.
// It initializes the Tekmetric API client, creates the MCP server,
// and registers all available tools.
//
// The server is configured to communicate via stdio (standard input/output)
// which is the standard communication method for MCP servers.
//
// Parameters:
//   - cfg: Server configuration including Tekmetric API credentials
//   - logger: Structured logger for server operations
//
// Returns:
//   - *Server: Configured MCP server ready to start
//   - error: Any error during initialization
func NewServer(cfg *config.Config, logger *slog.Logger) (*Server, error) {
	// Create Tekmetric API client with OAuth2 authentication
	tekmetricClient := tekmetric.NewClient(&cfg.Tekmetric, logger)

	// Create MCP server instance
	// Tools are automatically enabled when registered via AddTool
	mcpServer := server.NewMCPServer(
		cfg.Server.Name,
		cfg.Server.Version,
		server.WithLogging(),
	)

	s := &Server{
		server: mcpServer,
		client: tekmetricClient,
		config: cfg,
		logger: logger,
	}

	// Register all Tekmetric tools (shops, customers, vehicles, etc.)
	toolRegistry := tools.NewRegistry(tekmetricClient, cfg, logger)
	toolRegistry.RegisterAll(mcpServer)

	// Register analysis tools
	analysisRegistry := analysis.NewRegistry(tekmetricClient, cfg, logger)
	analysisRegistry.Register(analysis.NewVehicleServiceAnalysis(tekmetricClient, cfg, logger))
	analysisRegistry.RegisterAll(mcpServer)

	return s, nil
}

// Start starts the MCP server and begins listening for requests.
// It first authenticates with the Tekmetric API to obtain an access token,
// then starts serving MCP requests via stdio.
//
// This is a blocking call that runs until the context is cancelled or
// an error occurs. The server communicates with Claude Desktop via
// standard input/output streams.
//
// Parameters:
//   - ctx: Context for server lifecycle management
//
// Returns:
//   - error: Any error during authentication or server operation
func (s *Server) Start(ctx context.Context) error {
	// Authenticate with Tekmetric API before starting server
	// This obtains an OAuth2 access token for API requests
	if err := s.client.Authenticate(ctx); err != nil {
		return err
	}

	s.logger.Info("MCP server starting",
		"name", s.config.Server.Name,
		"version", s.config.Server.Version)

	// Start serving MCP requests via stdio
	// This blocks until the server is stopped or encounters an error
	return server.ServeStdio(s.server)
}

