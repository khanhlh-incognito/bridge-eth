const Vault = artifacts.require("Vault");

module.exports = function (deployer) {
  const incognitoProxyAddress = "0xcb859a5fC20EEeCc4Cec191d8CCe5e31a2CC1dAF";
  deployer.deploy(Vault, incognitoProxyAddress);
};
