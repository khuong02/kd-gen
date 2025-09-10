package tools

import (
	"context"

	"github.com/khuong02/kd-gen/pkg/enum"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ListMethodsParams defines input params for the listMethods tool.
// Currently unused, but kept for future extensibility.
type ListMethodsParams struct {
	Query string `json:"query" jsonschema:"description=Optional filter string to match methods"`
}

// MethodInfo holds name and description for an enum generation method.

// listMethods returns all supported enum generation methods with descriptions.
func listMethods(ctx context.Context, args ListMethodsParams) (map[enum.Method]string, error) {
	return enum.MethodDescriptions, nil
}

// ListMethods is the MCP tool definition for listing enum generation methods.
var ListMethods = MustTool(
	"list_methods",
	"List all supported enum generation methods available in this server.",
	listMethods,
	mcp.WithTitleAnnotation("List methods"),
	mcp.WithIdempotentHintAnnotation(true),
	mcp.WithReadOnlyHintAnnotation(true),
)

// AddListMethodsTool registers the ListMethods MCP tool with the server.
func AddListMethodsTool(s *server.MCPServer) {
	ListMethods.Register(s)
}
