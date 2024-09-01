package cmd

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/internal/domain/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/domain/services/promptbuilder"
	"github.com/guiyomh/aicommitter/internal/domain/usecases/analyze"
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

			config := cmd.Context().Value(entities.ConfigKey).(entities.Config)
			ctx := context.Background()

			usecase := analyze.New()
			input := analyze.Input{
				AdapterType:         commitmessage.GoogleGenAI,
				CommitSpecification: promptbuilder.Conventional,
				CommitOptions:       buildAnalyzeOptions(cmd),
				ApiKey:              config.APIKey,
			}
			usecase.Analyze(ctx, input)
		},
	}

	cmd.Flags().StringP("scope", "s", "", "Force the scope of the commit")
	cmd.Flags().StringP("type", "t", "", "Force the type of the commit")
	cmd.Flags().StringP("issue", "i", "", "Add the issue number of the commit")
	cmd.Flags().StringP("language", "l", "", "Force the language of the message")
	cmd.Flags().StringP("adapter", "a", "google_genai", "The adapter to use to generate the commit message")

	return cmd
}

func buildAnalyzeOptions(cmd *cobra.Command) []commitmessage.Option {
	options := make([]commitmessage.Option, 0)

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
