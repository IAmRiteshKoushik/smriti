package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	URL     string
	Headers map[string]string
	HTTP    *http.Client
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (c *Client) Send(ctx context.Context, payload any, sessionID string) (*Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json, text/event-stream")

	if sessionID != "" {
		httpReq.Header.Set("Mcp-Session-Id", sessionID)
	}

	for k, v := range c.Headers {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: httpResp.StatusCode,
		Headers:    httpResp.Header,
		Body:       body,
	}, nil
}
