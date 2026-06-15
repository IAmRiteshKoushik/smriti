package server

import (
	"github.com/IAmRiteshKoushik/smriti/internal/common"
)

type TransportConfig struct {
	Type string `koanf:"type"`
	URL  string `koanf:"url"`
}

type Server struct {
	Name      string            `koanf:"name"`
	Transport TransportConfig   `koanf:"transport"`
	Headers   map[string]string `koanf:"headers"`
}

func Load(path string) (*Server, error) {
	return common.Load[Server](path)
}
