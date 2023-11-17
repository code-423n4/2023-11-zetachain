import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getZetaFactoryEth, getZetaFactoryNonEth, isEthNetworkName } from "../../lib/contracts.helpers";

const approvalAmount = ethers.utils.parseEther("10000000.0");

export async function setTokenApproval() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const zetaTokenAddress = getAddress("zetaToken", network.name);
  const connectorAddress = getAddress("connector", network.name);

  let contract;
  if (isEthNetworkName(network.name)) {
    contract = await getZetaFactoryEth({ deployParams: null, existingContractAddress: zetaTokenAddress });
  } else {
    contract = await getZetaFactoryNonEth({ deployParams: null, existingContractAddress: zetaTokenAddress });
  }

  let tx = await contract.approve(connectorAddress, approvalAmount);
  tx.wait();

  console.log(`Approved Connector Contract ${connectorAddress} for ${approvalAmount} `);
}

setTokenApproval()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
