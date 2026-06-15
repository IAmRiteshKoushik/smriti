package mcp

import (
	"fmt"

	"github.com/IAmRiteshKoushik/smriti/internal/scenario"
)

type Request struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Method  string         `json:"method"`
	Params  ToolCallParams `json:"params"`
}

type ToolCallParams struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

func BuildRequest(s *scenario.Scenario) (*Request, error) {
	switch s.RequestType {
	case "tools/list":
		return &Request{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "tools/list",
		}, nil
	case "tools/call":
		return &Request{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "tools/list",
			Params: ToolCallParams{
				Name:      s.Tool,
				Arguments: s.Arguments,
			},
		}, nil
	}

	return nil, fmt.Errorf("unsupported request type %s", s.RequestType)
}
