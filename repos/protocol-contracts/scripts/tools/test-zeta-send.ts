import { networks } from "@zetachain/networks";
import { AbiCoder } from "ethers/lib/utils";
import { ethers, network } from "hardhat";

import { getAddress, isProtocolNetworkName, ZetaProtocolNetwork } from "../../lib/address.tools";
import { ZetaConnectorEth__factory as ZetaConnectorEthFactory } from "../../typechain-types";

const encoder = new AbiCoder();

export const getChainId = (network: ZetaProtocolNetwork): number => {
  //@ts-ignore
  return networks[network].chain_id;
};

async function main() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();

  const connectorAddress = getAddress("connector", network.name);

  const factory = (await ethers.getContractFactory("ZetaConnectorEth")) as ZetaConnectorEthFactory;
  const contract = factory.attach(connectorAddress);

  console.log(`Sending To ${accounts[0].address}`);
  await (
    await contract.send({
      destinationAddress: encoder.encode(["address"], [accounts[0].address]),
      destinationChainId: getChainId("bsc_testnet"),
      destinationGasLimit: 1_000_000,
      message: encoder.encode(["address"], [accounts[0].address]),
      zetaParams: [],
      zetaValueAndGas: "10000000000000000000",
    })
  ).wait();
  console.log("Sent");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
