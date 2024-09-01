package entities

type ConfigKeyType string

const ConfigKey ConfigKeyType = "config"

type Config struct {
	APIKey string `mapstructure:"api_key"`
}
