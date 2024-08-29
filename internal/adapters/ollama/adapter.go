package ollama

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/ports"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type OllamaAdapter struct {
	llm           *ollama.LLM
	promptBuilder ports.PromptBuilder
}

func NewOllamaAdapter(builder ports.PromptBuilder) (*OllamaAdapter, error) {
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		return nil, err
	}
	return &OllamaAdapter{
		llm:           llm,
		promptBuilder: builder,
	}, nil
}

func (o *OllamaAdapter) Generate(ctx context.Context, prompt domain.Prompt) (domain.CommitMessage, error) {
	fullPrompt := o.promptBuilder.Build(prompt)
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		o.llm,
		fullPrompt,
	)
	if err != nil {
		return domain.CommitMessage{}, err
	}

	fmt.Println(completion)

	return domain.CommitMessage{}, nil
}
