package transport

import (
	"context"

	"github.com/IAmRiteshKoushik/smriti/internal/mcp"
)

type Transport interface {
	Send(ctx context.Context, req *mcp.Request) ([]byte, error)
}
