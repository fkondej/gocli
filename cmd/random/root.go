package random

import (
	rootCmd "github.com/fkondej/gocli/cmd"
	"github.com/spf13/cobra"
)

type RandomArgs struct {
	*rootCmd.RootArgs
}

var randomArgs RandomArgs

var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random names",
	Long:  `Random names`,
}

func init() {
	randomArgs.RootArgs = &rootCmd.Args
}
