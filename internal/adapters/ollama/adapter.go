package ollama

import (
	"context"

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

func (o *OllamaAdapter) Generate(ctx context.Context, prompt domain.Prompt) (string, error) {
	fullPrompt := o.promptBuilder.Build(prompt)
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, fullPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt.Diff.Content),
	}

	response, err := o.llm.GenerateContent(ctx, content)
	if err != nil {
		return "", err
	}

	var completion string
	for _, choice := range response.Choices {
		completion += choice.Content
	}

	return completion, nil
}
