package commitmessage

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/adapters/external_services"
	"github.com/guiyomh/aicommitter/internal/interfaces"
)

// NewAdapter creates a new adapter based on the adapter type
func NewAdapter(
	ctx context.Context,
	adapterType AdapterType,
	promptbuilder interfaces.PromptBuilder,
	apiKey string,
) (interfaces.MessageGenerator, error) {
	switch adapterType {
	case GoogleGenAI:
		return external_services.NewGoogleGenAIAdapter(ctx, apiKey, promptbuilder)
	case Ollama:
		return external_services.NewOllamaAdapter(promptbuilder)
	default:
		return external_services.NewOllamaAdapter(promptbuilder)
	}
}
