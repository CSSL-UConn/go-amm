const TokenBank = artifacts.require('TokenBank')
const TokenA = artifacts.require('AMMBoostTokenX')
const TokenB = artifacts.require('AMMBoostTokenY')
module.exports = function(_deployer, network, accounts) {
  admin = accounts[0]
  // Use deployer to state migration tasks.
  _deployer.deploy(TokenA, admin).then(function(){
    _deployer.deploy(TokenB, admin).then(function(){
      _deployer.deploy(TokenBank, TokenA.address, TokenB.address);
    }
    );
  });
};
