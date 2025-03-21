const hre = require("hardhat");
const BigNumber = require('ethers');

async function main() {
  //ABTX
  const AMMBoostTokenX = await hre.ethers.getContractFactory("AMMBoostTokenX");
  const abtx = await AMMBoostTokenX.deploy('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266');
  await abtx.waitForDeployment();
  const abtxADDR = await abtx.getAddress();
  console.log('abtx deployed to', abtxADDR);

  const AMMBoostTokenY = await hre.ethers.getContractFactory("AMMBoostTokenY");
  const abty = await AMMBoostTokenY.deploy('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266');
  await abty.waitForDeployment();
  const abtyADDR = await abty.getAddress();
  console.log('abty deployed to', abtyADDR);

  
  const TokenBank = await hre.ethers.getContractFactory("TokenBank");
  const tb = await TokenBank.deploy(abtxADDR,abtyADDR);
  await tb.waitForDeployment();
  const tbADDR = await tb.getAddress();
  console.log('TB deployed to', tbADDR);
  

}

main();
