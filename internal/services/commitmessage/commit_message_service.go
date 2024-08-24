package commitmessage

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/ports"
)

// CommitMessageService is a service that generates commit messages.
type CommitMessageService struct {
	generator ports.MessageGenerator
}

// NewCommitMessageService creates a new instance of CommitMessageService.
func NewCommitMessageService(generator ports.MessageGenerator) *CommitMessageService {
	return &CommitMessageService{generator: generator}
}

// CreateCommitMessage generates a commit message based on the given diff.
func (s *CommitMessageService) CreateCommitMessage(
	ctx context.Context,
	diff domain.Diff,
	options ...CommitMessageOption,
) (domain.CommitMessage, error) {
	prompt := domain.Prompt{
		Diff: diff,
	}
	for _, option := range options {
		prompt = option(prompt)
	}
	return s.generator.Generate(ctx, prompt)
}
