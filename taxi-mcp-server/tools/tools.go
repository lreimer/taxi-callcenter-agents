package tools

import "github.com/mark3labs/mcp-go/server"

// add all the Taxi call center tools to the MCP server instance
func AddTools(s *server.MCPServer) {
	verifyAddress(s)
	checkAvailability(s)
	orderTaxi(s)
}
