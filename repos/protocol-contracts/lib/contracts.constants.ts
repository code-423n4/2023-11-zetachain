import { parseEther } from "ethers/lib/utils";

export const ZETA_INITIAL_SUPPLY = 2_100_000_000;

export const MAX_ETH_ADDRESS = "0xffffffffffffffffffffffffffffffffffffffff";

// dev: this values should be calculated using get-salt script
export const ZETA_TOKEN_SALT_NUMBER_ETH = "84108";
export const ZETA_TOKEN_SALT_NUMBER_NON_ETH = "29265";

// dev: this values should be calculated using get-salt script
export const ZETA_CONNECTOR_SALT_NUMBER_ETH = "1414";
export const ZETA_CONNECTOR_SALT_NUMBER_NON_ETH = "71733";

export const ERC20_CUSTODY_SALT_NUMBER_ETH = "87967";
export const ERC20_CUSTODY_SALT_NUMBER_NON_ETH = "195084";
export const ERC20_CUSTODY_ZETA_FEE = "0";
export const ERC20_CUSTODY_ZETA_MAX_FEE = parseEther("1000");

export const ZETA_CONSUMER_SALT_NUMBER = "0";
