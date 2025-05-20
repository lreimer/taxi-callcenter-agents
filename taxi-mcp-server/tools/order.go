package tools

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func orderTaxi(s *server.MCPServer) {
	projectsTool := mcp.NewTool("order_taxi",
		mcp.WithDescription("Order a taxi for given passenger details and pickup location (street, city)."),
		mcp.WithOpenWorldHintAnnotation(true),
		mcp.WithString("passenger_name",
			mcp.Description("The name of the passenger to order a taxi for"),
			mcp.Required(),
		),
		mcp.WithString("passenger_phone",
			mcp.Description("The phone number of the passenger to order a taxi for"),
			mcp.Required(),
		),
		mcp.WithString("street",
			mcp.Description("The name and house number to pick up the passenger"),
			mcp.Required(),
		),
		mcp.WithString("city",
			mcp.Description("The name of the city to pick up the passenger"),
			mcp.Required(),
		),
		mcp.WithString("pickup_time",
			mcp.Description("The time when the taxi should pick up the passenger"),
			mcp.DefaultString("asap"),
		),
		mcp.WithString("pickup_details",
			mcp.Description("Additional details for the pickup location"),
			mcp.DefaultString(""),
		),
	)

	// add the tool to the server
	s.AddTool(projectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the location parameter from the request
		passengerName := request.Params.Arguments["passenger_name"].(string)
		passengerPhone := request.Params.Arguments["passenger_phone"].(string)
		street := request.Params.Arguments["street"].(string)
		city := request.Params.Arguments["city"].(string)
		pickupTime := request.Params.Arguments["pickup_time"].(string)
		pickupDetails := request.Params.Arguments["pickup_details"].(string)

		log.Printf("ordering taxi for %s (%s) at %s, %s\n", passengerName, passengerPhone, street, city)
		log.Printf("Pickup time: %s\n", pickupTime)
		log.Printf("Pickup details: %s\n", pickupDetails)

		return mcp.NewToolResultText("Taxi ordered successfully!"), nil
	})
}
