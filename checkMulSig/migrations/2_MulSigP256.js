const CurveP256 = artifacts.require('MulSigP256');
const EcRec = artifacts.require('EcRec');

module.exports = function(deployer) {
  deployer.deploy(CurveP256);
  deployer.deploy(EcRec)
};
