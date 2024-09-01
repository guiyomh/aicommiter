// Package gitdiff contains the service for git diff
package gitdiff

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

// GitDiffGenerator implements the DiffGenerator interface to generate diffs using Git.
type GitDiffGenerator struct{}

// New crée une nouvelle instance de GitDiffGenerator.
func New() *GitDiffGenerator {
	return &GitDiffGenerator{}
}

// GenerateDiff executes the git diff --staged command and returns the diff.
func (*GitDiffGenerator) GenerateDiff() (entities.Diff, error) {
	cmd := exec.Command("git", "diff", "--staged", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum")
	cmd.Env = append(cmd.Env, "GIT_PAGER=cat")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return entities.Diff{}, err
	}

	diffContent := out.String()

	if len(diffContent) == 0 {
		return entities.Diff{}, fmt.Errorf("no diff content")
	}

	return entities.Diff{Content: diffContent}, nil
}
