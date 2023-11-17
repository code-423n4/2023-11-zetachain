# ZetaChain Protocol Contracts

This repository contains the smart contracts for ZetaChain. The smart contracts
are written in Solidity, and the repository includes scripts to compile the
contracts and generate Go bindings.

## Importing Protocol Contracts

As a dApp developer, you can install the protocol contracts package into your
project:

```
yarn add --dev @zetachain/protocol-contracts
```

Importing
[`ZetaInterfaces`](https://www.zetachain.com/docs/developers/cross-chain-messaging/connector/)
and `ZetaInteractor` for cross-chain messaging:

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaInteractor.sol";
```

Importing [ZRC20](https://www.zetachain.com/docs/developers/concepts/zrc-20/)
and the [system
contract](https://www.zetachain.com/docs/developers/concepts/system-contract/)
for omni-chain smart contracts:

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
```

## Prerequisites

Before you can contribute to this project, you must have the following installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Getting Started

To get started with this project, you should first clone the repository:

```
git clone https://github.com/zeta-chain/protocol-contracts
```

Once you have cloned the repository, you can navigate to the project directory
and run the following command to install the project dependencies:

```
yarn
```

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts
to the `artifacts` directory.

## Generating Go Bindings

To generate Go bindings for the Solidity contracts, run the following command:

```
yarn generate
```

This will use `abigen` to generate Go bindings for the contracts and output the
resulting Go files to the `pkg` directory.

## Contributing

If you would like to contribute to this project, please fork the repository and
submit a pull request. All contributions are welcome!
