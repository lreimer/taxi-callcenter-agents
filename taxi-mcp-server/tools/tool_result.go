package tools

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func newToolResult(data any) (*mcp.CallToolResult, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to marshal response data.", err), err
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}
