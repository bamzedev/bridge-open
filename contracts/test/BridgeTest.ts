import { expect } from "chai";
import { Contract, Signer, Signature } from "ethers";
import { signERC2612Permit } from "eth-permit";
import hre from "hardhat";
import { ethers } from "hardhat";

let validator: Signer, validator2: Signer, feeCollector: Signer, user: Signer;
let token: Contract;
let wToken: Contract;
let bridge: Contract;
let mintTx: any;
let transactionHashForMint: string;
let transactionHashForUnlock: string;

const mint = async (transactionHash: string, messageHash: string, token: Contract, validatorA: Signer, validatorB: Signer) => {
  const arrayfiedHash = ethers.utils.arrayify(messageHash);
  const signedMessage = await validatorA.signMessage(arrayfiedHash);
  const signedMessage2 = await validatorB.signMessage(arrayfiedHash);
  const sig = ethers.utils.splitSignature(signedMessage);
  const sig2 = ethers.utils.splitSignature(signedMessage2);
  let sigs: Signature[] = [sig, sig2];
  let tx = await bridge.connect(user).mintTokens(token.address, "Wrapped Test Token", "WTT", 100, transactionHash, [sigs[0].v, sigs[1].v], [sigs[0].r, sigs[1].r], [sigs[0].s, sigs[1].s]);
  let receipt = await tx.wait();
  let wrappedTokenEvent = receipt.events.filter((x: any) => {
    return x.event == "NewTokenDeployed";
  });
  if (wrappedTokenEvent.length > 0) {
    const WToken = await hre.ethers.getContractFactory("WrappedToken");
    let wt = WToken.attach(wrappedTokenEvent[0].args.tokenContract);
    return {
      wToken: wt.attach(wrappedTokenEvent[0].args.tokenContract),
      tx: tx,
    };
  } else {
    return { wToken, tx: tx };
  }
};

const unlock = async (transactionHash: string, messageHash: string, token: Contract, validatorA: Signer, validatorB: Signer) => {
  const arrayfiedHash = ethers.utils.arrayify(messageHash);
  const signedMessage = await validatorA.signMessage(arrayfiedHash);
  const signedMessage2 = await validatorB.signMessage(arrayfiedHash);
  const sig = ethers.utils.splitSignature(signedMessage);
  const sig2 = ethers.utils.splitSignature(signedMessage2);
  let sigs: Signature[] = [sig, sig2];
  let tx = await bridge.connect(user).unlockTokens(token.address, 100, transactionHash, [sigs[0].v, sigs[1].v], [sigs[0].r, sigs[1].r], [sigs[0].s, sigs[1].s]);
  await tx.wait();
  return tx;
};

