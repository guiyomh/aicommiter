package gemini

import (
	"context"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/ports"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type GoogleGenAIAdapter struct {
	llm           *googleai.GoogleAI
	promptBuilder ports.PromptBuilder
}

func NewGoogleGenAIAdapter(
	ctx context.Context,
	apiKey string,
	builder ports.PromptBuilder,
) (ports.MessageGenerator, error) {
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &GoogleGenAIAdapter{
		llm:           llm,
		promptBuilder: builder,
	}, nil
}

func (g *GoogleGenAIAdapter) Generate(ctx context.Context, prompt domain.Prompt) (domain.CommitMessage, error) {
	fullPrompt := g.promptBuilder.Build(prompt)

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		g.llm,
		fullPrompt,
	)
	if err != nil {
		return domain.CommitMessage{}, err
	}

	// Décomposer le contenu en header, body et footer
	header, body, footer := parseCommitMessage(completion)

	// Logique simple pour construire le message à partir de la réponse
	return domain.CommitMessage{
		Header: header,
		Body:   body,
		Footer: footer,
	}, nil
}

// parseCommitMessage splits the content into header, body, and footer
func parseCommitMessage(content string) (header, body, footer string) {
	// Diviser le contenu en lignes
	lines := strings.Split(strings.Trim(content, "`\n"), "\n")

	// Identifier l'index des parties
	header = lines[0] // Première ligne comme header
	bodyLines := []string{}
	footerLines := []string{}
	footerStart := false

	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "BREAKING CHANGE:") || strings.HasPrefix(line, "Fixes #") {
			footerStart = true
		}
		if footerStart {
			footerLines = append(footerLines, line)
		} else {
			bodyLines = append(bodyLines, line)
		}
	}

	body = strings.Join(bodyLines, "\n")
	footer = strings.Join(footerLines, "\n")

	return header, strings.TrimSpace(body), strings.TrimSpace(footer)
}
