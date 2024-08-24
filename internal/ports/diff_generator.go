package ports

import "github.com/guiyomh/aicommitter/internal/domain"

// DiffGenerator is an interface that generates diffs.
type DiffGenerator interface {
	GenerateDiff() (domain.Diff, error)
}
