package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IAmRiteshKoushik/smriti/internal/common"
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

	// Request building
	req, err := mcp.BuildRequest(sc)
	if err != nil {
		return err
	}

	// Client and session building
	client := &transport.Client{
		URL:     serverCfg.Transport.URL,
		Headers: serverCfg.Headers,
		HTTP:    http.DefaultClient,
	}
	session := &transport.Session{
		Client: client,
	}

	// Initialize session
	if err := session.Initialize(ctx); err != nil {
		return err
	}

	// Execute request after setting up session
	resp, err := session.Execute(ctx, req)
	if err != nil {
		return err
	}

	data, _ := common.ExtractSSEData(resp)
	var pretty bytes.Buffer
	_ = json.Indent(&pretty, data, "", " ") // unhandled error

	fmt.Println(pretty.String())

	return nil
}
