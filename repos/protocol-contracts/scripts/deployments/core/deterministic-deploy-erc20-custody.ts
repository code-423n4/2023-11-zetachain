import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import {
  ERC20_CUSTODY_SALT_NUMBER_ETH,
  ERC20_CUSTODY_SALT_NUMBER_NON_ETH,
  ERC20_CUSTODY_ZETA_FEE,
  ERC20_CUSTODY_ZETA_MAX_FEE,
} from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { ERC20Custody__factory } from "../../../typechain-types";

export const deterministicDeployERC20Custody = async () => {
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

  const saltNumber = isEthNetworkName(network.name) ? ERC20_CUSTODY_SALT_NUMBER_ETH : ERC20_CUSTODY_SALT_NUMBER_NON_ETH;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const zetaFee = ERC20_CUSTODY_ZETA_FEE;
  const zetaMaxFee = ERC20_CUSTODY_ZETA_MAX_FEE;

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

  const constructorTypes = ["address", "address", "uint256", "uint256", "address"];
  const constructorArgs = [tssAddress, tssUpdaterAddress, zetaFee.toString(), zetaMaxFee.toString(), zetaTokenAddress];
  const contractBytecode = ERC20Custody__factory.bytecode;

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  console.log("Deployed ERC20 Custody. Address:", address);
  console.log("Constructor Args", constructorArgs);
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployERC20Custody()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
