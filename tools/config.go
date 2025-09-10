package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ShowConfigStructureParams defines input params for the showConfigStructure tool.
// Currently unused, but kept for future extensibility.
type ShowConfigStructureParams struct{}

// showConfigStructure returns an example YAML structure of the config file.
func showConfigStructure(ctx context.Context, args ShowConfigStructureParams) (string, error) {
	yamlExample := `enums:
  - name: EnumName      # Name of the enum type
    type: string|int    # Type of the enum (string, int, etc.)
    values:
      - name: ValueName     # Name of the enum constant
        display: "Display"  # Optional string representation
        code: value         # Optional associated value
`
	return yamlExample, nil
}

// ShowConfigStructure is the MCP tool definition.
var ShowConfigStructure = MustTool(
	"show_config_structure",
	"Show the YAML structure of the config file for generating enums.",
	showConfigStructure,
	mcp.WithTitleAnnotation("Show config structure (YAML)"),
	mcp.WithIdempotentHintAnnotation(true),
	mcp.WithReadOnlyHintAnnotation(true),
)

// AddShowConfigStructureTool registers the ShowConfigStructure MCP tool with the server.
func AddShowConfigStructureTool(s *server.MCPServer) {
	ShowConfigStructure.Register(s)
}
