import { ethers, Signature } from "ethers";
import { signERC2612Permit } from "eth-permit";

export const lockTokens = async (provider, bridgeContract, tokenContract, amount) => {
  console.log(await bridgeContract.fee());
  const senderAddress = await provider.getSigner().getAddress();
  const result = await signERC2612Permit(provider.getSigner(), tokenContract.address, senderAddress, bridgeContract.address, amount);
  let receipt = await bridgeContract.connect(provider.getSigner()).lock(await provider.getSigner().getAddress(), tokenContract.address, amount, result.deadline, result.v, result.r, result.s, {value: await bridgeContract.fee()});
  await receipt.wait(1);
  return receipt;
};
export const burnTokens = async (provider, bridgeContract, tokenContract, amount) => {
  console.log(bridgeContract.address);
  const burnMessageHash = ethers.utils.solidityKeccak256(
    ["address", "address", "address", "uint256"],
    [await provider.getSigner().getAddress(), bridgeContract.address, tokenContract.address, amount]
  );
  const burnArrayfiedHash = ethers.utils.arrayify(burnMessageHash);
  const burnSignedMessage = await provider.getSigner().signMessage(burnArrayfiedHash);
  const burnSig = ethers.utils.splitSignature(burnSignedMessage);

  let receipt = await bridgeContract.connect(provider.getSigner()).burn(await provider.getSigner().getAddress(), tokenContract.address, amount, burnSig.v, burnSig.r, burnSig.s, {value: await bridgeContract.fee()});
  await receipt.wait(1);
  return receipt;
};

export const claimTokens = async (provider, bridgeContract, depositTransaction) => {
  const sig = ethers.utils.splitSignature(depositTransaction.signatures[0]);
  const sig2 = ethers.utils.splitSignature(depositTransaction.signatures[1]);
  let sigs: Signature[] = [sig, sig2];
  console.log(depositTransaction.tokenAddress, depositTransaction.amount, depositTransaction.depositTransactionHash);
  let receipt
  if(depositTransaction.isBurn){
    receipt = await bridgeContract
      .connect(provider.getSigner())
      .unlockTokens(depositTransaction.tokenAddress, depositTransaction.amount, depositTransaction.depositTransactionHash, [sigs[0].v, sigs[1].v], [sigs[0].r, sigs[1].r], [sigs[0].s, sigs[1].s]);
    await receipt.wait(1);
  }else{
    receipt = await bridgeContract
      .connect(provider.getSigner())
      .mintTokens(depositTransaction.tokenAddress, "Wrapped "+depositTransaction.name, "W"+depositTransaction.symbol, depositTransaction.amount, depositTransaction.depositTransactionHash, [sigs[0].v, sigs[1].v], [sigs[0].r, sigs[1].r], [sigs[0].s, sigs[1].s]);
    await receipt.wait(1);
  }
  depositTransaction.claimTransactionHash = receipt.hash;
  return depositTransaction;
};
