const Proxy = artifacts.require("IncognitoProxy");

module.exports = function(deployer) {
    deployer.deploy(Proxy);
};
