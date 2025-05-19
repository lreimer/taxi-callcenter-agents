# Taxi Callcenter Agents

Intelligent taxi call center agent implementations that can process a taxi order.
This demo repository uses Google ADK for now.

## Project Setup

```bash
# initialize project with deps and tools
uv init
uv add google-adk
uv tool install google-adk

adk web
adk run weather_time_agent
```

## Taxi MCP Server

```bash
goreleaser build --snapshot --clean
goreleaser release --skip=publish --snapshot --clean
```

If you want to use the tool locally, e.g. with Claude Desktop, use the following
configuration for the MCP server.

```json
{
    "mcpServers": {
      "gcloud": {
        "command": "/Users/mario-leander.reimer/Applications/gcp-mcp-server",
        "args": ["--transport", "stdio"],
        "env": {
        }
      }
    }
}
```

Alternatively, you can use the MCP introspector for easy local development:
```bash
# as stdio binary
npx @modelcontextprotocol/inspector go run main.go

# as SSE server using 
go run main.go --transport sse
npx @modelcontextprotocol/inspector npx mcp-remote@next http://localhost:8001/sse
npx @modelcontextprotocol/inspector
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the GPL open source license, read the `LICENSE` file for details.