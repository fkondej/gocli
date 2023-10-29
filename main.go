package main

import (
	rootCmd "github.com/fkondej/gocli/cmd"
	"github.com/fkondej/gocli/cmd/ethereum"
	"github.com/fkondej/gocli/cmd/random"
	"github.com/fkondej/gocli/cmd/secrets"
	"github.com/fkondej/gocli/cmd/version"
)

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd.RootCmd.AddCommand(ethereum.EthereumCmd)
	rootCmd.RootCmd.AddCommand(random.RandomCmd)
	rootCmd.RootCmd.AddCommand(secrets.SecretsCmd)
	rootCmd.RootCmd.AddCommand(version.VersionCmd)
}
