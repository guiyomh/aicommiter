package ports

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/domain"
)

// MessageGenerator is an interface that generates commit messages.
type MessageGenerator interface {
	Generate(ctx context.Context, prompt domain.Prompt) (domain.CommitMessage, error)
}
