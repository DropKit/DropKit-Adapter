const auditLog = artifacts.require("auditLog");

module.exports = function (deployer) {
  deployer.deploy(auditLog);
};
