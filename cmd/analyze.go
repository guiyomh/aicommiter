package cmd

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/adapters/gemini"
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/services/gitdiff"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command

func NewAnalyzeCmd() *cobra.Command {
	cmd := &cobra.Command{
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
				log.Error().Err(err).Msg("Error while generating diff")
				cobra.CheckErr(err)
			}

			options := buildAnalyzeOptions(cmd)

			commitMessage, err := service.CreateCommitMessage(
				ctx,
				diff,
				options...,
			)
			if err != nil {
				log.Error().Err(err).Msg("Error while generating commit message")
				cobra.CheckErr(err)
			}
			fmt.Printf("Generated Commit Message:\n%s\n%s\n%s\n", commitMessage.Header, commitMessage.Body, commitMessage.Footer)
		},
	}

	cmd.Flags().StringP("scope", "s", "", "Force the scope of the commit")
	cmd.Flags().StringP("type", "t", "", "Force the type of the commit")
	cmd.Flags().StringP("issue", "i", "", "Add the issue number of the commit")
	cmd.Flags().StringP("language", "l", "", "Force the language of the message")

	return cmd
}

func buildAnalyzeOptions(cmd *cobra.Command) []commitmessage.CommitMessageOption {
	options := make([]commitmessage.CommitMessageOption, 0)

	scope := cmd.Flags().Lookup("scope")
	if scope != nil {
		options = append(options, commitmessage.WithScope(scope.Value.String()))
	}
	commitType := cmd.Flags().Lookup("type")
	if commitType != nil {
		options = append(options, commitmessage.WithType(commitType.Value.String()))
	}
	issue := cmd.Flags().Lookup("issue")
	if issue != nil {
		options = append(options, commitmessage.WithIssue(issue.Value.String()))
	}
	language := cmd.Flags().Lookup("language")
	if language != nil {
		options = append(options, commitmessage.WithLanguage(language.Value.String()))
	}

	return options
}
