package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/spf13/viper"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (*Loader) Load() (entities.Config, error) {
	var config entities.Config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	configPath := filepath.Join(homeDir, ".config", "aicommitter.yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil
}
