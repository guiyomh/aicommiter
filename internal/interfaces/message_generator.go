package interfaces

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

// MessageGenerator is an interface that generates commit messages.
type MessageGenerator interface {
	Generate(ctx context.Context, prompt entities.Prompt) (string, error)
}
