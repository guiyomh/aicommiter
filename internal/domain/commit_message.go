package domain

// CommitMessage represents a commit message.
type CommitMessage struct {
	Header string
	Body   string
	Footer string
}

// DiffGenerator is the interface for generating a diff from source code.
type DiffGenerator interface {
	GenerateDiff() (Diff, error)
}

// Diff represents a diff generated from the source code.
type Diff struct {
	Content string
}

// Prompt represents the data required to generate a commit message, including diff and metadata.
type Prompt struct {
	Diff        Diff
	CommitType  string
	Scope       string
	IssueNumber string
	Language    string
}
