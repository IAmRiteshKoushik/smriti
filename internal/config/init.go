package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// Indicates whether the project has been initialized or not
	InitFlagFilename = "init"
)

// Represents the initialization status for a project directory
type ProjectInitFloat struct {
	Initialized bool `json:"initialized"`
}

// Check if the initialization dialog box should be shown for current directory
func ShouldShowInitDialog() (bool, error) {
	if cfg == nil {
		return false, fmt.Errorf("config not loaded")
	}

	flagFilePath := filepath.Join(cfg.Data.Directory, InitFlagFilename)

	// Check if the flag file exists
	_, err := os.Stat(flagFilePath)
	if err == nil {
		// File exists, don't show the dialog
		return false, nil
	}

	// If the error is not "file not found", return the error
	if !os.IsNotExist(err) {
		return false, fmt.Errorf("failed to check init flag file: %w", err)
	}

	// File does not exist, show the dialog box
	return true, nil
}

func MarkProjectInitialized() error {
	if cfg == nil {
		return fmt.Errorf("config not loaded")
	}

	flagFilePath := filepath.Join(cfg.Data.Directory, InitFlagFilename)

	file, err := os.Create(flagFilePath)
	if err != nil {
		return fmt.Errorf("failed to create init flag file: %w", err)
	}
	defer file.Close()

	return nil
}
