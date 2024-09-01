package commitmessage

import (
	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

// Option is a function that modifies a Prompt.
type Option = func(prompt entities.Prompt) entities.Prompt

// WithScope returns a CommitMessageOption that sets the scope of a Prompt.
func WithScope(scope string) Option {
	return func(prompt entities.Prompt) entities.Prompt {
		prompt.Scope = scope
		return prompt
	}
}

// WithType returns a CommitMessageOption that sets the type of a Prompt.
func WithType(commitType string) Option {
	return func(prompt entities.Prompt) entities.Prompt {
		prompt.CommitType = commitType
		return prompt
	}
}

// WithIssue returns a CommitMessageOption that sets the issue number of a Prompt.
func WithIssue(issue string) Option {
	return func(prompt entities.Prompt) entities.Prompt {
		prompt.IssueNumber = issue
		return prompt
	}
}

// WithLanguage returns a CommitMessageOption that sets the language of a Prompt.
func WithLanguage(language string) Option {
	return func(prompt entities.Prompt) entities.Prompt {
		prompt.Language = language
		return prompt
	}
}
