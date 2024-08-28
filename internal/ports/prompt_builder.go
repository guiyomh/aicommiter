package ports

import "github.com/guiyomh/aicommitter/internal/domain"

type PromptBuilder interface {
	Build(domain.Prompt) string
}
