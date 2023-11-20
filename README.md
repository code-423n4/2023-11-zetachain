# ZetaChain audit details
- Total Prize Pool: $150,000 USDC 
  - HM awards: $121,400 USDC 
  - QA awards: $3,100 USDC 
  - Judge awards: $15,000 USDC
  - Lookout awards: $10,000 USDC 
  - Scout awards: $500 USDC 
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2023-11-zetachain/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts November 20, 2023 20:00 UTC
- Ends December 18, 2023 20:00 UTC 

## Automated Findings / Publicly Known Issues

The 4naly3er report can be found [here](https://github.com/code-423n4/2023-10-zetachain/blob/main/4naly3er-report.md).

_Note for C4 wardens: Anything included in the 4naly3er **or** the automated findings output is considered a publicly known issue and is ineligible for awards._

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

ZetaChain is based on Cosmos-SDK - [see here our usage of the framework](docs/usage-cosmos-sdk.md)

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

- Some of the areas of concern for the protocol security [can be found here](https://github.com/code-423n4/2023-11-zetachain/docs/important-areas.md). 
- Threat models from the ZetaChain Alpha competition [can be viewed here](https://github.com/code-423n4/2023-11-zetachain/blob/main/threat_models/README.md).

## Links

- **Previous audits:** https://drive.google.com/drive/folders/10PFcoASYKhllalv5n1AW4mYD12urPgWJ
- **Documentation:** https://www.zetachain.com/docs/
- **Website:** https://www.zetachain.com/
- **Twitter:** https://twitter.com/zetablockchain
- **Discord:** https://discord.gg/zetachain

# Scope

[ ⭐️ SPONSORS: add scoping and technical details here ]

- [ ] In the table format shown below, provide the name of each contract and:
  - [ ] source lines of code (excluding blank lines and comments) in each *For line of code counts, we recommend running prettier with a 100-character line length, and using [cloc](https://github.com/AlDanial/cloc).* 
  - [ ] external contracts called in each
  - [ ] libraries used in each

*List all files in scope in the table below (along with hyperlinks) -- and feel free to add notes here to emphasize areas of focus.*

| Contract | SLOC | Purpose | Libraries used |  
| ----------- | ----------- | ----------- | ----------- |
| [contracts/folder/sample.sol](https://github.com/code-423n4/repo-name/blob/contracts/folder/sample.sol) | 123 | This contract does XYZ | [`@openzeppelin/*`](https://openzeppelin.com/contracts/) |

## Out of scope

*List any files/contracts that are out of scope for this audit.*

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
- Check all that apply (e.g. timelock, NFT, AMM, ERC20, rollups, etc.): ERC-29 Token, Multi-Chain
- Is there a need to understand a separate part of the codebase / get context in order to audit this part of the protocol?: Yes   
- Please describe required context: You need to understand how the protocol works to understand how the cross chain elements come into play.   
- Does it use an oracle?: No 
- Describe any novel or unique curve logic or mathematical models your code uses: Protocol using shared TSS key to manage assets using decentralized validators 
- Is this either a fork of or an alternate implementation of another project?: True, Some concepts borrowed from ThorChain and Evmos   
- Does it use a side-chain?: False
- Describe any specific areas you would like addressed: Try to break cross-chain elements, ZRC20, accounting between external chain state and Zetachain state
```
