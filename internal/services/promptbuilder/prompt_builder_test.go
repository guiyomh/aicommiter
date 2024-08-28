package promptbuilder

import (
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPromptBuilder_Build(t *testing.T) {
	tests := []struct {
		name     string
		prompt   domain.Prompt
		expected string
	}{
		{
			name: "Prompt avec tous les champs",
			prompt: domain.Prompt{
				Diff:        domain.Diff{Content: "diff --git a/file.txt b/file.txt\nindex abc123..def456 100644\n--- a/file.txt\n+++ b/file.txt\n@@ -1,4 +1,4 @@\n-Hello World\n+Hello Go"},
				CommitType:  "feat",
				Scope:       "app",
				IssueNumber: "123",
			},
			expected: `Based on the following git diff --staged output, generate a commit message following the Angular commit message convention.
The commit message should include:
A header with a <type>(<scope>): <short summary> format.
A detailed body explaining the motivation for the change.
If applicable, a footer with information on breaking changes, deprecations, or issue references.

Type: feat
Scope: app
Issue: #123
git diff --staged output:` + "\n```\n" +
				`diff --git a/file.txt b/file.txt
index abc123..def456 100644
--- a/file.txt
+++ b/file.txt
@@ -1,4 +1,4 @@
-Hello World
+Hello Go` + "\n```",
		},
		{
			name: "Prompt sans scope et issue",
			prompt: domain.Prompt{
				Diff:       domain.Diff{Content: "diff --git a/file.txt b/file.txt\nindex abc123..def456 100644\n--- a/file.txt\n+++ b/file.txt\n@@ -1,4 +1,4 @@\n-Hello World\n+Hello Go"},
				CommitType: "fix",
			},
			expected: `Based on the following git diff --staged output, generate a commit message following the Angular commit message convention.
The commit message should include:
A header with a <type>(<scope>): <short summary> format.
A detailed body explaining the motivation for the change.
If applicable, a footer with information on breaking changes, deprecations, or issue references.

Type: fix
git diff --staged output:` + "\n```\n" +
				`diff --git a/file.txt b/file.txt
index abc123..def456 100644
--- a/file.txt
+++ b/file.txt
@@ -1,4 +1,4 @@
-Hello World
+Hello Go` + "\n```",
		},
		{
			name: "Prompt minimal",
			prompt: domain.Prompt{
				Diff: domain.Diff{Content: "diff --git a/file.txt b/file.txt\nindex abc123..def456 100644\n--- a/file.txt\n+++ b/file.txt\n@@ -1,4 +1,4 @@\n-Hello World\n+Hello Go"},
			},
			expected: `Based on the following git diff --staged output, generate a commit message following the Angular commit message convention.
The commit message should include:
A header with a <type>(<scope>): <short summary> format.
A detailed body explaining the motivation for the change.
If applicable, a footer with information on breaking changes, deprecations, or issue references.

git diff --staged output:` + "\n```\n" +
				`diff --git a/file.txt b/file.txt
index abc123..def456 100644
--- a/file.txt
+++ b/file.txt
@@ -1,4 +1,4 @@
-Hello World
+Hello Go` + "\n```",
		},
	}

	for _, tt := range tests {
		builder := NewDefaultPromptBuilder()
		t.Run(tt.name, func(t *testing.T) {
			result := builder.Build(tt.prompt)
			assert.Equal(t, tt.expected, result)
		})
	}
}
