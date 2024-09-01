package commitmessage

import (
	"context"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/internal/interfaces"
)

// Service is a service that generates commit messages.
type Service struct {
	generator interfaces.MessageGenerator
}

// New creates a new instance of CommitMessageService.
func New(generator interfaces.MessageGenerator) *Service {
	return &Service{generator: generator}
}

// CreateCommitMessage generates a commit message based on the given diff.
func (s *Service) CreateCommitMessage(
	ctx context.Context,
	diff entities.Diff,
	options ...Option,
) (entities.CommitMessage, error) {
	prompt := entities.Prompt{
		Diff: diff,
	}
	for _, option := range options {
		prompt = option(prompt)
	}
	content, err := s.generator.Generate(ctx, prompt)
	if err != nil {
		return entities.CommitMessage{}, err
	}

	header, body, footer := s.parseCommitMessage(content)

	return entities.CommitMessage{
		Header: header,
		Body:   body,
		Footer: footer,
	}, nil
}

// parseCommitMessage splits the content into header, body, and footer
func (*Service) parseCommitMessage(content string) (header, body, footer string) {
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
