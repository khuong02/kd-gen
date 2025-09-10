package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ------------------------------
// Tool: generate_enum
// ------------------------------
type GenerateEnumParams struct {
	Config string `json:"config" jsonschema:"description=YAML or JSON config content"`
}

func generateEnum(ctx context.Context, args GenerateEnumParams) (string, error) {
	// TODO: integrate with kd-gen core logic
	return "// generated Go enum code here", nil
}

var GenerateEnum = MustTool(
	"generate_enum",
	"Generate Go enum code from a config file.",
	generateEnum,
	mcp.WithTitleAnnotation("Generate enum"),
)

func AddGenerateEnumTool(s *server.MCPServer) {
	GenerateEnum.Register(s)
}

// ------------------------------
// Tool: explain_enum
// ------------------------------
type ExplainEnumParams struct {
	Name string `json:"name" jsonschema:"description=Enum name to explain"`
}

func explainEnum(ctx context.Context, args ExplainEnumParams) (string, error) {
	// TODO: lookup enum and return usage example
	return "Enum '" + args.Name + "' provides methods: String(), Parse(), Values(), IsValid() ...", nil
}

var ExplainEnum = MustTool(
	"explain_enum",
	"Explain how to use a generated enum with examples.",
	explainEnum,
	mcp.WithTitleAnnotation("Explain enum"),
)

func AddExplainEnumTool(s *server.MCPServer) {
	ExplainEnum.Register(s)
}

// ------------------------------
// Tool: diff_config
// ------------------------------
type DiffConfigParams struct {
	ConfigA string `json:"configA"`
	ConfigB string `json:"configB"`
}

func diffConfig(ctx context.Context, args DiffConfigParams) (string, error) {
	// TODO: compute a YAML diff
	return "No differences found", nil
}

var DiffConfig = MustTool(
	"diff_config",
	"Compare two config files and return differences.",
	diffConfig,
	mcp.WithTitleAnnotation("Diff config"),
)

func AddDiffConfigTool(s *server.MCPServer) {
	DiffConfig.Register(s)
}

// ------------------------------
// Tool: format_config
// ------------------------------
type FormatConfigParams struct {
	Config string `json:"config"`
}

func formatConfig(ctx context.Context, args FormatConfigParams) (string, error) {
	// TODO: pretty-print YAML
	return args.Config, nil
}

var FormatConfig = MustTool(
	"format_config",
	"Format and normalize a config file.",
	formatConfig,
	mcp.WithTitleAnnotation("Format config"),
)

func AddFormatConfigTool(s *server.MCPServer) {
	FormatConfig.Register(s)
}
