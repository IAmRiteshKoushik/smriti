package workspace

import (
	"github.com/IAmRiteshKoushik/smriti/internal/server"
)

type Workspace struct {
	Name          string `koanf:"name"`
	DefaultServer string `koanf:"default_server"`
}

type LoadedWorkspace struct {
	Root      string
	Workspace Workspace
	Servers   map[string]server.Server
}
