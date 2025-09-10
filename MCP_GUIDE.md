# MCP Integration Guide for Claude
This project can be used as an [MCP server](https://modelcontextprotocol.io/) so Claude Desktop can call its tools directly.

## 1. Configure Claude Desktop

**Claude Desktop looks for a file named claude_desktop_config.json.**

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`

- Linux: `~/.config/Claude/claude_desktop_config.json`

- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

**Edit the file and add the MCP server:**

```json
{
  "mcpServers": {
    "kd-gen": {
      "command": "kd-gen",
      "args": ["mcp"]
    }
  }
}
```
## 2. Restart Claude

**Quit and restart Claude Desktop.
If everything is configured correctly, you’ll see logs indicating:**

```text
[kd-gen] Server started and connected successfully
```

## 3. Using the MCP Tools

**You can now use MCP tools exposed by this server.
For example:**

> Prompt:
“List all supported enum generate methods from the kd-gen MCP server.”

Claude will call the list_method tool and return available methods with descriptions.
