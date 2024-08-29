package configservice

import (
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/ports"
)

type ConfigService struct {
	loader ports.ConfigLoader
}

func NewConfigService(loader ports.ConfigLoader) *ConfigService {
	return &ConfigService{
		loader: loader,
	}
}

func (cs *ConfigService) GetConfig() (domain.Config, error) {
	return cs.loader.LoadConfig()
}
