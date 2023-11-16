import { AddressZero } from "@ethersproject/constants";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SystemContract, ZRC20 } from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import { FUNGIBLE_MODULE_ADDRESS } from "./test.helpers";
const hre = require("hardhat");

describe("ZRC20 tests", () => {
  let ZRC20Contract: ZRC20;
  let systemContract: SystemContract;
  let owner, fungibleModuleSigner: SignerWithAddress;
  let addrs: SignerWithAddress[];

  beforeEach(async () => {
    [owner, ...addrs] = await ethers.getSigners();

    // Impersonate the fungible module account
    await hre.network.provider.request({
      method: "hardhat_impersonateAccount",
      params: [FUNGIBLE_MODULE_ADDRESS],
    });

    // Get a signer for the fungible module account
    fungibleModuleSigner = await ethers.getSigner(FUNGIBLE_MODULE_ADDRESS);
    hre.network.provider.send("hardhat_setBalance", [FUNGIBLE_MODULE_ADDRESS, parseEther("1000000").toHexString()]);

    const SystemContractFactory = await ethers.getContractFactory("SystemContractMock");
    systemContract = (await SystemContractFactory.deploy(AddressZero, AddressZero, AddressZero)) as SystemContract;

    const ZRC20Factory = await ethers.getContractFactory("ZRC20");
    ZRC20Contract = (await ZRC20Factory.connect(fungibleModuleSigner).deploy(
      "TOKEN",
      "TKN",
      18,
      1,
      1,
      0,
      systemContract.address
    )) as ZRC20;
  });

  describe("ZRC20Contract", () => {
    it("Should return name", async () => {
      const name = await ZRC20Contract.name();
      expect(name).to.equal("TOKEN");
    });

    it("Should return symbol", async () => {
      const symbol = await ZRC20Contract.symbol();
      expect(symbol).to.equal("TKN");
    });

    it("Should return decimals", async () => {
      const decimals = await ZRC20Contract.decimals();
      expect(decimals).to.equal(18);
    });

    it("Should return chainId", async () => {
      const chainId = await ZRC20Contract.CHAIN_ID();
      expect(chainId).to.equal(1);
    });

    it("Should return coinType", async () => {
      const coinType = await ZRC20Contract.COIN_TYPE();
      expect(coinType).to.equal(1);
    });

    it("Should return gasLimit", async () => {
      const gasLimit = await ZRC20Contract.GAS_LIMIT();
      expect(gasLimit).to.equal(0);
    });

    it("Should return systemContractAddress", async () => {
      const systemContractAddress = await ZRC20Contract.SYSTEM_CONTRACT_ADDRESS();
      expect(systemContractAddress).to.equal(systemContract.address);
    });

    it("Should return totalSupply", async () => {
      const totalSupply = await ZRC20Contract.totalSupply();
      expect(totalSupply).to.equal(0);
    });

    it("Should return balanceOf", async () => {
      const balanceOf = await ZRC20Contract.balanceOf(owner.address);
      expect(balanceOf).to.equal(0);
    });

    it("Should return allowance", async () => {
      const allowance = await ZRC20Contract.allowance(owner.address, owner.address);
      expect(allowance).to.equal(0);
    });

    it("Should deposit if is called by fungible module", async () => {
      const deposit = await ZRC20Contract.connect(fungibleModuleSigner).deposit(owner.address, parseEther("1"));
      expect(deposit).to.emit(ZRC20Contract, "Deposit").withArgs(owner.address, parseEther("1"));
    });

    it("Should not deposit if is not called by fungible module", async () => {
      const depositTx = ZRC20Contract.connect(addrs[0]).deposit(owner.address, parseEther("1"));
      await expect(depositTx).to.be.revertedWith("InvalidSender");
    });

    it("Should transfer", async () => {
      await ZRC20Contract.connect(fungibleModuleSigner).deposit(owner.address, parseEther("1"));
      const initialBalance = await ZRC20Contract.balanceOf(addrs[0].address);

      const transfer = await ZRC20Contract.connect(owner).transfer(addrs[0].address, parseEther("1"));
      expect(transfer).to.emit(ZRC20Contract, "Transfer").withArgs(owner.address, addrs[0].address, parseEther("1"));

      const finalBalance = await ZRC20Contract.balanceOf(addrs[0].address);
      expect(finalBalance.sub(initialBalance)).to.equal(parseEther("1"));
    });
  });
});
