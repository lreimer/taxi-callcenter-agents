# Taxi Callcenter Agents

A call center agent that can help with ordering a taxi.
This demo repository uses Google ADK for now. Maybe CrewAI later.

## Project Setup

```bash
# initialize project with deps and tools
uv init
uv add google-adk
uv tool install google-adk

# run and open the demo console
# choose the simple_taxi_agent
# the agent only supports voice
adk web
open http://localhost:8000
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