package commitmessage

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/adapters/gemini"
	"github.com/guiyomh/aicommitter/internal/adapters/ollama"
	"github.com/guiyomh/aicommitter/internal/ports"
)

type AdapterType string

const (
	GoogleGenAI AdapterType = "google_genai"
	Ollama      AdapterType = "ollama"
)

func (a *AdapterType) FromString(adapter string) error {
	switch adapter {
	case string(GoogleGenAI):
		*a = GoogleGenAI
	case string(Ollama):
		*a = Ollama
	default:
		return fmt.Errorf(
			"unknown adapter type %s. Valid values are: %v",
			adapter,
			validAdapterTypes(),
		)
	}
	return nil
}

// validAdapterTypes returns the list of valid adapter types
func validAdapterTypes() []string {
	return []string{string(GoogleGenAI), string(Ollama)}
}

// CreateAdapter creates a new adapter based on the adapter type
func CreateAdapter(
	ctx context.Context,
	adapterType AdapterType,
	promptbuilder ports.PromptBuilder,
	apiKey string,
) (ports.MessageGenerator, error) {
	switch adapterType {
	case GoogleGenAI:
		return gemini.NewGoogleGenAIAdapter(ctx, apiKey, promptbuilder)
	case Ollama:
		return ollama.NewOllamaAdapter(promptbuilder)
	default:
		return ollama.NewOllamaAdapter(promptbuilder)
	}
}
