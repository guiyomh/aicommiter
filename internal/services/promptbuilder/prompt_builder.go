package promptbuilder

import "github.com/guiyomh/aicommitter/internal/domain"

type DefaultPromptBuilder struct{}

func NewDefaultPromptBuilder() *DefaultPromptBuilder {
	return &DefaultPromptBuilder{}
}

func (*DefaultPromptBuilder) Build(p domain.Prompt) string {
	prompt := `Based on the following git diff --staged output,
generate a commit message following the Angular commit message convention.
The commit message should include:
A header with a <type>(<scope>): <short summary> format.
A detailed body explaining the motivation for the change.
If applicable, a footer with information on breaking changes, deprecations, or issue references.
`

	// Ajoutez des métadonnées supplémentaires si disponibles
	if p.CommitType != "" {
		prompt += "\nType: " + p.CommitType
	}
	if p.Scope != "" {
		prompt += "\nScope: " + p.Scope
	}
	if p.IssueNumber != "" {
		prompt += "\nIssue: #" + p.IssueNumber
	}
	if p.Language != "" {
		prompt += "\nGenerate the commit message with  " + p.Language + " language."
	}

	prompt += "\ngit diff --staged output:\n```\n"
	prompt += p.Diff.Content
	prompt += "\n```"

	return prompt
}
