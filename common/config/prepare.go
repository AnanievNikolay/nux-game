package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func PrepareConfig(
	configPath string,
) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(fmt.Sprintf("file.Close: %s", err))
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var cfg Config
	if err = json.Unmarshal(
		bytes,
		&cfg,
	); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &cfg, nil
}
