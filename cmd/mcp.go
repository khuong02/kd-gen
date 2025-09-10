package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/khuong02/kd-gen/tools"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Run kd-gen as an MCP server",
	Long:  "Start the kd-gen MCP server so tools like Claude Desktop can call kd-gen programmatically.",
	RunE:  kdMcp,
}

func init() {
	RootCmd.AddCommand(mcpCmd)
}

func kdMcp(cmd *cobra.Command, args []string) error {
	return run("stdio")
}

func maybeAddTools(s *server.MCPServer, tf func(*server.MCPServer)) {
	tf(s)
}

func newServer() *server.MCPServer {
	s := server.NewMCPServer("mcp-grafana", version, server.WithInstructions(`
	This server provides access to kd-gen functionality.

	Available Capabilities:
	- Enum Generation: Generate Go enums from YAML configuration files.
	`))

	maybeAddTools(s, tools.AddListMethodsTool)
	maybeAddTools(s, tools.AddShowConfigStructureTool)
	return s
}

func run(transport string) error {
	s := newServer()

	// Create a context that will be cancelled on shutdown
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	// Handle shutdown signals
	go func() {
		<-sigChan
		slog.Info("Received shutdown signal")
		//cancel()

		// For stdio, close stdin to unblock the Listen call
		if transport == "stdio" {
			if err := os.Stdin.Close(); err != nil {
				fmt.Printf("Warning: failed to close stdin: %v\n", err)
			}
		}
	}()

	switch transport {
	case "stdio":

		slog.Info("Starting KD Gen MCP server transport=stdio", "version", version)

		err := server.ServeStdio(s)
		if err != nil && errors.Is(err, context.Canceled) {
			return fmt.Errorf("server error: %w", err)
		}
		return nil

	case "sse", "streamable-http":
		// TODO: implement when needed
		return fmt.Errorf("transport %q not yet implemented", transport)

	default:
		return fmt.Errorf("invalid transport type %q (must be 'stdio', 'sse', or 'streamable-http')", transport)
	}
}
