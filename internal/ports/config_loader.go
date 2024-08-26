package ports

import "github.com/guiyomh/aicommitter/internal/domain"

// ConfigLoader is an interface for loading configuration
type ConfigLoader interface {
	LoadConfig() (domain.Config, error)
}
