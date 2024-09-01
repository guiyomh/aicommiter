package commitmessage

import "fmt"

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
