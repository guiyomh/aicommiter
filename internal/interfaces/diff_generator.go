package interfaces

import (
	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

// DiffGenerator is an interface that generates diffs.
type DiffGenerator interface {
	GenerateDiff() (entities.Diff, error)
}
