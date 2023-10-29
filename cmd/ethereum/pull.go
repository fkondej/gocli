package ethereum

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fkondej/gocli/clients/etherscan"
	"github.com/fkondej/gocli/ethutils"
	"github.com/fkondej/gocli/types"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type smartContract struct {
	Name       string
	HexAddress string
	Network    types.ETHNetwork
}

type PullArgs struct {
	*EthereumArgs
	SmartContracts   []smartContract
	EthereumEndpoint string
	EtherscanAPIKey  string
}

var pullArgs = PullArgs{
	SmartContracts: []smartContract{
		// {Name: "MultisigControl", Network: types.ETHMainnet, HexAddress: "0x164D322B2377C0fdDB73Cd32f24e972A7d9C72F9"},
		{Name: "ERC20Bridge", Network: types.ETHMainnet, HexAddress: "0xCd403f722b76366f7d609842C589906ca051310f"},
		// {Name: "ERC20AssetPool", Network: types.ETHMainnet, HexAddress: "0xA226E2A13e07e750EfBD2E5839C5c3Be80fE7D4d"},
		// {Name: "StakingBridge", Network: types.ETHMainnet, HexAddress: "0x195064D33f09e0c42cF98E665D9506e0dC17de68"},
		// {Name: "VestingBridge", Network: types.ETHMainnet, HexAddress: "0x23d1bFE8fA50a167816fBD79D7932577c06011f4"},
		// {Name: "ERC20BridgeRestricted", Network: types.ETHMainnet, HexAddress: "0x23872549cE10B40e31D6577e0A920088B0E0666a"},
		// {Name: "MultisigControl", Network: types.ETHMainnet, HexAddress: "0xDD2df0E7583ff2acfed5e49Df4a424129cA9B58F"},
		// {Name: "ClaimCodes", Network: types.ETHMainnet, HexAddress: "0x0ee1fb382caf98e86e97e51f9f42f8b4654020f3"},
	},
}

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Get Smart Contracts bytecode and source code from Ethereum Network to local",
	Long:  `Get Smart Contracts bytecode and source code`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunPull(pullArgs); err != nil {
			pullArgs.Logger.Error("Error", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	pullArgs.EthereumArgs = &ethereumArgs

	EthereumCmd.AddCommand(pullCmd)
	pullCmd.PersistentFlags().StringVar(&pullArgs.EthereumEndpoint, "ethereum-endpoint", "https://mainnet.infura.io/v3/57a0d24661374c0ab7136d0ffca35297", "Ethereum endpoint")
	pullCmd.PersistentFlags().StringVar(&pullArgs.EtherscanAPIKey, "etherscan-api-key", "", "Etherscan API Key - it will work 25x faster with key provided")

}

func RunPull(args PullArgs) error {
	// Get clients
	ethClient, err := ethclient.DialContext(context.Background(), args.EthereumEndpoint)
	if err != nil {
		return fmt.Errorf("failed to get ethereum client for %s, %w", args.EthereumEndpoint, err)
	}
	etherescanClient, err := etherscan.NewEtherscanClientWithNetwork(types.ETHMainnet, args.EtherscanAPIKey)
	if err != nil {
		return fmt.Errorf("failed to create etherscan cleint, %w", err)
	}

	// pull data for all smart contracts
	for _, contract := range args.SmartContracts {
		dir := filepath.Join("smartcontracts", contract.Name)

		if err := ethutils.PullAndStoreSmartContractImmutableData(
			contract.HexAddress, contract.Network, contract.Name, dir, ethClient, etherescanClient,
		); err != nil {
			return err
		}
		fmt.Printf(" - Downloaded %s(%s) from '%s' and stored in %s\n",
			contract.Name, contract.HexAddress, contract.Network, dir)
	}
	return nil
}
