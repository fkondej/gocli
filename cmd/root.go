package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RootArgs struct {
	Debug  bool
	Logger *zap.Logger
}

var Args RootArgs

var RootCmd = &cobra.Command{
	Use:   "gocli",
	Short: "CLI helper scripts for various things",
	Long: `CLI helper scripts for various things.
	An approach to replace bash scripting with Golang`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		cfg := zap.NewProductionConfig()
		if Args.Debug {
			cfg.Level.SetLevel(zap.DebugLevel)
		}
		// https://github.com/uber-go/zap/issues/584
		cfg.OutputPaths = []string{"stdout"}
		cfg.Encoding = "console"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		Args.Logger, err = cfg.Build()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to setup logger: %s\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run gocli '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.PersistentFlags().BoolVar(&Args.Debug, "debug", false, "Print debug logs")
}
