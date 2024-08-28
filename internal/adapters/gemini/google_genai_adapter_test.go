package gemini

import (
	"testing"

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