describe("Bridge", async () => {
  beforeEach(async () => {
    [validator, validator2, feeCollector, user] = await ethers.getSigners();
    transactionHashForMint = "0x06bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f";
    transactionHashForUnlock = "0x16bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f";
    const Token = await hre.ethers.getContractFactory("WrappedToken");
    token = await Token.deploy("Test Token", "TT");
    const Bridge = await hre.ethers.getContractFactory("Bridge");
    bridge = await Bridge.deploy([await validator.getAddress(), await validator2.getAddress()], await feeCollector.getAddress(), 1);
    await token.mintTo(await user.getAddress(), 100);
    await token.mintTo(await bridge.address, 100);
    expect(await token.balanceOf(await user.getAddress())).to.equal(100);
    const mintMessageHash = ethers.utils.solidityKeccak256(
      ["uint256", "string", "address", "string", "string", "uint256", "address"],
      [(await ethers.provider.getNetwork()).chainId, transactionHashForMint, token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
    );
    const result = await mint(transactionHashForMint, mintMessageHash, token, validator, validator2);
    wToken = result.wToken;
    mintTx = result.tx;
  });

  describe("mint", async () => {
    it("Should mint without deployed wrapped token contract", async () => {
      await expect(mintTx).to.emit(bridge, "NewTokenDeployed").withArgs(wToken.address);
      await expect(mintTx)
        .to.emit(bridge, "TokenMint")
        .withArgs(await user.getAddress(), wToken.address, 100);
      await expect(await wToken.balanceOf(await user.getAddress())).to.equal(100);
    });
    it("Should mint with already deployed wrapped token contract", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x26bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      let result = await mint("0x26bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2);
      let mintTx = result.tx;
      await expect(mintTx)
        .to.emit(bridge, "TokenMint")
        .withArgs(await user.getAddress(), wToken.address, 100);
      await expect(mintTx).to.not.emit(bridge, "NewTokenDeployed");
      await expect(await wToken.balanceOf(await user.getAddress())).to.equal(200);
    });
    it("Should not mint again with already proccessed transaction", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForMint, token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      await expect(mint(transactionHashForMint, messageHash, token, validator, validator2)).to.be.revertedWith("Transaction already processed");
    });
    it("Should not mint with wrong chainId", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [
          (await ethers.provider.getNetwork()).chainId + 1,
          "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f",
          token.address,
          "Wrapped Test Token",
          "WTT",
          100,
          await user.getAddress(),
        ]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with wrong token address", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", wToken.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with wrong token name", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrong Test Token", "WTT", 100, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with wrong token symbol", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WRONG", 100, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with wrong amount", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WTT", 101, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with wrong receiver address", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [
          (await ethers.provider.getNetwork()).chainId,
          "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f",
          token.address,
          "Wrapped Test Token",
          "WTT",
          100,
          await validator.getAddress(),
        ]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint with a signature from not a validator", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, user)).to.be.revertedWith("Wrong signature");
    });
    it("Should not mint without 2 different validator signatures", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      await expect(mint("0x36bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", messageHash, token, validator, validator)).to.be.revertedWith("Same signature");
    });
    it("Should not mint with only 1 signature", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "string", "string", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, "0x46bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f", token.address, "Wrapped Test Token", "WTT", 100, await user.getAddress()]
      );
      const arrayfiedHash = ethers.utils.arrayify(messageHash);
      const signedMessage = await validator.signMessage(arrayfiedHash);
      const sig = ethers.utils.splitSignature(signedMessage);
      let sigs: Signature[] = [sig];
      await expect(bridge.connect(user).mintTokens(token.address, "Wrapped Test Token", "WTT", 100, "0x46bfc9e0f64e8b2476863bf51e22d12a698ee6cefd26226ff6f0771301ab2e4f",
       [sigs[0].v], [sigs[0].r], [sigs[0].s])).to.be.revertedWith("Need at least 2 signatures")

    });
  });

  describe("lock", async () => {
    it("Should lock tokens", async () => {
      const result = await signERC2612Permit(user, token.address, await user.getAddress(), bridge.address, 1);
      await expect(
        bridge.lock(await user.getAddress(), token.address, 1, result.deadline, result.v, result.r, result.s, {
          value: 1,
        })
      )
        .to.emit(bridge, "TokenLock")
        .withArgs(await user.getAddress(), token.address, 1);
      expect(await token.balanceOf(await user.getAddress())).to.equal(99);
      expect(await token.balanceOf(await bridge.address)).to.equal(101);
    });
    it("Should not lock with wrong permit signature", async () => {
      const result = await signERC2612Permit(validator, token.address, await user.getAddress(), bridge.address, 1);
      await expect(bridge.lock(await user.getAddress(), token.address, 1, result.deadline, result.v, result.r, result.s)).to.be.reverted;
    });
    it("Should not lock with not enough ether value", async () => {
      const result = await signERC2612Permit(user, token.address, await user.getAddress(), bridge.address, 1);
      await expect(bridge.lock(await user.getAddress(), token.address, 1, result.deadline, result.v, result.r, result.s)).to.be.revertedWith("Wrong msg value");
    });
  });

  describe("burn", async () => {
    it("Should burn tokens", async () => {
      const burnMessageHash = ethers.utils.solidityKeccak256(["address", "address", "address", "uint256"], [await user.getAddress(), bridge.address, wToken.address, 1]);

      console.log("burnMessageHash: " + burnMessageHash);
      const burnArrayfiedHash = ethers.utils.arrayify(burnMessageHash);
      const burnSignedMessage = await user.signMessage(burnArrayfiedHash);
      const burnSig = ethers.utils.splitSignature(burnSignedMessage);
      let tx = await bridge.connect(user).burn(await user.getAddress(), wToken.address, 1, burnSig.v, burnSig.r, burnSig.s, { value: 1 });
      await tx.wait(1);

      expect(await wToken.balanceOf(await user.getAddress())).to.equal(99);
      await expect(tx)
        .to.emit(bridge, "TokenBurn")
        .withArgs(await user.getAddress(), token.address, 1);
    });
    it("Should not burn tokens with signature from another address", async () => {
      const burnMessageHash = ethers.utils.solidityKeccak256(["address", "address", "address", "uint256"], [await validator.getAddress(), bridge.address, wToken.address, 1]);

      console.log("burnMessageHash: " + burnMessageHash);
      const burnArrayfiedHash = ethers.utils.arrayify(burnMessageHash);
      const burnSignedMessage = await user.signMessage(burnArrayfiedHash);
      const burnSig = ethers.utils.splitSignature(burnSignedMessage);
      await expect(bridge.connect(user).burn(await validator.getAddress(), wToken.address, 1, burnSig.v, burnSig.r, burnSig.s, { value: 1 })).to.be.revertedWith("Wrong signature");
    });
    it("Should not burn tokens with wrong amount of fee", async () => {
      const burnMessageHash = ethers.utils.solidityKeccak256(["address", "address", "address", "uint256"], [await validator.getAddress(), bridge.address, wToken.address, 1]);

      console.log("burnMessageHash: " + burnMessageHash);
      const burnArrayfiedHash = ethers.utils.arrayify(burnMessageHash);
      const burnSignedMessage = await user.signMessage(burnArrayfiedHash);
      const burnSig = ethers.utils.splitSignature(burnSignedMessage);
      await expect(bridge.connect(user).burn(await validator.getAddress(), wToken.address, 1, burnSig.v, burnSig.r, burnSig.s, { value: 0 })).to.be.revertedWith("Wrong msg value");
    });
  });

  describe("unlock", async () => {
    it("Should unlock tokens", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );

      let tx = await unlock(transactionHashForUnlock, messageHash, token, validator, validator2);
      await tx.wait(1);

      expect(await token.balanceOf(await user.getAddress())).to.equal(200);
      expect(await token.balanceOf(bridge.address)).to.equal(0);
      await expect(tx)
        .to.emit(bridge, "TokenUnlock")
        .withArgs(await user.getAddress(), token.address, 100);
    });
    it("Should not unlock tokens with already proccessed transaction", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );
      await unlock(transactionHashForUnlock, messageHash, token, validator, validator2);
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator2)).to.be.revertedWith("Transaction already processed");
    });
    it("Should not unlock tokens with wrong chainId", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId + 1, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not unlock tokens with wrong token address", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, wToken.address, 100, await user.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not unlock tokens with wrong token amount", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 101, await user.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not unlock tokens with wrong receiver address", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await validator.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator2)).to.be.revertedWith("Wrong signature");
    });
    it("Should not unlock tokens with a signature from not a validator", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, user)).to.be.revertedWith("Wrong signature");
    });
    it("Should not unlock without 2 different validator signatures", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );
      await expect(unlock(transactionHashForUnlock, messageHash, token, validator, validator)).to.be.revertedWith("Same signature");
    });
    it("Should not unlock with only 1 signature", async () => {
      const messageHash = ethers.utils.solidityKeccak256(
        ["uint256", "string", "address", "uint256", "address"],
        [(await ethers.provider.getNetwork()).chainId, transactionHashForUnlock, token.address, 100, await user.getAddress()]
      );
      const arrayfiedHash = ethers.utils.arrayify(messageHash);
      const signedMessage = await validator.signMessage(arrayfiedHash);
      const sig = ethers.utils.splitSignature(signedMessage);
      let sigs: Signature[] = [sig];
      await expect(bridge.connect(user).unlockTokens(token.address, 100, transactionHashForUnlock, [sigs[0].v], [sigs[0].r], [sigs[0].s])).to.be.revertedWith("Need at least 2 signatures");
    });
  });
});
