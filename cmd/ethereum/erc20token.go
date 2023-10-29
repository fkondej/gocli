package ethereum

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fkondej/gocli/smartcontracts/ERC20Token"
	"github.com/fkondej/gocli/utils"
	"github.com/leekchan/accounting"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type ERC20TokenArgs struct {
	*EthereumArgs
	EthereumEndpoint string
	TokenHexAddress  string
}

var erc20TokenArgs ERC20TokenArgs

var erc20TokenCmd = &cobra.Command{
	Use:   "erc20-token",
	Short: "Get information about ERC20 token",
	Long:  `Get information about ERC20 token`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunERC20Token(erc20TokenArgs); err != nil {
			erc20TokenArgs.Logger.Error("Error", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	erc20TokenArgs.EthereumArgs = &ethereumArgs
	EthereumCmd.AddCommand(erc20TokenCmd)

	erc20TokenCmd.PersistentFlags().StringVar(&erc20TokenArgs.EthereumEndpoint, "ethereum-endpoint", "https://mainnet.infura.io/v3/57a0d24661374c0ab7136d0ffca35297", "Ethereum endpoint")

	erc20TokenCmd.PersistentFlags().StringVar(&erc20TokenArgs.TokenHexAddress, "address", "", "Token Hex address")
	if err := erc20TokenCmd.MarkPersistentFlagRequired("address"); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func RunERC20Token(args ERC20TokenArgs) error {
	var errMsg = fmt.Sprintf("token %s, from %s", args.TokenHexAddress, args.EthereumEndpoint)

	ethClient, err := ethclient.DialContext(context.Background(), args.EthereumEndpoint)
	if err != nil {
		return fmt.Errorf("failed to get ethereum client for %s, %w", args.EthereumEndpoint, err)
	}
	address := common.HexToAddress(args.TokenHexAddress)

	token, err := ERC20Token.NewERC20Token(address, ethClient)
	if err != nil {
		return fmt.Errorf("failed to get info about %s, %w", errMsg, err)
	}

	fmt.Printf("Info about token %s\n", args.TokenHexAddress)

	// Name
	name, err := token.Name(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get Name about %s, %w", errMsg, err)
	}
	fmt.Printf(" - Name: %s\n", name)

	// Symbol
	symbol, err := token.Symbol(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get Symbol about %s, %w", errMsg, err)
	}
	fmt.Printf(" - Symbol: %s\n", symbol)

	// Decimals
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get Decimals about %s, %w", errMsg, err)
	}
	fmt.Printf(" - Decimals: %d\n", decimals)

	// TotalSupply
	totalSupply, err := token.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get TotalSupply about %s, %w", errMsg, err)
	}
	supply := utils.TokenToFullTokens(totalSupply, decimals)
	ac := accounting.Accounting{Symbol: symbol, Precision: int(decimals), Format: "%v %s"}
	fmt.Printf(" - TotalSupply: %s\n", ac.FormatMoneyBigFloat(supply))

	return nil
}
