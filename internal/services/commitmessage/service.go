package commitmessage

import (
	"context"
	"strings"

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
	content, err := s.generator.Generate(ctx, prompt)
	if err != nil {
		return domain.CommitMessage{}, err
	}

	header, body, footer := s.parseCommitMessage(content)

	return domain.CommitMessage{
		Header: header,
		Body:   body,
		Footer: footer,
	}, nil
}

// parseCommitMessage splits the content into header, body, and footer
func (*CommitMessageService) parseCommitMessage(content string) (header, body, footer string) {
	// Diviser le contenu en lignes
	lines := strings.Split(strings.Trim(content, "`\n"), "\n")

	// Identifier l'index des parties
	header = lines[0] // Première ligne comme header
	bodyLines := []string{}
	footerLines := []string{}
	footerStart := false

	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "BREAKING CHANGE:") || strings.HasPrefix(line, "Fixes #") {
			footerStart = true
		}
		if footerStart {
			footerLines = append(footerLines, line)
		} else {
			bodyLines = append(bodyLines, line)
		}
	}

	body = strings.Join(bodyLines, "\n")
	footer = strings.Join(footerLines, "\n")

	return header, strings.TrimSpace(body), strings.TrimSpace(footer)
}
