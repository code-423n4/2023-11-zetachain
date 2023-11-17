import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getZetaFactoryNonEth } from "../../lib/contracts.helpers";

async function updateZetaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const [, tssUpdaterSigner] = await ethers.getSigners();

  const zetaTokenAddress = getAddress("zetaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const connectorAddress = getAddress("connector", network.name);

  const contract = (
    await getZetaFactoryNonEth({ deployParams: null, existingContractAddress: zetaTokenAddress })
  ).connect(tssUpdaterSigner);

  await (await contract.updateTssAndConnectorAddresses(tssAddress, connectorAddress)).wait();

  console.log(`Updated TSS address to ${tssAddress}.`);
  console.log(`Updated Connector address to ${connectorAddress}.`);
}

updateZetaConnector()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
