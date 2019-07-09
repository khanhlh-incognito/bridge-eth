const CurveP256 = artifacts.require('MulSigP256');

module.exports = function(deployer) {
  deployer.deploy(CurveP256);
};
