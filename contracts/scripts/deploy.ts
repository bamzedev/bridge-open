// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { run, ethers } from "hardhat";

async function deployBridgeContract() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  await run("compile");
  console.log("Deploying...");
  const Bridge = await ethers.getContractFactory("Bridge");
  const bridge = await Bridge.deploy(['0xEaDc3D0D313f90D2E9A6B8982b71eB05e5c09D8c','0x80211bE2056A5C6C00aecDb5dfCd43F243E4d2E0']);
  //const WrappedTether = await ethers.getContractFactory("WrappedToken");
  //const wtoken = await WrappedTether.deploy("Test Token", "TESTT");
  await bridge.deployed();
  //await wtoken.deployed()
  console.log("Bridge contract deployed to:", bridge.address);
  //console.log("Wrapped Token contract deployed to:", wtoken.address);
}
module.exports = deployBridgeContract;
// // We recommend this pattern to be able to use async/await everywhere
// // and properly handle errors.
// deployBridgeContract().catch((error) => {
//   console.error(error);
//   process.exitCode = 1;
// });
