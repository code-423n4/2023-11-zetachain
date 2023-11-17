import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import {
  ZETA_INITIAL_SUPPLY,
  ZETA_TOKEN_SALT_NUMBER_ETH,
  ZETA_TOKEN_SALT_NUMBER_NON_ETH,
} from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { ZetaEth__factory, ZetaNonEth__factory } from "../../../typechain-types";

export async function deterministicDeployZetaToken() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const saltNumber = isEthNetworkName(network.name) ? ZETA_TOKEN_SALT_NUMBER_ETH : ZETA_TOKEN_SALT_NUMBER_NON_ETH;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

  let constructorTypes;
  let constructorArgs;
  let contractBytecode;

  if (isEthNetworkName(network.name)) {
    constructorTypes = ["address", "uint256"];
    constructorArgs = [DEPLOYER_ADDRESS, ZETA_INITIAL_SUPPLY.toString()];
    contractBytecode = ZetaEth__factory.bytecode;
  } else {
    constructorTypes = ["address", "address"];
    constructorArgs = [tssAddress, tssUpdaterAddress];
    contractBytecode = ZetaNonEth__factory.bytecode;
  }

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  console.log("Deployed zetaToken. Address:", address);
  console.log("Constructor Args", constructorArgs);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaToken()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
