# Go CLI to replace Bash scripting

This repository is a showcase to replace bash scripting with Golang.

This is particularly useful for projects created in Go:
- Minimise maintenance time
  - resolving any breaking changes can be done in a couple of minutes
- Reduce creation and testing time for scripts interacting with a developed product
  - access to all code from your main project, specifically API and other structs,
- Looks and feels like bash
  - thanks to `go run main.go ...`, you can quickly edit a file and rerun the script,
- You can create more scripts and automate more
  - thanks to: 1. reduced maintenance time, 2. golang (strictly typed, modularised, imports directly from GitHub), 3. go-tightly coupled with the most changing software, i.e. main product

There are also disadvantages:
- Some simple things in Golang require more code than in, e.g. Python.

## Setup

1. Install Go `v1.21.1+`
2. Clone this repo

## Run

```bash
cd gocli
go run main.go --help
```

## Implementation

All scripts are organised into groups, and there are no subgroups.

Usage:
```bash
go run main.go [group-name] [scrpt-name] <script-arguments>
```

All CLI scripts are in [cmd](./cmd) directory, while logic is organised into modules in other directories.

Use `--help` to discover all commands and options. Docs might not be up to date.

## Examples

### Random Server Data

[code](cmd/random/serverdata.go)

```bash
go run main.go random server-data --json
```

Example result:
```bash
{
    "about-url": "https://en.wikipedia.org/wiki/All_Saints_Church,_West_Dulwich",
    "avatar": "https://www.gravatar.com/avatar/lk32lxrvaks3odqww4ix6bv8zwfvch5o?d=identicon",
    "country-code": "SE",
    "server-name": "green-bee"
}
```

### Get info about ERC20 token from Ethereum

[code](cmd/ethereum/erc20token.go)

```bash
go run main.go ethereum erc20-token --address 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
```

Example result:
```bash
Info about token 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
 - Name: USD Coin
 - Symbol: USDC
 - Decimals: 6
 - TotalSupply: 23,028,924,373.681125 USDC
```

### Pull code, binary and generate go bindings of existing Smart Contract

[code](cmd/ethereum/pull.go)

```bash
go run main.go ethereum pull
```

Example result:
```bash
- ERC20Bridge: updated Source Code: smartcontracts/ERC20Bridge/ERC20_Bridge_Logic.sol
- ERC20Bridge: updated Binary Code: smartcontracts/ERC20Bridge/ERC20Bridge.bin
- ERC20Bridge: updated Hash: smartcontracts/ERC20Bridge/hash.txt
- ERC20Bridge: updated ABI: smartcontracts/ERC20Bridge/abi.json
- ERC20Bridge: updated Go Bindings: smartcontracts/ERC20Bridge/ERC20Bridge.go
 - Downloaded ERC20Bridge(0xCd403f722b76366f7d609842C589906ca051310f) from 'mainnet' and stored in smartcontracts/ERC20Bridge
```

### Test of Ethereum Endpoints

[code](cmd/ethereum/testEndpoints.go)

```bash
go run main.go ethereum test-endpoints --rpc https://eth-rpc.gateway.pokt.network,https://mainnet.infura.io/v3/57a0d24661374c0ab7136d0ffca35297
```

Example result:
```bash
Results:
- https://eth-rpc.gateway.pokt.network: 18455271
- https://mainnet.infura.io/v3/<API-KEY>: 18455271
Get events from block 18455261 for USDC: 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
https://eth-rpc.gateway.pokt.network [18455261]: 1
https://eth-rpc.gateway.pokt.network [18455262]: 6
https://eth-rpc.gateway.pokt.network [18455263]: 7
https://eth-rpc.gateway.pokt.network [18455264]: 7
https://eth-rpc.gateway.pokt.network [18455265]: 1
https://eth-rpc.gateway.pokt.network [18455266]: 3
https://eth-rpc.gateway.pokt.network [18455267]: 1
https://eth-rpc.gateway.pokt.network [18455268]: 3
https://eth-rpc.gateway.pokt.network [18455269]: 5
https://eth-rpc.gateway.pokt.network [18455270]: 6
https://eth-rpc.gateway.pokt.network [18455271]: 3
https://mainnet.infura.io/v3/<API-KEY> [18455261]: 1
https://mainnet.infura.io/v3/<API-KEY> [18455262]: 6
https://mainnet.infura.io/v3/<API-KEY> [18455263]: 7
https://mainnet.infura.io/v3/<API-KEY> [18455264]: 7
https://mainnet.infura.io/v3/<API-KEY> [18455265]: 1
https://mainnet.infura.io/v3/<API-KEY> [18455266]: 3
https://mainnet.infura.io/v3/<API-KEY> [18455267]: 1
https://mainnet.infura.io/v3/<API-KEY> [18455268]: 3
https://mainnet.infura.io/v3/<API-KEY> [18455269]: 5
https://mainnet.infura.io/v3/<API-KEY> [18455270]: 6
https://mainnet.infura.io/v3/<API-KEY> [18455271]: 3
```
