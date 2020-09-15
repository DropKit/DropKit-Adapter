const authentication = artifacts.require("Authentication");
const sk2AddrTable = artifacts.require("SK2AddrTable");

const addrTest = "0xf17f52151EbEF6C7334FAD080c5704D77216b732";

module.exports = function (deployer) {
  deployer.deploy(sk2AddrTable, addrTest);
  deployer.link(sk2AddrTable, authentication);
  deployer.deploy(authentication);
};
