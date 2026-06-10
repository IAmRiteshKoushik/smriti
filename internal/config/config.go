package config

import "fmt"

// Global configuration instance
var cfg *Config

const (
	defaultDataDirectory = ".smriti"
	defaultLogLevel      = "info"
	appName              = "smriti"
)

type Config struct {
	Data       Data                 `json:"data"`
	WorkingDir string               `json:"wd"`
	MCPServers map[string]MCPServer `json:"mcpServers,omitempty"`
	Debug      bool                 `json:"debug"`
	TUI        TUIConfig            `json:"tui"`
	Shell      ShellConfig          `json:"shell"`
}

// TUI specific configuration
type TUIConfig struct {
	Theme string `json:"theme,omitempty"`
}

// Shell specific configuration
type ShellConfig struct {
	Path string   `json:"path,omitempty"`
	Args []string `json:"args,omitempty"`
}

// Data defines the storage configuration
type Data struct {
	Directory string `json:"directory,omitempty"`
}

// MCP specific configuration
type MCPType string

const (
	MCPStdio          MCPType = "stdio"
	MCPSse            MCPType = "sse"
	MCPStreamableHttp MCPType = "streamableHttp"
)

type MCPServer struct {
	Command string            `json:"command"`
	Env     []string          `json:"env"`
	Args    []string          `json:"args"`
	Type    MCPType           `json:"type"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func Load(workingDir string, debug bool) (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		WorkingDir: workingDir,
		MCPServers: make(map[string]MCPServer),
	}

	setDefaults(debug)

	// if err := readConfig(); err != nil {
	// 	return cfg, err
	// }

	return nil, nil
}

func configureKoanf() {
}

func setDefaults(debug bool) {
}

func readConfig(err error) error {
	if err == nil {
		return nil
	}

	// if _, ok := err.(); ok {
	// 	return nil
	// }
	return fmt.Errorf("failed to read config: %w", err)
}

func applyDefaultValues() {
	for k, v := range cfg.MCPServers {
		if v.Type == "" {
			v.Type = MCPStreamableHttp
			cfg.MCPServers[k] = v
		}
	}
}
