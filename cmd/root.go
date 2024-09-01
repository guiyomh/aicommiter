package cmd

import (
	"context"

	"github.com/guiyomh/aicommitter/internal/adapters/config"
	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/internal/domain/services/configservice"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

const (
	normalVerbosity = 0
	infoVerbosity   = 1
	debugVerbosity  = 2
	traceVerbosity  = 3
)

func NewRootCmd() *cobra.Command {
	verboseCount := 0
	cmd := &cobra.Command{
		Use:   "aicommitter",
		Short: "A tool to generate commit messages using AI",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Initialize the logger
			initializeLogger(verboseCount, cmd)

			// Initialize and load config using the service
			loader := config.NewLoader()
			configService := configservice.NewConfigService(loader)
			config, err := configService.GetConfig()
			if err != nil {
				return err
			}
			ctx := context.WithValue(cmd.Context(), entities.ConfigKey, config)
			cmd.SetContext(ctx)

			return nil
		},
	}

	cmd.PersistentFlags().CountVarP(&verboseCount, "verbose", "v", "verbose output")

	return cmd
}

func initializeLogger(verboseCount int, cmd *cobra.Command) {
	level := zerolog.ErrorLevel
	switch verboseCount {
	case normalVerbosity:
		level = zerolog.ErrorLevel
	case infoVerbosity:
		level = zerolog.InfoLevel
	case debugVerbosity:
		level = zerolog.DebugLevel
	case traceVerbosity:
		level = zerolog.TraceLevel
	default:
		level = zerolog.TraceLevel
	}

	if level == zerolog.TraceLevel {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: cmd.OutOrStdout()}).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)
	log.Info().Msgf("Log level set to %s", level.String())
}
