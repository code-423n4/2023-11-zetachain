import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ZetaNonEth__factory as ZetaNonEthFactory } from "../../typechain-types";

export async function setZetaAddresses(connectorAddress: string, zetaTokenAddress: string) {
  const [owner] = await ethers.getSigners();

  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  if (owner.address !== tssAddress && owner.address !== tssUpdaterAddress) {
    console.log("Only TSS or TSS Updater can set Zeta addresses.");
    console.log("Please execute this step with a valid account.");
    return;
  }

  const [tssSigner] = await ethers.getSigners();

  const contract = ZetaNonEthFactory.connect(zetaTokenAddress, tssSigner);

  console.log("Updating");
  console.log("connectorAddress", connectorAddress);
  console.log("zetaTokenAddress", zetaTokenAddress);
  const tx = await contract.updateTssAndConnectorAddresses(tssAddress, connectorAddress);
  await tx.wait();
  console.log("Updated");
}
