/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/adapters/gemini"
	"github.com/guiyomh/aicommitter/internal/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/services/gitdiff"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := cmd.Flags().GetString("apiKey")
		if err != nil {
			cobra.CheckErr(err)
		}
		model := "gemini-1.5-flash"
		ctx := context.Background()
		adapter, err := gemini.NewGoogleGenAIAdapter(ctx, apiKey, model)
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

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringP("apiKey", "k", "", "Google GenAI API Key")
	analyzeCmd.MarkFlagRequired("apiKey")

}
