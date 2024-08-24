// Package gitdiff contains the service for git diff
package gitdiff

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/guiyomh/aicommitter/internal/domain"
)

// GitDiffGenerator implements the DiffGenerator interface to generate diffs using Git.
type GitDiffGenerator struct{}

// NewGitDiffGenerator crée une nouvelle instance de GitDiffGenerator.
func NewGitDiffGenerator() *GitDiffGenerator {
	return &GitDiffGenerator{}
}

// GenerateDiff executes the git diff --staged command and returns the diff.
func (*GitDiffGenerator) GenerateDiff() (domain.Diff, error) {
	cmd := exec.Command("git", "diff", "--staged", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum")
	cmd.Env = append(cmd.Env, "GIT_PAGER=cat")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return domain.Diff{}, err
	}

	diffContent := out.String()

	if len(diffContent) == 0 {
		return domain.Diff{}, fmt.Errorf("no diff content")
	}

	return domain.Diff{Content: diffContent}, nil
}
