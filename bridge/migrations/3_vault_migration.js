const Vault = artifacts.require("Vault");
const Proxy = artifacts.require("IncognitoProxy");

module.exports = function (deployer) {
  deployer.deploy(Vault, Proxy.address);
};
