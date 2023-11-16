import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";
import { ZetaConnectorEth__factory, ZetaConnectorNonEth__factory } from "../../../typechain-types";

const MAX_ITERATIONS = BigNumber.from(100000);

export async function deterministicDeployGetSaltZetaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const zetaTokenAddress = getAddress("zetaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  // @todo: decide which address use as pauser
  const constructorTypes = ["address", "address", "address", "address"];
  const constructorArgs = [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];

  let contractBytecode;

  if (isEthNetworkName(network.name)) {
    contractBytecode = ZetaConnectorEth__factory.bytecode;
  } else {
    contractBytecode = ZetaConnectorNonEth__factory.bytecode;
  }

  calculateBestSalt(MAX_ITERATIONS, DEPLOYER_ADDRESS, constructorTypes, constructorArgs, contractBytecode);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployGetSaltZetaConnector()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
