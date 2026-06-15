package scenario

import (
	"errors"
	"fmt"
)

type Scenario struct {
	Name        string         `koanf:"name"`
	Description string         `koanf:"description"`
	RequestType string         `koanf:"request_type"`
	Tool        string         `koanf:"tool"`
	Arguments   map[string]any `koanf:"arguments"`
}

func (s *Scenario) Validate() error {
	if s.Name == "" {
		return errors.New("Missing name in scenario")
	}

	switch s.RequestType {
	case "tools/list":
	// TODO: Not implemented
	case "tools/call":
	// TODO: Not implemented
	default:
		return fmt.Errorf("invalid request_type: %s", s.RequestType)
	}

	if s.RequestType == "tools/call" && s.Tool == "" {
		return errors.New("tool is required")
	}

	return nil
}
