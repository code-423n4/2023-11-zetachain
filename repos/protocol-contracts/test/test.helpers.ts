import { ZetaTokenConsumer__factory } from "@typechain-types";
import { BigNumber, ContractReceipt } from "ethers";

export const FUNGIBLE_MODULE_ADDRESS = "0x735b14BB79463307AAcBED86DAf3322B1e6226aB";

export const parseZetaConsumerLog = (logs: ContractReceipt["logs"]) => {
  const iface = ZetaTokenConsumer__factory.createInterface();

  const eventNames = logs.map((log) => {
    try {
      const parsedLog = iface.parseLog(log);

      return parsedLog.name;
    } catch (e) {
      return "NO_ZETA_LOG";
    }
  });

  return eventNames;
};
