package common

import (
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load[T any](path string) (*T, error) {
	k := koanf.New(".")

	if err := k.Load(
		file.Provider(path),
		toml.Parser(),
	); err != nil {
		return nil, err
	}

	var out T

	if err := k.Unmarshal("", &out); err != nil {
		return nil, err
	}

	return &out, nil
}
