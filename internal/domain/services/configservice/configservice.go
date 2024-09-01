package configservice

import (
	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/internal/interfaces"
)

type ConfigService struct {
	loader interfaces.ConfigLoader
}

func NewConfigService(loader interfaces.ConfigLoader) *ConfigService {
	return &ConfigService{
		loader: loader,
	}
}

func (service *ConfigService) GetConfig() (entities.Config, error) {
	return service.loader.Load()
}
