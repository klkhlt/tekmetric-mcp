package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterJobTools registers all job-related tools
func (r *Registry) RegisterJobTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("jobs",
			mcp.WithDescription("Get job details by ID. Jobs are work items/services on repair orders. Usually accessed via repair orders."),
			mcp.WithNumber("id",
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		r.handleJobs,
	)

	r.logger.Debug("registered job tools")
}

// handleJobs gets a specific job by ID
func (r *Registry) handleJobs(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	ctx := context.Background()

	id, errResult := requireFloatArg(arguments, "id")
	if errResult != nil {
		return errResult, nil
	}

	job, err := r.client.GetJob(ctx, id)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get job: %v", err)), nil
	}

	return formatJSON(job)
}
