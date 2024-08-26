package cmd

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/adapters/gemini"
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/services/gitdiff"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command

func NewAnalyzeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "analyze",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, _ []string) {
			config := cmd.Context().Value(domain.ConfigKey).(domain.Config)

			model := "gemini-1.5-flash"
			ctx := context.Background()
			adapter, err := gemini.NewGoogleGenAIAdapter(ctx, config.APIKey, model)
			if err != nil {
				cobra.CheckErr(err)
			}

			service := commitmessage.NewCommitMessageService(adapter)

			diffGenerator := gitdiff.NewGitDiffGenerator()
			diff, err := diffGenerator.GenerateDiff()

			if err != nil {
				cobra.CheckErr(err)
			}

			commitMessage, err := service.CreateCommitMessage(
				ctx,
				diff,
			)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Printf("Generated Commit Message:\n%s\n%s\n%s\n", commitMessage.Header, commitMessage.Body, commitMessage.Footer)
		},
	}
}
