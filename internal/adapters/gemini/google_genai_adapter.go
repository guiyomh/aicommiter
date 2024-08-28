package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/ports"
	"google.golang.org/api/option"
)

type GoogleGenAIAdapter struct {
	client        *genai.Client
	model         string
	promptBuilder ports.PromptBuilder
}

func NewGoogleGenAIAdapter(
	ctx context.Context,
	apiKey string,
	model string,
	builder ports.PromptBuilder,
) (ports.MessageGenerator, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &GoogleGenAIAdapter{
		client:        client,
		model:         model,
		promptBuilder: builder,
	}, nil
}

func (g *GoogleGenAIAdapter) Generate(ctx context.Context, prompt domain.Prompt) (domain.CommitMessage, error) {
	defer g.client.Close()

	fullPrompt := g.promptBuilder.Build(prompt)

	resp, err := g.client.GenerativeModel(g.model).GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		return domain.CommitMessage{}, err
	}

	// Suppose the response contains just one candidate
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return domain.CommitMessage{}, fmt.Errorf("no content generated")
	}

	// Récupérer la seule partie générée
	content := fmt.Sprint(resp.Candidates[0].Content.Parts[0])

	// Décomposer le contenu en header, body et footer
	header, body, footer := parseCommitMessage(content)

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
