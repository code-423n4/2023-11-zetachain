import { ZetaProtocolNetwork } from "./address.tools";

export declare type TestAddress = "dai" | "usdc";

export const getTestAddress = (address: TestAddress, networkName: ZetaProtocolNetwork): string => {
  if (networkName !== "eth_mainnet") throw new Error("Invalid network name");
  if (address === "dai") {
    return "0x6b175474e89094c44da98b954eedeac495271d0f";
  } else if (address === "usdc") {
    return "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48";
  }

  throw new Error(`Unknown external address ${address} for network ${networkName}`);
};
