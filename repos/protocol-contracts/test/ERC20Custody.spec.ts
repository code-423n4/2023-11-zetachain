import { MaxUint256 } from "@ethersproject/constants";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { AttackerContract, ERC20Custody, ERC20Mock, ZetaEth } from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import { deployZetaEth } from "../lib/contracts.helpers";

const ZETA_FEE = parseEther("0.01");
const ZETA_MAX_FEE = parseEther("0.05");

describe("ERC20Custody tests", () => {
  let ZRC20CustodyContract: ERC20Custody;
  let zetaTokenEthContract: ZetaEth;
  let randomTokenContract: ERC20Mock;
  let owner: SignerWithAddress;
  let tssUpdater: SignerWithAddress;
  let tssSigner: SignerWithAddress;
  let addrs: SignerWithAddress[];

  beforeEach(async () => {
    [owner, tssUpdater, tssSigner, ...addrs] = await ethers.getSigners();

    zetaTokenEthContract = await deployZetaEth({
      args: [tssUpdater.address, 100_000],
    });

    await zetaTokenEthContract.connect(tssUpdater).transfer(owner.address, parseEther("1000"));

    const ERC20Factory = await ethers.getContractFactory("ERC20Mock");
    randomTokenContract = (await ERC20Factory.deploy(
      "RandomToken",
      "RT",
      owner.address,
      parseEther("1000000")
    )) as ERC20Mock;

    const ERC20CustodyFactory = await ethers.getContractFactory("ERC20Custody");
    ZRC20CustodyContract = (await ERC20CustodyFactory.deploy(
      tssSigner.address,
      tssUpdater.address,
      ZETA_FEE,
      ZETA_MAX_FEE,
      zetaTokenEthContract.address
    )) as ERC20Custody;
  });

  describe("ERC20Custody", () => {
    it("Should update the tss address", async () => {
      await ZRC20CustodyContract.connect(tssUpdater).updateTSSAddress(addrs[0].address);
      const tssAddress = await ZRC20CustodyContract.TSSAddress();
      expect(tssAddress).to.equal(addrs[0].address);
    });

    it("Should revert if updateTSSAddress not called by tssUpdater", async () => {
      const tx = ZRC20CustodyContract.connect(owner).updateTSSAddress(addrs[0].address);
      await expect(tx).to.be.revertedWith("InvalidTSSUpdater");
    });

    it("Should update zeta fee", async () => {
      await ZRC20CustodyContract.connect(tssSigner).updateZetaFee(parseEther("0.02"));
      const zetaFee = await ZRC20CustodyContract.zetaFee();
      expect(zetaFee).to.equal(parseEther("0.02"));
    });

    it("Should revert if updateZetaFee not called by tssUpdater", async () => {
      const tx = ZRC20CustodyContract.connect(tssUpdater).updateZetaFee(parseEther("0.02"));
      await expect(tx).to.be.revertedWith("InvalidSender");
    });

    it("Should revert if updateZetaFee is higher than max", async () => {
      const tx = ZRC20CustodyContract.connect(tssSigner).updateZetaFee(parseEther("0.07"));
      await expect(tx).to.be.revertedWith("ZetaMaxFeeExceeded");
    });

    it("Should renounce tssUpdater", async () => {
      await ZRC20CustodyContract.connect(tssUpdater).renounceTSSAddressUpdater();
      const tssUpdaterAddress = await ZRC20CustodyContract.TSSAddressUpdater();
      expect(tssUpdaterAddress).to.equal(tssSigner.address);
    });

    it("Should whitelist", async () => {
      let isWhitelisted = await ZRC20CustodyContract.whitelisted(addrs[0].address);
      expect(isWhitelisted).to.equal(false);

      await ZRC20CustodyContract.connect(tssSigner).whitelist(addrs[0].address);
      isWhitelisted = await ZRC20CustodyContract.whitelisted(addrs[0].address);
      expect(isWhitelisted).to.equal(true);
    });

    it("Should revert if whitelist is not called by tss", async () => {
      const tx = ZRC20CustodyContract.connect(tssUpdater).whitelist(addrs[0].address);
      await expect(tx).to.be.revertedWith("InvalidSender");
    });

    it("Should unwhitelist", async () => {
      let isWhitelisted = await ZRC20CustodyContract.whitelisted(addrs[0].address);
      expect(isWhitelisted).to.equal(false);

      await ZRC20CustodyContract.connect(tssSigner).whitelist(addrs[0].address);
      isWhitelisted = await ZRC20CustodyContract.whitelisted(addrs[0].address);
      expect(isWhitelisted).to.equal(true);

      await ZRC20CustodyContract.connect(tssSigner).unwhitelist(addrs[0].address);
      isWhitelisted = await ZRC20CustodyContract.whitelisted(addrs[0].address);
      expect(isWhitelisted).to.equal(false);
    });

    it("Should revert if unwhitelist is not called by tss", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(addrs[0].address);
      const tx = ZRC20CustodyContract.connect(tssUpdater).unwhitelist(addrs[0].address);
      await expect(tx).to.be.revertedWith("InvalidSender");
    });

    it("Should emit Deposit event on deposit", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);
      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);

      const amount = parseEther("1");
      await randomTokenContract.approve(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.deposit(owner.address, randomTokenContract.address, amount, "0x00");
      await expect(tx)
        .to.emit(ZRC20CustodyContract, "Deposited")
        .withArgs(owner.address.toLowerCase(), randomTokenContract.address, amount, "0x00");
    });

    it("Should emit Deposit event and do the right math if is called twice", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);
      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);

      const amount = parseEther("1");
      await randomTokenContract.approve(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.deposit(owner.address, randomTokenContract.address, amount, "0x00");
      await expect(tx)
        .to.emit(ZRC20CustodyContract, "Deposited")
        .withArgs(owner.address.toLowerCase(), randomTokenContract.address, amount, "0x00");

      const amount2 = parseEther("5");
      await randomTokenContract.approve(ZRC20CustodyContract.address, amount2);

      const tx2 = ZRC20CustodyContract.deposit(owner.address, randomTokenContract.address, amount2, "0x00");
      await expect(tx2)
        .to.emit(ZRC20CustodyContract, "Deposited")
        .withArgs(owner.address.toLowerCase(), randomTokenContract.address, amount2, "0x00");
    });

    it("Should revert deposit if is paused", async () => {
      await ZRC20CustodyContract.connect(tssSigner).pause();
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);
      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);

      const amount = parseEther("1");
      await randomTokenContract.approve(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.deposit(owner.address, randomTokenContract.address, amount, "0x00");
      await expect(tx).to.be.revertedWith("IsPaused");
    });

    it("Should revert deposit if is not whitelisted", async () => {
      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);

      const amount = parseEther("1");
      await randomTokenContract.approve(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.deposit(owner.address, randomTokenContract.address, amount, "0x00");
      await expect(tx).to.be.revertedWith("NotWhitelisted");
    });

    it("Should revert deposit if has no zeta to pay fee", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);
      await zetaTokenEthContract.connect(addrs[0]).approve(ZRC20CustodyContract.address, MaxUint256);

      const amount = parseEther("1");
      await randomTokenContract.connect(addrs[0]).approve(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.connect(addrs[0]).deposit(
        owner.address,
        randomTokenContract.address,
        amount,
        "0x00"
      );
      await expect(tx).to.be.revertedWith("ERC20: transfer amount exceeds balance");
    });

    it("Should emit Withdrawn event", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);

      const amount = parseEther("1");
      await randomTokenContract.transfer(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.connect(tssSigner).withdraw(owner.address, randomTokenContract.address, amount);
      await expect(tx)
        .to.emit(ZRC20CustodyContract, "Withdrawn")
        .withArgs(owner.address, randomTokenContract.address, amount);
    });

    it("Should revert if withdraw is not called by tss", async () => {
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);

      const amount = parseEther("1");
      await randomTokenContract.transfer(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.withdraw(owner.address, randomTokenContract.address, amount);
      await expect(tx).to.be.revertedWith("InvalidSender");
    });

    it("Should not revert withdraw if is paused", async () => {
      await ZRC20CustodyContract.connect(tssSigner).pause();
      await ZRC20CustodyContract.connect(tssSigner).whitelist(randomTokenContract.address);

      const amount = parseEther("1");
      await randomTokenContract.transfer(ZRC20CustodyContract.address, amount);

      const tx = ZRC20CustodyContract.connect(tssSigner).withdraw(owner.address, randomTokenContract.address, amount);
      await expect(tx)
        .to.emit(ZRC20CustodyContract, "Withdrawn")
        .withArgs(owner.address, randomTokenContract.address, amount);
    });

    it("Should revert withdraw if is not whitelisted", async () => {
      const amount = parseEther("1");
      const tx = ZRC20CustodyContract.connect(tssSigner).withdraw(owner.address, randomTokenContract.address, amount);
      await expect(tx).to.be.revertedWith("NotWhitelisted");
    });

    it("Should not allow reentrant calls to Deposit method", async () => {
      // Deploy AttackerContract with the address of the ZRC20CustodyContract
      const Attacker = await ethers.getContractFactory("AttackerContract");
      const attacker = (await Attacker.deploy(
        ZRC20CustodyContract.address,
        zetaTokenEthContract.address,
        1
      )) as AttackerContract;
      await attacker.deployed();

      // Whitelist token and approve as before
      await ZRC20CustodyContract.connect(tssSigner).whitelist(attacker.address);

      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);
      await zetaTokenEthContract.transfer(attacker.address, parseEther("1"));

      const amount = parseEther("1");

      // Attempt to attack by calling the deposit method reentrantly
      const attackTx = ZRC20CustodyContract.deposit(owner.address, attacker.address, amount, "0x00");

      // The test expects the attack to fail due to the reentrancy guard
      await expect(attackTx).to.be.revertedWith("ReentrancyGuard: reentrant call");
    });

    it("Should not allow reentrant calls to Witdraw method", async () => {
      // Deploy AttackerContract with the address of the ZRC20CustodyContract
      const Attacker = await ethers.getContractFactory("AttackerContract");
      const attacker = (await Attacker.deploy(
        ZRC20CustodyContract.address,
        zetaTokenEthContract.address,
        2
      )) as AttackerContract;
      await attacker.deployed();

      // Whitelist token and approve as before
      await ZRC20CustodyContract.connect(tssSigner).whitelist(attacker.address);

      await zetaTokenEthContract.approve(ZRC20CustodyContract.address, MaxUint256);
      await zetaTokenEthContract.transfer(attacker.address, parseEther("1"));

      const amount = parseEther("1");

      // Attempt to attack by calling the deposit method reentrantly
      const attackTx = ZRC20CustodyContract.connect(tssSigner).withdraw(owner.address, attacker.address, amount);

      // The test expects the attack to fail due to the reentrancy guard
      await expect(attackTx).to.be.revertedWith("ReentrancyGuard: reentrant call");
    });
  });
});
