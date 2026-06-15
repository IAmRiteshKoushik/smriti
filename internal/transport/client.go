package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/IAmRiteshKoushik/smriti/internal/mcp"
)

type Client struct {
	URL     string
	Headers map[string]string
	HTTP    *http.Client
}

func (c *Client) Send(ctx context.Context, req *mcp.Request) ([]byte, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json, text/event-stream")

	for k, v := range c.Headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
