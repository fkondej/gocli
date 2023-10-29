package secrets

import (
	rootCmd "github.com/fkondej/gocli/cmd"
	"github.com/spf13/cobra"
)

type SecretsArgs struct {
	*rootCmd.RootArgs
}

var secretsArgs SecretsArgs

var SecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "General secrets tasks",
	Long:  `General secrets tasks`,
}

func init() {
	secretsArgs.RootArgs = &rootCmd.Args
}
