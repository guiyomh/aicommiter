package entities

// Prompt represents the data required to generate a commit message, including diff and metadata.
type Prompt struct {
	Diff        Diff
	CommitType  string
	Scope       string
	IssueNumber string
	Language    string
}
