import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getZetaConnectorEth, getZetaConnectorNonEth, isEthNetworkName } from "../../lib/contracts.helpers";

async function sendGas() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const [, tssUpdaterSigner] = await ethers.getSigners();

  const tssAddress = getAddress("tss", network.name);
  const connectorAddress = getAddress("connector", network.name);

  const newTssAddress = "0x0000000000000000000000000000000000000000";

  if (isEthNetworkName(network.name)) {
    const contract = (
      await getZetaConnectorEth({ deployParams: null, existingContractAddress: connectorAddress })
    ).connect(tssUpdaterSigner);

    await (await contract.updateTssAddress(newTssAddress)).wait();
  } else {
    const contract = (
      await getZetaConnectorNonEth({ deployParams: null, existingContractAddress: connectorAddress })
    ).connect(tssUpdaterSigner);

    await (await contract.updateTssAddress(newTssAddress)).wait();
  }

  console.log(`Updated TSS address from ${tssAddress} to ${newTssAddress}.`);
}

sendGas()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
