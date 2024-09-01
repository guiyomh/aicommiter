package analyze

import (
	"github.com/guiyomh/aicommitter/internal/domain/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/domain/services/promptbuilder"
)

type Input struct {
	AdapterType         commitmessage.AdapterType
	CommitSpecification promptbuilder.CommitSpecification
	CommitOptions       []commitmessage.Option
	ApiKey              string
}
