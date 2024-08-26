package configloader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/spf13/viper"
)

type ConfigLoader struct{}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (*ConfigLoader) LoadConfig() (domain.Config, error) {
	var config domain.Config

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
