package secrets

import (
	"fmt"
	"log"
	"os"

	"github.com/fkondej/gocli/generate/secrets"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type EthereumArgs struct {
	*SecretsArgs
	Passpharse string
	WalletPath string
}

var ethereumArgs EthereumArgs

var ethereumCmd = &cobra.Command{
	Use:   "ethereum",
	Short: "Create Ethereum wallet",
	Long:  `Create Ethereum wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunEthereum(ethereumArgs); err != nil {
			ethereumArgs.Logger.Error("Error", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	ethereumArgs.SecretsArgs = &secretsArgs
	SecretsCmd.AddCommand(ethereumCmd)

	ethereumCmd.PersistentFlags().StringVar(&ethereumArgs.WalletPath, "wallet-path", "", "File where to store generated wallet")
	if err := ethereumCmd.MarkPersistentFlagRequired("wallet-path"); err != nil {
		log.Fatalf("%v\n", err)
	}
	ethereumCmd.PersistentFlags().StringVar(&ethereumArgs.Passpharse, "passphrase", "", "Passphrase to access generated wallet")
	if err := ethereumCmd.MarkPersistentFlagRequired("passphrase"); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func RunEthereum(args EthereumArgs) error {
	var errMsg = "failed to genereate new ethereum wallet, %w"

	newWallet, err := secrets.GenerateEthereumWallet("", "", "")
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}

	if err := secrets.StoreEthereumWallet(newWallet.PrivateKey, args.Passpharse, args.WalletPath); err != nil {
		return fmt.Errorf(errMsg, err)
	}

	fmt.Printf("Generated new Ethereum wallet and stored in %s\n", args.WalletPath)
	fmt.Printf(" - Mnemonic: %s\n", newWallet.Mnemonic)
	fmt.Printf(" - Address: %s\n", newWallet.Address)
	fmt.Printf(" - Seed: %s\n", newWallet.Seed)
	fmt.Printf(" - PrivateKey: %s\n", newWallet.PrivateKey)

	return nil
}
