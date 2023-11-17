# Usage of Cosmos SDK

ZetaChain is a EVM compatible blockchain based on Cosmos SDK. The blockchain imports the default Cosmos SDK module for DPoS chains: 

- `auth`
- `bank`
- `staking`
- `slashing`
- `evidence`
- `gov`
- `distribution`
- `upgrade`
- `feemarket`
- `authz`
- `group`
- `crisis`
- `params`

`mint` and IBC modules are not imported.

The blockchain also imports the `evm` module from Ethermint to be EVM compatible.

The blockchain is composed of 4 custom modules: `crosschain`, `fungible`, `observer` and `emissions` . Each play a specific role for cross-chain transactions.

The ZetaChain EVM is called zEVM.

The `bank` module is used to manage the supply of ZETA tokens. Every other assets in the protocol like foreign coins are represented through smart contracts in the zEVM.