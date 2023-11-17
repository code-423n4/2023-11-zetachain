__Goal: Steal funds\n__

- **Exploit Inbound \n Transaction Processing**\n\n
  + Inject malicious\n code into a \n contract triggered by a CCTX \n\n
  + Submit transactions\n with manipulated \n contract addresses to redirect \n funds. \n\n
  + Use incorrect \n destination data to \n reroute funds to an attacker-controlled \n address. \n\n
  + Modify transaction \n amounts in CCTXs \n to withdraw unauthorized token \n amounts. \n\n

- **Manipulate \n Outbound Transactions**\n\n
  + Intercept and alter \n outbound transaction data \n to misappropriate funds.\n\n
  + Compromise observer \n nodes to validate \n fraudulent transactions.\n\n
  + Hack into the TSS \n keysign process \n to authorize illegitimate withdrawals.\n\n

- **Control Over \n Network Consensus**\n\n
  + Gain majority control \n of consensus nodes \n to validate false transactions.\n\n
  + Sybil attack to create \n multiple fake identities \n and influence network decisions.\n\n

- **Disrupt Transaction \n Verification**\n\n
  + Send transactions \n with invalid signatures \n to create verification confusion.\n\n
  + Flood the network \n with transactions \n to overwhelm validation processes.\n\n

- **Abuse Smart Contract \n Vulnerabilities**\n\n
  + Identify and exploit \n smart contract bugs \n to drain funds from contracts.\n\n
  + Initiate transactions \n that trigger contract functions \n to release funds improperly.\n\n

- **Undermine System \n Configuration**\n\n
  + Exploit misconfigurations \n in node software \n to gain unauthorized access.\n\n
  + Utilize system \n updates or patches \n to introduce vulnerabilities.\n\n

- **Market Manipulation \n for Economic Gain**\n\n
  + Conduct pump and dump schemes \n to inflate token value and sell at a peak.\n\n
  + Manipulate token \n liquidity to trigger slippage \n and acquire tokens below market value.\n\n

- **Perform Double \n Spend Attacks**\n\n
  + Exploit transaction malleability \n to spend the same \n funds multiple times.\n\n
  + Race attack by broadcasting \n two conflicting transactions \n simultaneously.\n\n

- Double voucher submission
- Time Out 
  - Submit after timeout window
  - Force drop tx
- Spoofing 
  - sender
  - chain id  
  - funds 
  - receiver 
- Block reorganizations
  - [x] Waiting for confirmations
- Bad cryptography 
  - [ ] Voting used on many things
  - merkle tree verification problems
  - cryptographic signatures
  - nonce issues
  - Hash collisions 
  - Replay attacks
- Chain ID confusion 
  - [x] chains can only be added by admins
- Transaction ordering
  - [x] Cosmos is a FIFO queue
- Event/log truncation
- Event status mishandling 
  - Illegal state transitions i.e. inbound to reverted_failure
- Denomination issues 
  - [ ] only settable by admin 
  - Cross-chain mishandling still possible
- Asset inflation
  - Increasing amount of tokens being sent
- Precision issues
- Bad book keeping
  - Double counting 
  - Missing state updates
- Price manipulation on assets
- Coin type confusion on relayers
  - Zeta to gas
  - ERC20 to Zeta
  - Illegal use of 'cmd' coin
- Supply issues
  - Mint too many
  - Burn too much 
