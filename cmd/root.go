package cmd

import (
	"context"
	"os"

	"github.com/guiyomh/aicommitter/internal/adapters/configloader"
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/services/configservice"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aicommitter",
	Short: "A tool to generate commit messages using AI",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		// Initialize and load config using the service
		loader := configloader.NewConfigLoader()
		configService := configservice.NewConfigService(loader)
		config, err := configService.GetConfig()
		if err != nil {
			return err
		}
		ctx := context.WithValue(cmd.Context(), domain.ConfigKey, config)
		cmd.SetContext(ctx)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
