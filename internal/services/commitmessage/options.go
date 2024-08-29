package commitmessage

import "github.com/guiyomh/aicommitter/internal/domain"

// CommitMessageOption is a function that modifies a Prompt.
type CommitMessageOption = func(prompt domain.Prompt) domain.Prompt

// WithScope returns a CommitMessageOption that sets the scope of a Prompt.
func WithScope(scope string) CommitMessageOption {
	return func(prompt domain.Prompt) domain.Prompt {
		prompt.Scope = scope
		return prompt
	}
}

// WithType returns a CommitMessageOption that sets the type of a Prompt.
func WithType(commitType string) CommitMessageOption {
	return func(prompt domain.Prompt) domain.Prompt {
		prompt.CommitType = commitType
		return prompt
	}
}

// WithIssue returns a CommitMessageOption that sets the issue number of a Prompt.
func WithIssue(issue string) CommitMessageOption {
	return func(prompt domain.Prompt) domain.Prompt {
		prompt.IssueNumber = issue
		return prompt
	}
}

// WithLanguage returns a CommitMessageOption that sets the language of a Prompt.
func WithLanguage(language string) CommitMessageOption {
	return func(prompt domain.Prompt) domain.Prompt {
		prompt.Language = language
		return prompt
	}
}
