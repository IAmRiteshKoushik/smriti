package transport

import (
	"context"
	"encoding/json"
	"fmt"
)

type Session struct {
	Client    *Client
	SessionID string
}

func (s *Session) Initialize(ctx context.Context) error {
	initReq := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "initialize",
		"params": map[string]any{
			"protocolVersion": "2025-03-26",
			"capabilities":    map[string]any{},
			"clientInfo": map[string]any{
				"name":    "smriti",
				"version": "0.1.0",
			},
		},
	}

	resp, err := s.Client.Send(ctx, initReq, "")
	if err != nil {
		return err
	}

	sessionID := resp.Headers.Get("Mcp-Session-Id")
	if sessionID == "" {
		sessionID = resp.Headers.Get("MCP-Session-Id")
	}

	if sessionID == "" {
		return fmt.Errorf("server did not return session id")
	}

	s.SessionID = sessionID

	var parsed map[string]any
	if err := json.Unmarshal(resp.Body, &parsed); err != nil {
		return err
	}

	if _, ok := parsed["error"]; ok {
		return fmt.Errorf("initialize failed: %s", string(resp.Body))
	}
	return s.SendInitialized(ctx)
}

func (s *Session) SendInitialized(ctx context.Context) error {
	req := map[string]any{
		"jsonrpc": "2.0",
		"method":  "notifications/initialized",
	}

	_, err := s.Client.Send(ctx, req, s.SessionID)

	return err
}

func (s *Session) Execute(ctx context.Context, payload any) ([]byte, error) {
	resp, err := s.Client.Send(ctx, payload, s.SessionID)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
