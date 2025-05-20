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

## Cloud Run Deployment

```bash
# make sure you enabled these Google APIs
gcloud services enable cloudbuild.googleapis.com artifactregistry.googleapis.com
gcloud services enable run.googleapis.com 

# deploy the MCP server first
gcloud run deploy taxi-mcp-server --source=taxi-mcp-server/ \
  --region=europe-north1 \
  --port=8001 --allow-unauthenticated \
  --set-env-vars=BASE_URL=https://taxi-mcp-server-343509396461.europe-north1.run.app

# now we deploy the agent with UI
export GOOGLE_API_KEY=<insert here>

# we use plain gcloud CLI to perform the deployment
gcloud run deploy simple-taxi-agent --source=. \
  --region=europe-north1 \
  --port=8000 --allow-unauthenticated \
  --set-env-vars=MCP_SERVER_URL=https://taxi-mcp-server-343509396461.europe-north1.run.app/sse,GOOGLE_API_KEY=$GOOGLE_API_KEY,GOOGLE_GENAI_USE_VERTEXAI=FALSE

# there is a `adk deploy cloud_run command`
# however, it lacks the possibility to provide ENV variables
# also, the --with_ui option does not seem to have effect

# if you need to debug have a look at the logs
gcloud run services logs read taxi-mcp-server --region=europe-north1
gcloud run services logs read simple-taxi-agent --region=europe-north1

# use these commands to delete the workloads
gcloud run services list
gcloud run services delete taxi-mcp-server --async --region=europe-north1
gcloud run services delete simple-taxi-agent --async --region=europe-north1
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the GPL open source license, read the `LICENSE` file for details.