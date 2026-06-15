package runner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IAmRiteshKoushik/smriti/internal/mcp"
	"github.com/IAmRiteshKoushik/smriti/internal/scenario"
	"github.com/IAmRiteshKoushik/smriti/internal/server"
	"github.com/IAmRiteshKoushik/smriti/internal/transport"
)

func RunScenario(ctx context.Context, serverCfg server.Server, scenarioPath string) error {
	sc, err := scenario.Load(scenarioPath)
	if err != nil {
		return err
	}

	if err := sc.Validate(); err != nil {
		return err
	}

	req, err := mcp.BuildRequest(sc)
	if err != nil {
		return err
	}

	client := &transport.Client{
		URL:     serverCfg.Transport.URL,
		Headers: serverCfg.Headers,
		HTTP:    http.DefaultClient,
	}

	resp, err := client.Send(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))

	return nil
}
