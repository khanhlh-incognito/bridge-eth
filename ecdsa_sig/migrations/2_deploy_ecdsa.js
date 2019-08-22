const ECDSA = artifacts.require("./Ecdsa.sol");

module.exports = function(deployer) {
  deployer.deploy(ECDSA);
};
