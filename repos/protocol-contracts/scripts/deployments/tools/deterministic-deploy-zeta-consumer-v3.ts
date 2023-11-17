import { ethers, network } from "hardhat";
import { getAddress, getNonZetaAddress, isProtocolNetworkName } from "lib";

import { ZetaTokenConsumerUniV3__factory } from "../../../typechain-types";

export async function deterministicDeployZetaConsumer() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const zetaTokenAddress = getAddress("zetaToken", network.name);

  const uniswapV3Router = getNonZetaAddress("uniswapV3Router", network.name);
  const uniswapV3Factory = getNonZetaAddress("uniswapV3Factory", network.name);
  const WETH9Address = getNonZetaAddress("weth9", network.name);

  const zetaPoolFee = 500;
  const tokenPoolFee = 3000;

  console.log([zetaTokenAddress, uniswapV3Router, uniswapV3Factory, WETH9Address, zetaPoolFee, tokenPoolFee]);

  const Factory = new ZetaTokenConsumerUniV3__factory(signer);
  const contract = await Factory.deploy(
    zetaTokenAddress,
    uniswapV3Router,
    uniswapV3Factory,
    WETH9Address,
    zetaPoolFee,
    tokenPoolFee
  );
  await contract.deployed();
  const address = contract.address;

  console.log("Deployed ZetaConsumer. Address:", address);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployZetaConsumer()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
