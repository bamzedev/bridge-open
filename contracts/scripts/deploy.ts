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
  const bridge = await Bridge.deploy([process.env.VALIDATOR1ADDRESS, process.env.VALIDATOR2ADDRESS]);
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
