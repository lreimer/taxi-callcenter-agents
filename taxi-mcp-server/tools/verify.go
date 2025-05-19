package tools

import (
	"context"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func verifyAddress(s *server.MCPServer) {
	// create a new MCP tool for address verification
	projectsTool := mcp.NewTool("verify_address",
		mcp.WithDescription("Verify that a given address consisting of street and city actually exists."),
		mcp.WithOpenWorldHintAnnotation(true),
		mcp.WithString("street",
			mcp.Description("The name and house number of the street to lookup and verify"),
			mcp.Required(),
		),
		mcp.WithString("city",
			mcp.Description("The name of the city to lookup and verify"),
			mcp.Required(),
		),
	)

	// add the tool to the server
	s.AddTool(projectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		street := request.Params.Arguments["street"].(string)
		city := request.Params.Arguments["city"].(string)

		if street == "" || city == "" {
			return mcp.NewToolResultError("street and city are required"), nil
		}

		if strings.ToLower(city) != "rosenheim" {
			log.Printf("Invalid address %s, %s\n", street, city)
			return mcp.NewToolResultText("Invalid."), nil
		}

		log.Printf("Valid address %s, %s\n", street, city)
		return mcp.NewToolResultText("Valid."), nil
	})
}
