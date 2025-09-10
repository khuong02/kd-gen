package tools

import (
	"context"
	"errors"
	"strings"

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
	methods:
    - String|Parse|Normalize|JSON|SQL
    type: string|int    # Type of the enum (string, int, etc.)
    values:
      - name: ValueName     # Name of the enum constant
        display: "Display"  # Optional string representation
        code: value         # Optional "Code" field → supports multiple types:
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

type ExampleConfigParams struct {
	Type string `json:"type" jsonschema:"description=Enum type (string, int, etc.)"`
}

// exampleConfig returns a YAML config example depending on args.Type.
func exampleConfig(ctx context.Context, args ExampleConfigParams) (string, error) {
	t := strings.ToLower(args.Type)

	// Exact match
	if yamlExample, ok := yamlExamples[t]; ok {
		return yamlExample, nil
	}

	// If type belongs to int-family → reuse int example with replaced type
	if intFamily[t] {
		return strings.Replace(yamlExamples["int"], "type: int", "type: "+t, 1), nil
	}

	// Unsupported type
	return "", errors.New("do not support this type: " + t)
}

var ExampleConfig = MustTool(
	"example_config",
	"Show an example enum config file for a given type.",
	exampleConfig,
	mcp.WithTitleAnnotation("Example config (YAML)"),
	mcp.WithIdempotentHintAnnotation(true),
	mcp.WithReadOnlyHintAnnotation(true),
)

func AddExampleConfigTool(s *server.MCPServer) {
	ExampleConfig.Register(s)
}

// ------------------------------
// Tool: validate_config
// ------------------------------
type ValidateConfigParams struct {
	Config string `json:"config" jsonschema:"description=YAML or JSON config content"`
}

type ValidationResult struct {
	Valid  bool     `json:"valid"`
	Errors []string `json:"errors,omitempty"`
}

func validateConfig(ctx context.Context, args ValidateConfigParams) (ValidationResult, error) {
	// TODO: parse and validate YAML/JSON
	return ValidationResult{
		Valid:  true,
		Errors: nil,
	}, nil
}

var ValidateConfig = MustTool(
	"validate_config",
	"Validate a config file for correctness.",
	validateConfig,
	mcp.WithTitleAnnotation("Validate config"),
)

func AddValidateConfigTool(s *server.MCPServer) {
	ValidateConfig.Register(s)
}
