import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ZETA_CONNECTOR_SALT_NUMBER_ETH, ZETA_CONNECTOR_SALT_NUMBER_NON_ETH } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { ZetaConnectorEth__factory, ZetaConnectorNonEth__factory } from "../../../typechain-types";

export async function deterministicDeployZetaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const zetaTokenAddress = getAddress("zetaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const saltNumber = isEthNetworkName(network.name)
    ? ZETA_CONNECTOR_SALT_NUMBER_ETH
    : ZETA_CONNECTOR_SALT_NUMBER_NON_ETH;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  const constructorTypes = ["address", "address", "address", "address"];
  const constructorArgs = [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];

  let contractBytecode;
  if (isEthNetworkName(network.name)) {
    contractBytecode = ZetaConnectorEth__factory.bytecode;
  } else {
    contractBytecode = ZetaConnectorNonEth__factory.bytecode;
  }

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  console.log("Deployed ZetaConnector. Address:", address);
  console.log("Constructor Args", constructorArgs);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaConnector()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
