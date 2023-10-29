package ethereum

import (
	rootCmd "github.com/fkondej/gocli/cmd"
	"github.com/spf13/cobra"
)

type EthereumArgs struct {
	*rootCmd.RootArgs
}

var ethereumArgs EthereumArgs

var EthereumCmd = &cobra.Command{
	Use:   "ethereum",
	Short: "Various functionality for the Ethereum",
	Long:  `Various functionality for the Ethereum`,
}

func init() {
	ethereumArgs.RootArgs = &rootCmd.Args
}
