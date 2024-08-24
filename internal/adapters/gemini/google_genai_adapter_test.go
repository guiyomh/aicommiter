package gemini

import (
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestParseCommitMessage(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    struct {
			header string
			body   string
			footer string
		}
	}{
		{
			name: "Message complet avec header, body, et footer",
			content: `feat(app): Rename root component

This commit renames the root component from "app-root" to "my-app-root".

This change is part of a larger effort to consistently name components in the application.

BREAKING CHANGE: The root component name has been changed.
Fixes #123`,
			want: struct {
				header string
				body   string
				footer string
			}{
				header: "feat(app): Rename root component",
				body: `This commit renames the root component from "app-root" to "my-app-root".

This change is part of a larger effort to consistently name components in the application.`,
				footer: "BREAKING CHANGE: The root component name has been changed.\nFixes #123",
			},
		},
		{
			name: "Wrapped message complet avec header, body, et footer",
			content: "```\n" + `feat(app): Rename root component

This commit renames the root component from "app-root" to "my-app-root".

This change is part of a larger effort to consistently name components in the application.

BREAKING CHANGE: The root component name has been changed.
Fixes #123` + "\n```",
			want: struct {
				header string
				body   string
				footer string
			}{
				header: "feat(app): Rename root component",
				body: `This commit renames the root component from "app-root" to "my-app-root".

This change is part of a larger effort to consistently name components in the application.`,
				footer: "BREAKING CHANGE: The root component name has been changed.\nFixes #123",
			},
		},
		{
			name: "Message sans footer",
			content: `fix(auth): Improve login performance

Refactored the login service to reduce latency during authentication.
The improvements will result in faster login times for all users.`,
			want: struct {
				header string
				body   string
				footer string
			}{
				header: "fix(auth): Improve login performance",
				body: `Refactored the login service to reduce latency during authentication.
The improvements will result in faster login times for all users.`,
				footer: "",
			},
		},
		{
			name:    "Message avec seulement un header",
			content: `chore: update dependencies`,
			want: struct {
				header string
				body   string
				footer string
			}{
				header: "chore: update dependencies",
				body:   "",
				footer: "",
			},
		},
		{
			name: "Message avec un footer sans body",
			content: `refactor(utils): optimize helper functions

BREAKING CHANGE: Refactored utility functions, some of which have changed signatures.`,
			want: struct {
				header string
				body   string
				footer string
			}{
				header: "refactor(utils): optimize helper functions",
				body:   "",
				footer: "BREAKING CHANGE: Refactored utility functions, some of which have changed signatures.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotBody, gotFooter := parseCommitMessage(tt.content)

			assert.Equal(t, tt.want.header, gotHeader, "parseCommitMessage() header = %v, want %v", gotHeader, tt.want.header)
			assert.Equal(t, tt.want.body, gotBody, "parseCommitMessage() body = %v, want %v", gotBody, tt.want.body)
			assert.Equal(t, tt.want.footer, gotFooter, "parseCommitMessage() footer = %v, want %v", gotFooter, tt.want.footer)
		})
	}
}

func TestBuildPrompt(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			result := buildPrompt(tt.prompt)
			assert.Equal(t, tt.expected, result)
		})
	}
}
