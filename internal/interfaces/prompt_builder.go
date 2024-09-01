package interfaces

import (
	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

type PromptBuilder interface {
	Build(entities.Prompt) string
}
