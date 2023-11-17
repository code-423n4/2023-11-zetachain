# Important Areas

## 1 - CCTX lifecycle

CCTXs are what define cross-chain interactions. The logics triggered during a CCTX lifecycle is one of the most important part of the protocol. The transported assets in CCTXs must be correctly handled when minted/burned on Zeta, used to pay for outbound tx gas, etc…

### Inbound Transaction Observation

If the destination chain is ZetaChain and the CCTX contains a message, ZRC20 tokens are deposited and a contract on ZetaChain is called. Destination contract address and the arguments for the contract call are contained within the message.

If the destination chain is not ZetaChain, the status of a transaction is changed to "pending outbound" and the CCTX to be processed as an outbound transaction.

More information on inbound tx observation can be found at [ZetaDocs (zetachain.com)](https://www.zetachain.com/docs/architecture/modules/crosschain/messages/#msgvoteonobservedinboundtx)

### Outbound Transaction Observation

**Pending Outbound**

Observers watch ZetaChain for pending outbound transactions. To process a pending outbound transactions observers enter into a TSS keysign ceremony to sign the transaction, and then broadcast the signed transaction to the connected blockchains.

****Observed Outbound****

Observers watch connected blockchains for the broadcasted outbound transactions. Once a transaction is "confirmed" (or "mined") on a connected blockchains, observers vote on ZetaChain by sending a `VoteOnObservedOutboundTx` message.

After the vote passes the threshold, the voting is finalized and a transaction's status is changed to final.

More information about outbound tx observation can be found at [ZetaDocs (zetachain.com)](https://www.zetachain.com/docs/architecture/modules/crosschain/messages/#msgvoteonobservedoutboundtx)

### Gas Payment on ZetaChain for outbound tx

Gas for outbound txs are paid on ZetaChain using the amount of token in the cross-chain tx (cctx)

The gas payment functionality has been extended to support 3 types of cctx:

- Gas cctx: the cctx contains the gas token of the chain and therefore it is directly subtracted from the amount for gas payment
- Zeta cctx: zeta tokens are subtracted from the amount are swapped into the zrc20 gas tokens for the gas payment
- ERC20 cctx: erc20 are subtracted from the amount and swapped erc20→zeta→gas for the gas payment

→ The swaps use Uniswap pools, native to ZetaChain, and expects the liquidity to be maintained

### Gas Price Increase for outbound tx

A  mechanism allows increasing the gas price used for an outbound tx on an external chain if it is not being mined

- When outbound txs are mined, part of the remaining fees (surplus gasUsed/gasLimit) are sent to a “gas stability pool”
- Every n blocks, every pending outbound txs are iterated, if pending for too long, the gas price for the outbound tx is increase and, funds from the gas stability pool is used to pay for the increase
- Gas is not increased if not enough funds in the gas stability pool

## 2 - Permissionless Tx Validation Model

A new architecture has been developed in order to validate txs from external chains permissionlessly instead of requiring observers to vote on each of them.

- Observer votes on block headers of the external chains instead of the txs. The block headers are stored on-chain on ZetaChain
- An inbound/outbound tx can be sent by anyone with a proof. The proof is checked against the header of the chain to validate it

→ The system is currently only used to validate txs to be added in the tracker through `MsgAddToInTxTracker` and `MsgAddToOutTxTracker` and not to validate txs to be observed as part of the ZetaChain cross-chain tx workflow

→ The system is currently implemented for Ethereum/BSC and Bitcoin

→ In the future, the validation of inbound txs and outbound txs will be refactored to use this system instead of observers. The observers will only vote for the block headers of external chains. Then anyone can validate a inbound or outbound tx for a cross-chain tx.

## 3 - Permissions and Admin Groups

- Certain action in the blockchain can only be executed by a specific address (multi-sig). There are two admin groups:
    - Group1
        - Emergency actions: halt a workflow of the blockchain in case of emergency
        - Example: freeze a ZRC20 if a vulnerabilty is found for this token
        - Rapidity over security: represented by an address that would require a single signature from the multisig
        - Should not concern action that can be financially exploitable
    - Group2
        - Operational: updating parameters influencing logic of the blockchain
        - Example: updating the system contract
        - Security over rapidity: requires several signature

**Permissions Table**

| Message | Group | Description | Module |
| --- | --- | --- | --- |
| MsgUpdateZRC20PausedStatus - Pause | 1 | Pause interaction with a provided list of zrc20 - no transfer or withdraw can be processed with the paused zrc20 | fungible |
| MsgUpdateZRC20PausedStatus - Unpause | 2 | Unpause a paused zrc20 | fungible |
| MsgAddToOutTxTracker | 1 | Add an outbound tx to the onchain tx tracked (msg to be removed) | crosschain |
| MsgRemoveFromOutTxTracker | 1 | Remove an outbound tx from the onchain tx tracked (msg to be removed) | crosschain |
| MsgWhitelistERC20 | 1 | Whitelist an ERC20 on an external chain so a zrc20 can be created | crosschain |
| MsgUpdateCrosschainFlags - Disabling outbounds/inbounds | 1 | Disable inbound txs or outbound txs to be observed on ZetaChain | observer |
| MsgUpdateCrosschainFlags - Other actions | 2 | Other actions for crosschain flags are:
Enabling disabled outbound/inbound
Changing parameters for gas price increase mechanism | observer |
| MsgUpdateKeygen | 1 | Update the block height of the keygen and sets the status to PendingKeygen | observer |
| MsgUpdateTssAddress | 2 | Update the address of the TSS | crosschain |
| MsgDeployFungibleCoinZRC20 | 2 | Deploy a new ZRC20 (gas token or ERC20) | fungible |
| MsgRemoveForeignCoin | 2 | Remove a foreign coin object from ZetaChain.
The ZRC20 will still exist on ZetaChain and can still be traded but cross-chain capabilities will no longer be available for it | fungible |
| MsgUpdateZRC20LiquidityCap | 2 | Set or remove a liquidity cap for a zrc20, the maximum supply that can exist for the ZRC20 on ZetaChain | fungible |
| MsgUpdateContractBytecode | 2 | Update the bytecode of a ZRC20 or a system contract | fungible |
| MsgUpdateSystemContract | 2 | Update the address of the system contract | fungible |
| MsgUpdateZRC20WithdrawFee | 2 | Update the withdraw protocol fees for a ZRC20.
A flat fee paid for each withdraw.  | fungible |
| MsgUpdateCoreParams | 2 | Updates core parameters for a specific chain.
Core parameters include: confirmation count, outbound transaction schedule interval, ZETA token,
connector and ERC20 custody contract addresses, etc. | observer |
| MsgMigrateTssFunds | 2 | Send a tx to an external chain to migrate tss funds from one address to another one | crosschain |
| MsgAddToInTxTracker | 1 | Add an inbound tx to the onchain tx tracked (msg to be removed) | crosschain |
| MsgAddObserver | 2 | Add a new observer | observer |
| MsgUpdateObserver | 2 | Updade the operator of an existing observer | observer |