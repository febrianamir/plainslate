package lib

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type Config struct {
	UserHomeDir string

	RootPath string `json:"rootPath"`
}

func NewConfig() (*Config, error) {
	// Load config
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return &Config{}, nil // Return default config if file doesn't exist
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	// Set internal configs
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		LogError(context.Background(), err.Error(),
			zap.Strings("tags", []string{"config", "UserHomeDir"}),
		)
		return nil, err
	}
	cfg.UserHomeDir = userHomeDir

	LogInfo(context.Background(), "done set app config",
		zap.Any("config", cfg),
	)
	return &cfg, nil
}

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appConfigDir := filepath.Join(configDir, "plainslate")
	err = os.MkdirAll(appConfigDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(appConfigDir, "preferences.json"), nil
}
