import { Contract } from "ethers";
import { network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib/address.tools";

import { deployZetaConnectorEth, deployZetaConnectorNonEth, isEthNetworkName } from "../../../lib/contracts.helpers";

export async function deployZetaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const zetaTokenAddress = getAddress("zetaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  let contract: Contract;
  console.log(`Deploying ZetaConnector to ${network.name}`);

  if (isEthNetworkName(network.name)) {
    contract = await deployZetaConnectorEth({
      args: [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress],
    });
  } else {
    contract = await deployZetaConnectorNonEth({
      args: [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress],
    });
  }

  console.log("Deployed ZetaConnector. Address:", contract.address);
  return contract.address;
}
