package workspace

import (
	"path/filepath"

	"github.com/IAmRiteshKoushik/smriti/internal/server"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func loadWorkspace(path string) (*Workspace, error) {
	k := koanf.New(".")

	if err := k.Load(
		file.Provider(path),
		toml.Parser(),
	); err != nil {
		return nil, err
	}

	var w Workspace

	if err := k.Unmarshal("", &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func Load(root string) (*LoadedWorkspace, error) {
	workspace, err := loadWorkspace(filepath.Join(root, "mcpx.toml"))
	if err != nil {
		return nil, err
	}

	loadedWorkspace := &LoadedWorkspace{
		Root:      root,
		Workspace: *workspace,
		Servers:   make(map[string]server.Server),
	}

	files, err := filepath.Glob(filepath.Join(root, "servers", ".toml"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		srv, err := server.Load(f)
		if err != nil {
			return nil, err
		}

		loadedWorkspace.Servers[srv.Name] = *srv
	}
	return loadedWorkspace, nil
}
