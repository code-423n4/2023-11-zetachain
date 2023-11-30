# ZetaChain audit details
- Total Prize Pool: $153,886 USDC 
  - HM awards: $125,286 USDC 
  - QA awards: $3,100 USDC 
  - Judge awards: $15,000 USDC
  - Lookout awards: $10,000 USDC 
  - Scout awards: $500 USDC 
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2023-11-zetachain/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts November 20, 2023 20:00 UTC
- Ends December 18, 2023 20:00 UTC 

❗️**Note for C4 wardens:** For this contest, analysis and gas optimizations are out of scope. ZetaChain will not be awarding prize funds for analysis or gas-specific submissions.

## Automated Findings / Publicly Known Issues

The 4naly3er report can be found [here](https://github.com/code-423n4/2023-10-zetachain/blob/main/4naly3er-report.md).

_Note for C4 wardens: Anything included in the 4naly3er **or** in the [previous audits](https://drive.google.com/drive/folders/10PFcoASYKhllalv5n1AW4mYD12urPgWJ) is considered a publicly known issue and is ineligible for awards._

# Overview

ZetaChain protocol is composed of two repositories:

- [`node`](repos/node/): ZetaChain source code based on Cosmos-SDK
- [`protocol-contracts`](repos/protocol-contracts/): Smart contracts deployed on ZetaChain or external chains to support interoperability

## Warden Resources for ZetaChain and Cosmos SDK: onboarding

Prior to this competitive audit, 3 teams of Code4rena wardens competed to produce a set of resources to help accelerate wardens’ ability to compete. Wardens unfamiliar with ZetaChain and/or Cosmos SDK are recommended to review the materials created by each team:

- [ZetaLotus](https://github.com/code-423n4/alpha-zetachain/tree/main/BlockChomper-0xladboy-reentrant)
- [032](https://github.com/code-423n4/alpha-zetachain/tree/main/032)
- [Gummy Bear Guardians](https://github.com/code-423n4/alpha-zetachain/tree/main/KrisApostolov-juancito)

In each team's workspace, you'll find: 

- Project 101/FAQ (“5 minutes to coding”)
- Testing 101/FAQ (“5 minutes to testing” plus best resources)
- A collection/documentation of classes of vulnerabilities

Teams ZetaLotus and 032 also produced: 

- Threat models
- architecture diagrams
- functional call flows

[The consolidated threat models, and links to additional resoures, can be viewed here.](https://github.com/code-423n4/2023-11-zetachain/blob/main/threat_models/README.md)

## Node

ZetaChain is based on Cosmos-SDK - [see here our usage of the framework](https://github.com/code-423n4/2023-11-zetachain/blob/main/docs/usage-cosmos-sdk.md)

Overview of the architecture of the node [can be found here](https://www.zetachain.com/docs/architecture/overview/)

The main sections of the source code are:
- [`x (modules)`](repos/node/x/): contains the source code of the Cosmos-SDK modules of the blockchain
  - [More info about the modules](docs/modules.md)
- [`zetaclient`](repos/node/zetaclient/): contains the code for the observer client validating cross-chain transactions on ZetaChain
- [`smoketests`](repos/node/contrib/): contains utilities to run smoke tests of the protocol, and local experimentation

## Protocol Contracts

The protocol contracts are separated into two sections:

- [`zevm`](repos/protocol-contracts/contracts/zevm/): contains contracts deployed on ZetaChain
- [`evm`](repos/protocol-contracts/contracts/evm/): contains contracts deployed on external EVM chains to be supported by ZetaChain

## Commands

### Requirements

- Go 1.20 (ZetaChain can only be built with this version of Go specifically)
- Docker (for smoke tests)
- Yarn

### `node`

Build the `zetacored` (blockchain node binary), and `zetacliend` (ZetaClient binary)
```
make install
```
Run the unit tests
```
make test
```
Run a standalone local blockchain node
```
make init
make run
```
Run the smoke tests
```
make start-smoketest
make stop-smoketest
```

### `protocol-contracts`

Compile the smart contracts
```
yarn
yarn compile
```

Run the unit tests
```
yarn test
```

## Experimentation with a local network

### Containers

The smoke tests under `contrib` allow testing of the different workflow of cross-chain functionalities on an E2E basis.

It also allows to experimentation of the protocol in a local environment. Running the smoke tests create several containers including:

- `zetacore0`: a ZetaChain node
- `zetaclient0`: a observer ZetaClient
- `eth`: a local Ethereum network connected to ZetaChain
- `bitcoin`: a local Bitcoin network connected to ZetaChain
- `orchestrator`: smoke tests runnner

After starting the networks with:

```
make start-smoketest
```

### ZetaChain

The user can connect to the `zetacore0` and directly use the node CLI with the `zetacored` binary with a funded account:
```
docker exec -it zetacore0 sh

/usr/local/bin # zetacored q bank balances zeta172uf5cwptuhllf6n4qsncd9v6xh59waxnu83kq
balances:
- amount: "4199000000000000000000000"
  denom: azeta
```

### Ethereum

The user can interact with the local Ethereum node with the exposed RPC on `http://0.0.0.0:8545`.
The following testing account is funded:
```
Address: 0xE5C5367B8224807Ac2207d350E60e1b6F27a7ecC
Private key: d87baf7bf6dc560a252596678c12e41f7d1682837f05b29d411bc3f78ae2c263
```

Examples with the `cast` CLI:
```
cast balance 0xE5C5367B8224807Ac2207d350E60e1b6F27a7ecC --rpc-url http://0.0.0.0:8545
98897999997945970464

cast send 0x9fd96203f7b22bCF72d9DCb40ff98302376cE09c --value 42 --rpc-url http://0.0.0.0:8545 --private-key "d87baf7bf6dc560a252596678c12e41f7d1682837f05b29d411bc3f78ae2c263"
```

### Custom programmatic interactions

The `smoketest` package contains an API to interact programmatically with the different network:
```go
type SmokeTest struct {
	zetaTxServer    ZetaTxServer
	cctxClient      crosschaintypes.QueryClient
	fungibleClient  fungibletypes.QueryClient
	authClient      authtypes.QueryClient
	bankClient      banktypes.QueryClient
	observerClient  observertypes.QueryClient
	goerliAuth      *bind.TransactOpts
	zevmAuth        *bind.TransactOpts
}
```
The user can use this API for custom testing on the networks and insert custom tests in `smoketest/main.go` (and commenting out unnecessary tests), the tests will automatically run upon starting the smoke tests. Current existing smoke tests are a good source to learn how to implement custom tests.

## Areas of Concerns

- Some of the areas of concern for the protocol security [can be found here](https://github.com/code-423n4/2023-11-zetachain/blob/main/docs/important-areas.md).
- Threat models from the ZetaChain Alpha competition [can be viewed here](https://github.com/code-423n4/2023-11-zetachain/blob/main/threat_models/README.md).

## Links

- **Previous audits:** https://drive.google.com/drive/folders/10PFcoASYKhllalv5n1AW4mYD12urPgWJ
- **Documentation:** https://www.zetachain.com/docs/
- **Website:** https://www.zetachain.com/
- **Twitter:** https://twitter.com/zetablockchain
- **Discord:** https://discord.gg/zetachain

# Scope

## Contracts

|File                                                               |SLOC|
|--------------------------------------------------------------------|--------|
|[contracts/zevm/ZRC20.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/ZRC20.sol)                                            |159     |
|[contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol)         |143     |
|[contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol)           |132     |
|[contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol)             |131     |
|[contracts/evm/ERC20Custody.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/ERC20Custody.sol)                                      |120     |
|[contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol)             |120     |
|[contracts/evm/tools/ImmutableCreate2Factory.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ImmutableCreate2Factory.sol)                     |93      |
|[contracts/evm/ZetaConnector.base.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/ZetaConnector.base.sol)                                |92      |
|[contracts/zevm/SystemContract.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/SystemContract.sol)                                   |89      |
|[contracts/zevm/ZetaConnectorZEVM.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/ZetaConnectorZEVM.sol)                                |73      |
|[contracts/evm/ZetaConnector.non-eth.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/ZetaConnector.non-eth.sol)                             |71      |
|[contracts/evm/ZetaConnector.eth.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/ZetaConnector.eth.sol)                                 |65      |
|[contracts/zevm/WZETA.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/WZETA.sol)                                            |47      |
|[contracts/evm/Zeta.non-eth.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/Zeta.non-eth.sol)                                      |43      |
|[contracts/evm/tools/ZetaInteractor.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/ZetaInteractor.sol)                              |36      |
|[contracts/evm/interfaces/ZetaInterfaces.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol)                         |36      |
|[contracts/evm/tools/interfaces/TridentIPoolRouter.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/interfaces/TridentIPoolRouter.sol)               |35      |
|[contracts/zevm/Interfaces.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/Interfaces.sol)                                       |9       |
|[contracts/evm/interfaces/ZetaErrors.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/interfaces/ZetaErrors.sol)                             |9       |
|[contracts/evm/interfaces/ConnectorErrors.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/interfaces/ConnectorErrors.sol)                        |9       |
|[contracts/zevm/interfaces/zContract.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/interfaces/zContract.sol)                             |8       |
|[contracts/zevm/interfaces/IWZETA.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/interfaces/IWZETA.sol)                                |7       |
|[contracts/evm/Zeta.eth.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/Zeta.eth.sol)                                          |7       |
|[contracts/evm/interfaces/ZetaInteractorErrors.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/interfaces/ZetaInteractorErrors.sol)                   |7       |
|[contracts/zevm/interfaces/IUniswapV2Router02.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/interfaces/IUniswapV2Router02.sol)                    |4       |
|[contracts/zevm/Uniswap.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/Uniswap.sol)                                          |4       |
|[contracts/evm/interfaces/ZetaNonEthInterface.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/interfaces/ZetaNonEthInterface.sol)                    |4       |
|[contracts/zevm/interfaces/IZRC20.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol)                                |3       |
|[contracts/evm/tools/interfaces/TridentConcentratedLiquidityPoolFactory.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/evm/tools/interfaces/TridentConcentratedLiquidityPoolFactory.sol)|3       |
|[contracts/zevm/UniswapPeriphery.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/UniswapPeriphery.sol)                                 |3       |
|[contracts/zevm/interfaces/IUniswapV2Router01.sol](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/protocol-contracts/contracts/zevm/interfaces/IUniswapV2Router01.sol)                    |3       |


## Node

|File                                                               |SLOC|
|--------------------------------------------------------------------|--------|
|[zetaclient/bitcoin_client.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/bitcoin_client.go)                                      |1082    |
|[zetaclient/evm_client.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/evm_client.go)                                          |1070    |
|[app/app.go](app/app.go)                                                        |612     |
|[x/fungible/keeper/system_contract.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/system_contract.go)                              |611     |
|[x/fungible/keeper/evm.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/evm.go)                                          |609     |
|[rpc/websockets.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/websockets.go)                                                 |592     |
|[zetaclient/evm_signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/evm_signer.go)                                          |555     |
|[server/start.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/start.go)                                                   |533     |
|[zetaclient/tss_signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/tss_signer.go)                                          |522     |
|[rpc/namespaces/ethereum/eth/filters/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/filters/api.go)                        |482     |
|[cmd/zetacored/observer_accounts.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/observer_accounts.go)                                |413     |
|[rpc/backend/blocks.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/blocks.go)                                             |391     |
|[zetaclient/query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/query.go)                                               |368     |
|[rpc/backend/tx_info.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/tx_info.go)                                            |362     |
|[zetaclient/btc_signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/btc_signer.go)                                         |312     |
|[rpc/backend/call_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/call_tx.go)                                            |311     |
|[rpc/namespaces/ethereum/eth/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/api.go)                                |287     |
|[x/crosschain/keeper/grpc_zevm.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/grpc_zevm.go)                                  |287     |
|[zetaclient/zetacore_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/zetacore_observer.go)                                   |284     |
|[x/crosschain/keeper/gas_payment.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/gas_payment.go)                                |268     |
|[rpc/types/events.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/events.go)                                               |267     |
|[x/crosschain/keeper/evm_hooks.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/evm_hooks.go)                                  |264     |
|[cmd/zetacored/root.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/root.go)                                             |261     |
|[zetaclient/config/types.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/config/types.go)                                        |255     |
|[zetaclient/tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/tx.go)                                                  |254     |
|[rpc/backend/node_info.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/node_info.go)                                          |244     |
|[rpc/types/utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/utils.go)                                                |243     |
|[cmd/zetaclientd/start.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/start.go)                                          |242     |
|[zetaclient/inbound_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/inbound_tracker.go)                                     |239     |
|[rpc/namespaces/ethereum/eth/filters/filter_system.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/filters/filter_system.go)              |238     |
|[server/config/config.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/config/config.go)                                           |236     |
|[x/crosschain/keeper/keeper_cross_chain_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_cross_chain_tx.go)                      |236     |
|[rpc/backend/utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/utils.go)                                              |228     |
|[x/crosschain/client/cli/cli_cctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_cctx.go)                               |222     |
|[zetaclient/zeta_supply_checker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/zeta_supply_checker.go)                                 |217     |
|[rpc/backend/chain_info.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/chain_info.go)                                         |204     |
|[scripts/gen-spec.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/scripts/gen-spec.go)                                               |203     |
|[zetaclient/zetacore_bridge.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/zetacore_bridge.go)                                     |197     |
|[rpc/namespaces/ethereum/eth/filters/filters.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/filters/filters.go)                    |196     |
|[cmd/zetaclientd/p2p_diagnostics.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/p2p_diagnostics.go)                                |193     |
|[zetaclient/telemetry.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/telemetry.go)                                           |191     |
|[common/pubkey.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/pubkey.go)                                                  |177     |
|[zetaclient/utils.go](zetaclient/utils.go)                                               |176     |
|[x/crosschain/keeper/keeper_cross_chain_tx_vote_outbound_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_cross_chain_tx_vote_outbound_tx.go)     |172     |
|[x/crosschain/keeper/msg_server_add_to_outtx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_server_add_to_outtx_tracker.go)            |171     |
|[common/chain.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/chain.go)                                                   |170     |
|[zetaclient/keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/keys.go)                                                |165     |
|[x/crosschain/keeper/keeper_gas_price.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_gas_price.go)                           |165     |
|[rpc/apis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/apis.go)                                                       |165     |
|[cmd/zetaclientd/keygen_tss.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/keygen_tss.go)                                     |164     |
|[rpc/backend/account_info.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/account_info.go)                                       |164     |
|[x/crosschain/keeper/keeper_cross_chain_tx_vote_inbound_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_cross_chain_tx_vote_inbound_tx.go)      |162     |
|[x/emissions/types/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/params.go)                                       |150     |
|[x/observer/types/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/params.go)                                        |149     |
|[rpc/types/block.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/block.go)                                                |148     |
|[cmd/zetacored/genaccounts.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/genaccounts.go)                                      |147     |
|[rpc/backend/tracing.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/tracing.go)                                            |147     |
|[cmd/zetacore_utils/main.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacore_utils/main.go)                                        |146     |
|[x/crosschain/keeper/msg_server_whitelist_erc20.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_server_whitelist_erc20.go)                 |143     |
|[x/observer/keeper/hooks.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/hooks.go)                                        |141     |
|[x/fungible/module.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/module.go)                                              |139     |
|[zetaclient/broadcast.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/broadcast.go)                                           |139     |
|[x/crosschain/client/cli/cli_tss.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_tss.go)                                |137     |
|[x/crosschain/module.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/module.go)                                            |136     |
|[app/export.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/export.go)                                                     |135     |
|[x/crosschain/client/cli/cli_out_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_out_tx_tracker.go)                     |134     |
|[x/observer/keeper/observer_mapper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/observer_mapper.go)                              |132     |
|[rpc/backend/backend.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/backend.go)                                            |132     |
|[common/headers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/headers.go)                                                 |130     |
|[x/crosschain/keeper/msg_migrate_tss_funds.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_migrate_tss_funds.go)                      |130     |
|[x/fungible/keeper/begin_blocker_deploy_system_contracts_privnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/begin_blocker_deploy_system_contracts_privnet.go)|129     |
|[x/observer/keeper/keeper_utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper_utils.go)                                 |129     |
|[x/crosschain/types/expected_keepers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/expected_keepers.go)                            |128     |
|[zetaclient/types/sql_evm.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/sql_evm.go)                                       |128     |
|[zetaclient/hsm/hsm_signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/hsm/hsm_signer.go)                                      |127     |
|[x/crosschain/keeper/keeper_out_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_out_tx_tracker.go)                      |126     |
|[x/observer/module.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/module.go)                                              |125     |
|[common/ethereum/proof.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/ethereum/proof.go)                                          |124     |
|[app/ante/handler_options.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/ante/handler_options.go)                                       |123     |
|[rpc/namespaces/ethereum/personal/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/personal/api.go)                           |118     |
|[x/emissions/module.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/module.go)                                             |115     |
|[rpc/ethereum/pubsub/pubsub.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/ethereum/pubsub/pubsub.go)                                     |114     |
|[zetaclient/interfaces.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/interfaces.go)                                          |110     |
|[rpc/backend/sign_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/sign_tx.go)                                            |110     |
|[x/emissions/abci.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/abci.go)                                               |110     |
|[x/crosschain/keeper/keeper_chain_nonces.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_chain_nonces.go)                        |109     |
|[x/fungible/keeper/gas_coin_and_pool.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/gas_coin_and_pool.go)                            |108     |
|[x/crosschain/keeper/keeper_tss.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_tss.go)                                 |107     |
|[x/observer/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/genesis.go)                                             |106     |
|[server/util.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/util.go)                                                    |105     |
|[x/crosschain/keeper/msg_tss_voter.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_tss_voter.go)                              |104     |
|[x/crosschain/client/cli/cli_chain_nonce.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_chain_nonce.go)                        |103     |
|[server/indexer_cmd.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/indexer_cmd.go)                                             |103     |
|[x/crosschain/keeper/evm_deposit.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/evm_deposit.go)                                |103     |
|[x/crosschain/keeper/cctx_utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/cctx_utils.go)                                 |100     |
|[app/ante/ante.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/ante/ante.go)                                                  |99      |
|[cmd/zetaclientd/aux.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/aux.go)                                            |98      |
|[x/crosschain/client/cli/cli_in_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_in_tx_tracker.go)                      |96      |
|[zetaclient/types/sql_btc.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/sql_btc.go)                                       |95      |
|[x/observer/types/message_update_core_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_update_core_params.go)                    |95      |
|[x/observer/keeper/msg_server_add_block_header.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_add_block_header.go)                  |94      |
|[x/crosschain/keeper/verify_proof.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/verify_proof.go)                               |93      |
|[server/json_rpc.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/json_rpc.go)                                                |92      |
|[x/crosschain/keeper/keeper_pending_nonces.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_pending_nonces.go)                      |91      |
|[x/observer/client/cli/query_blame.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_blame.go)                              |90      |
|[x/crosschain/client/cli/cli_gas_price.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_gas_price.go)                          |89      |
|[x/observer/keeper/blame.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/blame.go)                                        |89      |
|[x/crosschain/keeper/abci.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/abci.go)                                       |89      |
|[x/observer/keeper/ballot.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/ballot.go)                                      |89      |
|[server/indexer_service.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/indexer_service.go)                                         |88      |
|[x/observer/types/ballot.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/ballot.go)                                        |88      |
|[x/fungible/keeper/foreign_coins.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/foreign_coins.go)                                |87      |
|[x/crosschain/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/genesis.go)                                           |86      |
|[rpc/namespaces/ethereum/eth/filters/utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/filters/utils.go)                      |85      |
|[cmd/zetaclientd/init.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/init.go)                                           |85      |
|[x/observer/keeper/msg_server_update_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_update_observer.go)                   |83      |
|[x/crosschain/types/message_vote_on_observed_inbound_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_vote_on_observed_inbound_tx.go)         |83      |
|[common/proof.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/proof.go)                                                   |82      |
|[x/observer/client/cli/tx_add_blame_vote.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_add_blame_vote.go)                        |82      |
|[zetaclient/signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/signer.go)                                              |82      |
|[x/crosschain/keeper/grpc_query_get_tss_address.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/grpc_query_get_tss_address.go)                 |79      |
|[cmd/zetaclientd/hsm.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/hsm.go)                                            |79      |
|[x/crosschain/keeper/keeper_last_block_height.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_last_block_height.go)                   |78      |
|[zetaclient/metrics/metrics.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/metrics/metrics.go)                                     |78      |
|[x/crosschain/keeper/events.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/events.go)                                     |77      |
|[x/crosschain/client/cli/query_in_tx_hash_to_cctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/query_in_tx_hash_to_cctx.go)               |77      |
|[x/fungible/keeper/msg_server_update_system_contract.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_update_system_contract.go)            |75      |
|[x/fungible/keeper/gas_price.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/gas_price.go)                                    |75      |
|[app/ante/authz.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/ante/authz.go)                                                 |74      |
|[rpc/types/types.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/types.go)                                                |74      |
|[x/fungible/client/cli/query_gas_stability_pool.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/query_gas_stability_pool.go)                 |73      |
|[contrib/localnet/README.md](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/contrib/localnet/README.md)                                        |72      |
|[x/crosschain/types/message_vote_on_observed_outbound_tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_vote_on_observed_outbound_tx.go)        |72      |
|[x/observer/keeper/block_header.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/block_header.go)                                 |71      |
|[x/observer/keeper/keeper_node_account.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper_node_account.go)                          |71      |
|[server/flags/flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/flags/flags.go)                                             |70      |
|[zetaclient/config/config.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/config/config.go)                                       |69      |
|[x/crosschain/keeper/in_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/in_tx_tracker.go)                              |69      |
|[x/fungible/keeper/grpc_gas_stability_pool.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/grpc_gas_stability_pool.go)                      |68      |
|[x/fungible/keeper/deposits.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/deposits.go)                                     |67      |
|[common/cosmos/cosmos.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/cosmos/cosmos.go)                                           |67      |
|[x/observer/types/core_params_mainnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/core_params_mainnet.go)                           |66      |
|[standalone-network/evm-init](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/standalone-network/evm-init)                                       |65      |
|[server/config/toml.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/server/config/toml.go)                                             |65      |
|[x/observer/types/messages_add_block_header.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/messages_add_block_header.go)                     |64      |
|[x/fungible/keeper/msg_server_update_zrc20_withdraw_fee.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_update_zrc20_withdraw_fee.go)         |64      |
|[x/crosschain/keeper/keeper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper.go)                                     |64      |
|[x/crosschain/keeper/grpc_query_in_tx_hash_to_cctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/grpc_query_in_tx_hash_to_cctx.go)              |64      |
|[app/config.go](app/config.go)                                                     |63      |
|[x/fungible/keeper/msg_server_update_contract_bytecode.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_update_contract_bytecode.go)          |62      |
|[x/crosschain/types/keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/keys.go)                                        |62      |
|[x/crosschain/migrations/v2/migrate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/migrations/v2/migrate.go)                             |62      |
|[common/default_chains_mainnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/default_chains_mainnet.go)                                  |62      |
|[x/fungible/types/expected_keepers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/expected_keepers.go)                              |62      |
|[x/emissions/keeper/keeper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/keeper.go)                                      |61      |
|[x/fungible/keeper/keeper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/keeper.go)                                       |61      |
|[x/fungible/keeper/gas_stability_pool.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/gas_stability_pool.go)                           |60      |
|[x/fungible/client/cli/tx_deploy_fungible_coin_zrc_4.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/tx_deploy_fungible_coin_zrc_4.go)            |60      |
|[zetaclient/klaytn_client.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/klaytn_client.go)                                       |59      |
|[x/observer/client/cli/query_core_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_core_params.go)                        |58      |
|[x/crosschain/types/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/genesis.go)                                     |58      |
|[common/bitcoin/bitcoin_spv.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/bitcoin/bitcoin_spv.go)                                     |58      |
|[cmd/zetacored/collect_observer_info.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/collect_observer_info.go)                            |58      |
|[x/observer/keeper/grpc_query_prove.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/grpc_query_prove.go)                             |58      |
|[cmd/zetacored/get_pubkey.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/get_pubkey.go)                                       |55      |
|[x/observer/types/messages_crosschain_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/messages_crosschain_flags.go)                     |55      |
|[x/observer/keeper/keeper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper.go)                                       |55      |
|[x/crosschain/types/message_whitelist_erc20.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_whitelist_erc20.go)                     |55      |
|[x/fungible/client/cli/query_foreign_coins.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/query_foreign_coins.go)                      |55      |
|[x/observer/client/cli/tx_update_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_update_observer.go)                       |55      |
|[zetaclient/out_tx_processor_manager.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/out_tx_processor_manager.go)                            |54      |
|[x/observer/client/cli/query_observers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_observers.go)                          |54      |
|[x/crosschain/client/cli/cli_last_block_height.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_last_block_height.go)                  |54      |
|[x/observer/types/message_update_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_update_observer.go)                       |54      |
|[x/crosschain/types/message_add_to_out_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_add_to_out_tx_tracker.go)               |54      |
|[x/observer/client/cli/cli_node_account.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/cli_node_account.go)                         |54      |
|[cmd/zetaclientd/main.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/main.go)                                           |54      |
|[x/fungible/keeper/msg_server_update_zrc20_paused_status.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_update_zrc20_paused_status.go)        |53      |
|[x/observer/types/message_add_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_add_observer.go)                          |53      |
|[x/crosschain/types/message_add_to_in_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_add_to_in_tx_tracker.go)                |52      |
|[x/emissions/keeper/withdrawable_emissions.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/withdrawable_emissions.go)                      |52      |
|[x/fungible/keeper/msg_server_deploy_fungible_coin_zrc20.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_deploy_fungible_coin_zrc20.go)        |52      |
|[common/bitcoin/proof.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/bitcoin/proof.go)                                           |52      |
|[x/crosschain/types/message_set_node_keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_set_node_keys.go)                       |52      |
|[x/observer/types/core_params_privnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/core_params_privnet.go)                           |52      |
|[x/emissions/keeper/block_rewards_components.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/block_rewards_components.go)                    |51      |
|[x/fungible/types/message_update_zrc20_paused_status.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_update_zrc20_paused_status.go)            |51      |
|[x/fungible/types/message_deploy_fungible_coin_zrc20.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_deploy_fungible_coin_zrc20.go)            |50      |
|[x/crosschain/types/status.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/status.go)                                      |50      |
|[x/observer/client/cli/tx_update_core_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_update_core_params.go)                    |49      |
|[x/observer/types/message_add_blame_vote.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_add_blame_vote.go)                        |49      |
|[cmd/zetacored/addr_converter.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/addr_converter.go)                                   |48      |
|[x/fungible/types/message_update_zrc20_withdraw_fee.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_update_zrc20_withdraw_fee.go)             |48      |
|[common/default_chains_privnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/default_chains_privnet.go)                                  |48      |
|[x/fungible/types/message_update_contract_bytecode.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_update_contract_bytecode.go)              |47      |
|[zetaclient/config/config_mainnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/config/config_mainnet.go)                               |47      |
|[x/crosschain/keeper/in_tx_hash_to_cctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/in_tx_hash_to_cctx.go)                         |47      |
|[zetaclient/config/config_privnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/config/config_privnet.go)                               |46      |
|[x/observer/keeper/crosschain_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/crosschain_flags.go)                             |46      |
|[x/crosschain/types/message_migrate_tss_funds.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_migrate_tss_funds.go)                   |46      |
|[x/observer/keeper/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/params.go)                                       |46      |
|[x/observer/keeper/msg_server_add_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_add_observer.go)                      |46      |
|[rpc/namespaces/ethereum/eth/filters/subscription.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/eth/filters/subscription.go)               |45      |
|[x/crosschain/keeper/msg_server_add_to_intx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_server_add_to_intx_tracker.go)             |45      |
|[x/fungible/keeper/grpc_query_foreign_coins.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/grpc_query_foreign_coins.go)                     |45      |
|[x/fungible/types/message_update_zrc20_liquidity_cap.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_update_zrc20_liquidity_cap.go)            |44      |
|[x/observer/types/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/genesis.go)                                       |44      |
|[cmd/zetacored/main.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/main.go)                                             |43      |
|[x/fungible/keeper/evm_hooks.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/evm_hooks.go)                                    |43      |
|[x/crosschain/types/message_gas_price_voter.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_gas_price_voter.go)                     |43      |
|[x/crosschain/types/cctx_utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/cctx_utils.go)                                  |43      |
|[x/observer/types/parsers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/parsers.go)                                       |43      |
|[x/crosschain/types/message_tss_voter.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_tss_voter.go)                           |43      |
|[x/crosschain/keeper/zeta_conversion_rate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/zeta_conversion_rate.go)                       |43      |
|[x/crosschain/types/message_update_tss_address.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_update_tss_address.go)                  |42      |
|[x/observer/keeper/grpc_query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/grpc_query_params.go)                            |42      |
|[cmd/zetacored/parsers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/parsers.go)                                          |42      |
|[x/crosschain/types/codec.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/codec.go)                                       |42      |
|[x/fungible/types/message_update_system_contract.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_update_system_contract.go)                |41      |
|[rpc/namespaces/ethereum/net/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/net/api.go)                                |41      |
|[x/crosschain/types/message_remove_from_out_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_remove_from_out_tx_tracker.go)          |41      |
|[x/crosschain/types/message_nonce_voter.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/message_nonce_voter.go)                         |41      |
|[x/observer/keeper/msg_server_update_crosschain_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_update_crosschain_flags.go)           |41      |
|[zetaclient/chainmetrics.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/chainmetrics.go)                                        |41      |
|[cmd/zetaclientd/start_utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/start_utils.go)                                    |41      |
|[x/observer/module_simulation.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/module_simulation.go)                                   |40      |
|[x/observer/client/cli/tx_add_observer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_add_observer.go)                          |40      |
|[app/ante/vesting.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/ante/vesting.go)                                               |40      |
|[x/observer/keeper/events.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/events.go)                                       |40      |
|[x/observer/types/observer_mapper.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/observer_mapper.go)                               |39      |
|[rpc/types/query_client.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/query_client.go)                                         |39      |
|[x/crosschain/client/cli/query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/query.go)                                  |39      |
|[x/observer/keeper/msg_server_add_blame_vote.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_add_blame_vote.go)                    |38      |
|[x/observer/types/keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/keys.go)                                          |38      |
|[rpc/namespaces/ethereum/txpool/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/txpool/api.go)                             |37      |
|[x/fungible/types/message_remove_foreign_coin.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/message_remove_foreign_coin.go)                   |37      |
|[x/observer/client/cli/tx_permission_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_permission_flags.go)                      |37      |
|[x/observer/types/message_update_keygen.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_update_keygen.go)                         |37      |
|[x/observer/client/cli/tx_update_keygen.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx_update_keygen.go)                         |36      |
|[x/crosschain/types/errors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/errors.go)                                      |36      |
|[x/fungible/client/cli/tx_update_zrc20_liquidity_cap.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/tx_update_zrc20_liquidity_cap.go)            |36      |
|[.github/actions/deploy-binaries/functions](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/actions/deploy-binaries/functions)                         |36      |
|[x/crosschain/client/cli/query_get_tss_address.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/query_get_tss_address.go)                  |35      |
|[x/observer/keeper/keeper_keygen.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper_keygen.go)                                |35      |
|[x/emissions/types/keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/keys.go)                                         |34      |
|[common/utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/utils.go)                                                   |34      |
|[x/observer/client/cli/query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query.go)                                    |34      |
|[common/address.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/address.go)                                                 |33      |
|[x/observer/abci.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/abci.go)                                                |33      |
|[app/ante/interfaces.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/ante/interfaces.go)                                            |33      |
|[x/fungible/types/codec.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/codec.go)                                         |32      |
|[x/emissions/types/expected_keepers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/expected_keepers.go)                             |32      |
|[x/observer/client/cli/query_ballot.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_ballot.go)                             |32      |
|[x/fungible/client/cli/tx_remove_foreign_coin.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/tx_remove_foreign_coin.go)                   |32      |
|[x/emissions/client/cli/query_show_available_emissions.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/query_show_available_emissions.go)          |32      |
|[x/fungible/types/errors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/errors.go)                                        |32      |
|[zetaclient/authz_signer.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/authz_signer.go)                                        |32      |
|[x/observer/types/codec.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/codec.go)                                         |32      |
|[x/observer/types/errors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/errors.go)                                        |31      |
|[rpc/namespaces/ethereum/miner/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/miner/api.go)                              |31      |
|[x/observer/keeper/msg_server_update_keygen.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_update_keygen.go)                     |30      |
|[x/crosschain/types/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/params.go)                                      |30      |
|[x/fungible/module_simulation.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/module_simulation.go)                                   |29      |
|[.github/actions/cosmovisor-upgrade/functions](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/actions/cosmovisor-upgrade/functions)                      |29      |
|[x/emissions/client/cli/query_get_emmisons_factors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/query_get_emmisons_factors.go)              |29      |
|[x/observer/types/expected_keepers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/expected_keepers.go)                              |29      |
|[x/crosschain/module_simulation.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/module_simulation.go)                                 |29      |
|[common/chain_network.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/chain_network.go)                                           |29      |
|[rpc/namespaces/ethereum/miner/unsupported.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/miner/unsupported.go)                      |29      |
|[x/fungible/types/zrc20.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/zrc20.go)                                         |29      |
|[x/emissions/client/cli/query_list_balances.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/query_list_balances.go)                     |29      |
|[x/crosschain/client/cli/cli_zeta_height.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_zeta_height.go)                        |29      |
|[x/observer/client/cli/query_show_observer_count.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_show_observer_count.go)                |29      |
|[x/observer/keeper/msg_server_update_core_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server_update_core_params.go)                |29      |
|[x/crosschain/client/cli/tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/tx.go)                                     |29      |
|[cmd/zetacored/flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/flags.go)                                            |28      |
|[app/setup_handlers.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/setup_handlers.go)                                             |28      |
|[x/observer/migrations/v4/migrate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/migrations/v4/migrate.go)                               |28      |
|[cmd/zetacored/config/config.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/config/config.go)                                    |28      |
|[x/fungible/types/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/params.go)                                        |28      |
|[x/observer/migrations/v3/migrate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/migrations/v3/migrate.go)                               |28      |
|[x/observer/client/cli/cli_supported_chains.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/cli_supported_chains.go)                     |27      |
|[x/emissions/types/events.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/events.go)                                       |27      |
|[x/observer/client/cli/cli_keygen.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/cli_keygen.go)                               |27      |
|[x/observer/client/cli/query_permission_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_permission_flags.go)                   |27      |
|[x/crosschain/migrations/v3/migrate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/migrations/v3/migrate.go)                             |27      |
|[x/fungible/client/cli/query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/query.go)                                    |26      |
|[rpc/backend/filters.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/backend/filters.go)                                            |26      |
|[x/observer/client/cli/tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/tx.go)                                       |26      |
|[x/emissions/client/cli/query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/query_params.go)                            |26      |
|[x/observer/client/cli/query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/client/cli/query_params.go)                             |26      |
|[x/crosschain/client/cli/cli_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/client/cli/cli_params.go)                             |26      |
|[rpc/types/addrlock.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/types/addrlock.go)                                             |26      |
|[x/fungible/client/cli/query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/query_params.go)                             |26      |
|[x/fungible/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/genesis.go)                                             |25      |
|[x/fungible/client/cli/query_system_contract.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/query_system_contract.go)                    |25      |
|[x/fungible/types/keys.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/keys.go)                                          |25      |
|[x/observer/keeper/migrator.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/migrator.go)                                     |24      |
|[x/crosschain/keeper/grpc_query_in_tx_tracker.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/grpc_query_in_tx_tracker.go)                   |23      |
|[common/authz_tx_types.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/authz_tx_types.go)                                          |23      |
|[x/fungible/keeper/msg_server_udpate_zrc20_liquidity_cap.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_udpate_zrc20_liquidity_cap.go)        |22      |
|[x/observer/keeper/keeper_block_header.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper_block_header.go)                          |22      |
|[x/fungible/client/cli/tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/client/cli/tx.go)                                       |22      |
|[x/fungible/types/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/genesis.go)                                      |22      |
|[x/crosschain/keeper/keeper_last_zeta_height.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_last_zeta_height.go)                    |22      |
|[x/observer/migrations/v2/migrate.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/migrations/v2/migrate.go)                               |21      |
|[.github/actions/cosmovisor-upgrade/readme.md](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/actions/cosmovisor-upgrade/readme.md)                      |21      |
|[x/emissions/client/cli/query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/query.go)                                   |21      |
|[x/fungible/keeper/msg_server_remove_foreign_coin.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server_remove_foreign_coin.go)               |21      |
|[x/observer/keeper/grpc_query_show_observer_count.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/grpc_query_show_observer_count.go)               |21      |
|[x/crosschain/keeper/keeper_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/keeper_params.go)                              |21      |
|[x/crosschain/keeper/migrator.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/migrator.go)                                   |20      |
|[x/observer/types/crosschain_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/crosschain_flags.go)                              |20      |
|[x/fungible/keeper/grpc_query_system_contract.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/grpc_query_system_contract.go)                   |19      |
|[x/observer/keeper/grpc_query_crosschain_flags.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/grpc_query_crosschain_flags.go)                  |19      |
|[x/emissions/keeper/grpc_query_show_available_emissions.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/grpc_query_show_available_emissions.go)         |19      |
|[x/emissions/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/genesis.go)                                            |18      |
|[cmd/zetaclientd/root.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/root.go)                                           |18      |
|[zetaclient/block_height.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/block_height.go)                                        |18      |
|[rpc/namespaces/ethereum/web3/api.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/rpc/namespaces/ethereum/web3/api.go)                               |18      |
|[app/params/proto.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/params/proto.go)                                               |18      |
|[.github/pull_request_template.md](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/pull_request_template.md)                                  |17      |
|[x/emissions/client/cli/tx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/client/cli/tx.go)                                      |17      |
|[cmd/config_privnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/config_privnet.go)                                             |16      |
|[x/emissions/keeper/grpc_query_list_balances.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/grpc_query_list_balances.go)                    |15      |
|[x/emissions/keeper/grpc_query_get_emmisons_factors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/grpc_query_get_emmisons_factors.go)             |15      |
|[x/emissions/types/codec.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/codec.go)                                        |15      |
|[common/coin.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/coin.go)                                                    |15      |
|[cmd/zetaclientd/version.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetaclientd/version.go)                                        |15      |
|[x/emissions/keeper/grpc_query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/grpc_query_params.go)                           |15      |
|[x/fungible/keeper/grpc_query_params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/grpc_query_params.go)                            |15      |
|[.github/actions/upload-to-s3/readme.md](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/actions/upload-to-s3/readme.md])                            |14      |
|[x/fungible/keeper/zeta.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/zeta.go)                                         |14      |
|[x/crosschain/keeper/foreign_coins.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/foreign_coins.go)                              |13      |
|[x/fungible/types/key_foreign_coins.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/key_foreign_coins.go)                             |13      |
|[x/emissions/keeper/msg_server.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/msg_server.go)                                  |13      |
|[x/crosschain/types/key_in_tx_hash_to_cctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/key_in_tx_hash_to_cctx.go)                      |13      |
|[x/fungible/types/gas_stablity_pool.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/gas_stablity_pool.go)                             |13      |
|[app/prefix.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/prefix.go)                                                     |13      |
|[x/fungible/keeper/begin_blocker_deploy_system_contracts.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/begin_blocker_deploy_system_contracts.go)        |13      |
|[cmd/config_mainnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/config_mainnet.go)                                             |12      |
|[x/observer/types/utils.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/utils.go)                                         |12      |
|[app/params/encoding.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/params/encoding.go)                                            |12      |
|[x/fungible/types/evm.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/types/evm.go)                                           |12      |
|[x/observer/simulation/simap.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/simulation/simap.go)                                    |12      |
|[x/observer/keeper/keeper_supported_chain.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/keeper_supported_chain.go)                       |11      |
|[x/emissions/keeper/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/params.go)                                      |11      |
|[x/observer/keeper/msg_server.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/msg_server.go)                                   |11      |
|[x/fungible/keeper/msg_server.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/msg_server.go)                                   |11      |
|[x/crosschain/keeper/msg_server.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/msg_server.go)                                 |11      |
|[x/crosschain/types/client_ctx.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/types/client_ctx.go)                                  |11      |
|[cmd/zetacored/config/prefix_mainnet.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/cmd/zetacored/config/prefix_mainnet.go)                            |11      |
|[x/fungible/keeper/params.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/params.go)                                       |11      |
|[zetaclient/types/account_resp.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/account_resp.go)                                  |10      |
|[.github/actions/build-binaries/readme.md](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/.github/actions/build-binaries/readme.md)                          |10      |
|[x/emissions/types/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/genesis.go)                                      |10      |
|[app/encoding.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/encoding.go)                                                   |9       |
|[x/emissions/types/errors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/errors.go)                                      |9       |
|[app/genesis.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/app/genesis.go)                                                    |9       |
|[zetaclient/errors.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/errors.go)                                              |8       |
|[common/version.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/version.go)                                                 |7       |
|[zetaclient/types/ethish.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/ethish.go)                                        |7       |
|[zetaclient/supported_chains.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/supported_chains.go)                                    |7       |
|[zetaclient/misc.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/misc.go)                                                |7       |
|[x/crosschain/migrations/v2/legacytypes.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/migrations/v2/legacytypes.go)                         |6       |
|[x/fungible/keeper/grpc_query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/fungible/keeper/grpc_query.go)                                   |5       |
|[x/observer/keeper/grpc_query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/keeper/grpc_query.go)                                   |5       |
|[common/commands.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/common/commands.go)                                                |5       |
|[x/emissions/keeper/grpc_query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/keeper/grpc_query.go)                                  |5       |
|[x/crosschain/keeper/grpc_query.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/crosschain/keeper/grpc_query.go)                                 |5       |
|[zetaclient/types/telemetry_types.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/telemetry_types.go)                               |4       |
|[zetaclient/types/defs.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/zetaclient/types/defs.go)                                          |4       |
|[x/emissions/types/types.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/emissions/types/types.go)                                        |1       |
|[x/observer/types/message_set_supported_chains.go](https://github.com/code-423n4/2023-11-zetachain/tree/main/repos/node/x/observer/types/message_set_supported_chains.go)                  |1       |

## Out of scope

Please note that **ALL** the files are in scope, but the files listed in the previous section are where you should actually looks for bugs.

## Scoping Details 

```
- If you have a public code repo, please share it here: https://github.com/zeta-chain/node and https://github.com/zeta-chain/protocol-contracts/tree/main/contracts  
- How many contracts are in scope?: 13   
- Total SLoC for these contracts?: 0  
- How many external imports are there?:  
- How many separate interfaces and struct definitions are there for the contracts within scope?:  
- Does most of your code generally use composition or inheritance?: Inheritance  
- How many external calls?: 0   
- What is the overall line coverage percentage provided by your tests?:
- Is this an upgrade of an existing system?: False
- Check all that apply (e.g. timelock, NFT, AMM, ERC20, rollups, etc.): ERC-20 Token, Multi-Chain
- Is there a need to understand a separate part of the codebase / get context in order to audit this part of the protocol?: Yes   
- Please describe required context: You need to understand how the protocol works to understand how the cross chain elements come into play.   
- Does it use an oracle?: No 
- Describe any novel or unique curve logic or mathematical models your code uses: Protocol using shared TSS key to manage assets using decentralized validators 
- Is this either a fork of or an alternate implementation of another project?: True, Some concepts borrowed from ThorChain and Evmos   
- Does it use a side-chain?: False
- Describe any specific areas you would like addressed: Try to break cross-chain elements, ZRC20, accounting between external chain state and Zetachain state
```
