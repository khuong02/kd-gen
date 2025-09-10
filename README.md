# KD Gen

A powerful Go code generation tool for creating type-safe enums from YAML configuration files.

## Overview

**KD Gen** is a CLI tool that generates Go enum types with useful helper methods from YAML configuration files.
It helps you create type-safe enums with string representations, parsing functions, and more.

## ✨ Features
### V1
- [x] Generate type-safe enums (int, string, uint64, …)
- [x] JSON support (MarshalJSON, UnmarshalJSON)
- [x] Database/sql integration (Scan, Value for DB storage)

### V2
- [ ] MCP Tools → expose enum generation as interactive tools:

### Tools
| Tool                    | Category | Description                                                                |
| ----------------------- | -------- | -------------------------------------------------------------------------- |
| `list_methods`          | Enum     | List all supported enum generation methods with descriptions.              |
| `show_config_structure` | Config   | Show the YAML structure for a valid enum configuration file.               |
| `example_config`        | Config   | Generate an example enum config file for a given type (e.g., string, int). |
| `validate_config`       | Config   | Validate an enum YAML config against schema and supported types.           |
| `generate_enum`         | Codegen  | Generate Go enum code from a provided YAML configuration.                  |
| `explain_enum`          | Docs     | Explain what a given enum config means and how it maps to generated code.  |
| `diff_config`           | Utility  | Compare two enum YAML config files and show the differences.               |
| `format_config`         | Utility  | Format and normalize enum YAML config files.                               |


## Go Version

This project is developed using **Go 1.25.0**.
It should work with other recent Go versions as well.

## Libraries Used

- [github.com/dave/jennifer](https://github.com/dave/jennifer) – Code generation for Go
- [github.com/spf13/cobra](https://github.com/spf13/cobra) – CLI command framework
- [github.com/spf13/viper](https://github.com/spf13/viper) – Configuration management
- [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) – Decode generic map values to Go structs
- [github.com/shopspring/decimal](https://github.com/shopspring/decimal) – Arbitrary-precision fixed-point decimal numbers
- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) - A Go implementation of the Model Context Protocol (MCP), enabling seamless integration between LLM applications and external data sources and tools.

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/khuong02/kd-gen.git
cd kd-gen

# Build the binary
go build -o ./build/kd-gen main.go

# Optional: Move to a directory in your PATH
mv ./build/kd-gen /usr/local/bin/
```

### Using docker
```bash
# Build image
docker build -t kd-gen .

# Run
docker run --rm -v $(pwd):/app -w /app kd-gen \
  enum gen --package core --output ./example/enum/enum_gen.go --config ./example/enum/enum.yaml
```

### Using Go Install

```bash
go install github.com/khuong02/kd-gen@latest
```

### Using Homebrew

You can install **KD Gen** via [Homebrew](https://brew.sh/) using a custom tap:

```bash
# Add the tap
brew tap khuong02/tap https://github.com/khuong02/homebrew-tap

# Install kd-gen
brew install kd-gen

## Usage

### Basic Usage

Create a YAML configuration file defining your enums:

```yaml
enums:
  - name: Lang
    type: string
    values:
      - name: LangEnglish
        display: "English"
        code: "en"
      - name: LangVietnamese
        display: "Vietnamese"
        code: "vi"
  - name: LogLevel
    type: int
    values:
      - name: Debug
      - name: Info
      - name: Warn
      - name: Error
      - name: Fatal
```

Then run the tool to generate the Go code:

```bash
kd-gen enum gen --package core --output ./example/enum/enum_gen.go --config ./example/enum/enum.yaml
```

### Sample Run

```bash
# Create a config file
cat > enum.yaml << EOF
enums:
  - name: Status
    type: string
    values:
      - name: StatusActive
        display: "active"
        code: 1
      - name: StatusInactive
        display: "inactive"
        code: 0
EOF

# Generate the enum code
kd-gen enum gen --package core --output status_enum.go --config enum.yaml
```

The generated file will contain:

- Status type definition
- Constants for each enum value
- `String()` method for string representation
- `Parse` function to convert strings to enum values
- Map of codes
- Slice of all values

---

## Generated Code Features

For each enum defined in your YAML file, the tool generates:

1. A type definition (`string`, `int`, etc.)
2. Constants for each enum value
3. A map of code values (if provided)
4. A slice containing all enum values
5. A `String()` method for string representation
6. A `Parse` function to convert strings to enum values
7. A `Normalize()` method for string enums

---

## Configuration File Format

```yaml
enums:
  - name: EnumName      # Name of the enum type
    methods:
    - string
    - parse
    - normalize
    - json
    - sql
    type: string|int    # Type of the enum (string, int, etc.)
    values:
      - name: ValueName     # Name of the enum constant
        display: "Display"  # String representation (optional)
        code: value         # Optional "Code" field → supports multiple types:
```

---
## 🔌 MCP Integration

This project can also be used as an [MCP server](https://modelcontextprotocol.io/) so Claude Desktop can call its tools directly.

👉 See [MCP_GUIDE.md](./MCP_GUIDE.md) for full setup instructions.

---

## 📚 References

| Project      | Purpose                                           |
|--------------|---------------------------------------------------|
| [mcp-go](https://github.com/mark3labs/mcp-go) | Go implementation of the Model Context Protocol. |
| [mcp-grafana](https://github.com/grafana/mcp-grafana/tree/main) | Example MCP server integrating with Grafana. |

---

## License

[MIT License](LICENSE)
