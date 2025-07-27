package lib

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	RootPath string `json:"root_path"`
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
