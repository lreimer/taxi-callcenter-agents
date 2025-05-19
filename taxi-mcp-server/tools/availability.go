package tools

import (
	"context"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func checkAvailability(s *server.MCPServer) {
	// create a new MCP tool for Taxi availability
	projectsTool := mcp.NewTool("check_availability",
		mcp.WithDescription("Check if a taxi is available at a given address."),
		mcp.WithOpenWorldHintAnnotation(true),
		mcp.WithString("street",
			mcp.Description("The name and house number to check for availability"),
			mcp.Required(),
		),
		mcp.WithString("city",
			mcp.Description("The name of the city to check for availability"),
			mcp.Required(),
		),
	)

	// add the tool to the server
	s.AddTool(projectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the location parameter from the request
		street := request.Params.Arguments["street"].(string)
		city := request.Params.Arguments["city"].(string)

		if street == "" || city == "" {
			return mcp.NewToolResultError("street and city are required"), nil
		}

		if strings.Contains(strings.ToLower(street), "burgfried") {
			log.Printf("Taxi available for %s, %s\n", street, city)
			return mcp.NewToolResultText("There is a Taxi available for your address. It is 5 minutes away."), nil
		} else {
			log.Printf("No Taxi available for %s, %s\n", street, city)
			return mcp.NewToolResultText("Sorry, at the moment there is no Taxi available for your location."), nil
		}
	})
}
