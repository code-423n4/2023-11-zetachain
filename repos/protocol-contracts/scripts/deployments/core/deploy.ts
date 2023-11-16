import { network } from "hardhat";

import { isProtocolNetworkName } from "../../../lib/address.tools";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { setZetaAddresses } from "../../tools/set-zeta-token-addresses";
import { deployZetaConnector } from "./deploy-zeta-connector";
import { deployZetaToken } from "./deploy-zeta-token";

const networkName = network.name;

async function main() {
  if (!isProtocolNetworkName(networkName)) throw new Error("Invalid network name");

  const zetaTokenAddress = await deployZetaToken();
  const connectorAddress = await deployZetaConnector();

  /**
   * @description The Eth implementation of Zeta token doesn't need any address
   */
  if (isEthNetworkName(network.name)) return;

  /**
   * @description Avoid setting Zeta addresses for local network,
   * since it must be done after starting the local Zeta node
   */
  await setZetaAddresses(connectorAddress, zetaTokenAddress);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
