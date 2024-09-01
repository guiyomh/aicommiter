package interfaces

import (
	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

// ConfigLoader is an interface for loading configuration
type ConfigLoader interface {
	Load() (entities.Config, error)
}
