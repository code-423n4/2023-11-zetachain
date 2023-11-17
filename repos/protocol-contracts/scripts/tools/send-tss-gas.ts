import { ethers, network } from "hardhat";

import { getAddress, isProtocolNetworkName } from "../../lib/address.tools";

async function sendGas() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const [signer] = await ethers.getSigners();

  const tssAddress = getAddress("tss", network.name);

  const sendGasTx = {
    from: signer.address,
    to: tssAddress,
    value: ethers.utils.parseEther("1.0"),
  };
  await signer.sendTransaction(sendGasTx);
  console.log(`Sent 1.0 Ether from ${signer.address} to TSS address (${tssAddress}).`);
}

sendGas()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
