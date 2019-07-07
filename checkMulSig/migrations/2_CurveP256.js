const CurveP256 = artifacts.require('CurveP256');

module.exports = function(deployer) {
  deployer.deploy(CurveP256);
};
