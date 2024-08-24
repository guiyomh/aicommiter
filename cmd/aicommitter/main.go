package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/guiyomh/aicommitter/internal/adapters/gemini"
	"github.com/guiyomh/aicommitter/internal/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/services/gitdiff"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	model := "gemini-1.5-flash"
	ctx := context.Background()
	adapter, err := gemini.NewGoogleGenAIAdapter(ctx, apiKey, model)
	if err != nil {
		log.Fatal(err)
	}

	service := commitmessage.NewCommitMessageService(adapter)

	diffGenerator := gitdiff.NewGitDiffGenerator()
	diff, err := diffGenerator.GenerateDiff()

	if err != nil {
		log.Fatal(err)
	}

	commitMessage, err := service.CreateCommitMessage(
		ctx,
		diff,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generated Commit Message:\n%s\n%s\n%s\n", commitMessage.Header, commitMessage.Body, commitMessage.Footer)
}
