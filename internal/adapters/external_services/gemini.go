package external_services

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/internal/interfaces"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type GoogleGenAIAdapter struct {
	llm           *googleai.GoogleAI
	promptBuilder interfaces.PromptBuilder
}

func NewGoogleGenAIAdapter(
	ctx context.Context,
	apiKey string,
	builder interfaces.PromptBuilder,
) (interfaces.MessageGenerator, error) {
	llm, err := googleai.New(
		ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultModel("gemini-1.5-pro"),
	)
	if err != nil {
		return nil, err
	}

	return &GoogleGenAIAdapter{
		llm:           llm,
		promptBuilder: builder,
	}, nil
}

func (g *GoogleGenAIAdapter) Generate(ctx context.Context, prompt entities.Prompt) (string, error) {
	fullPrompt := g.promptBuilder.Build(prompt)

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, fullPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt.Diff.Content),
	}

	response, err := g.llm.GenerateContent(
		ctx,
		content,
		llms.WithTemperature(0.5),
		llms.WithTopP(0.1),
	)
	if err != nil {
		log.Error().Err(err).Any("response", response).Msg("failed to generate content")
		return "", err
	}

	var completion string
	for _, choice := range response.Choices {
		completion += choice.Content
	}

	return completion, nil

}
