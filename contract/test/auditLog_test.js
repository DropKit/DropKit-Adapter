const auditLog = artifacts.require("auditLog");

contract("auditLog", function (accounts) {
  const tableNameTest = "tableName";
  const addrTest = "0xf17f52151EbEF6C7334FAD080c5704D77216b732";

  it("Write with 'addTable'", async () => {
    const auditLogInstance = await auditLog.deployed();
    const auditLogAddTable = await auditLogInstance.addTable.call(
      tableNameTest,
      addrTest
    );

    assert.equal(auditLogAddTable, true, "Add Table with addTable failed");
  });
  it("Retrieve address from mapping", async () => {
    const auditLogInstance = await auditLog.deployed();
    const auditLogAddTable = await auditLogInstance.addTable.sendTransaction(
      tableNameTest,
      addrTest
    );
    const auditLogRetrieve = await auditLogInstance.retrieve.call(
      tableNameTest
    );
    assert.equal(
      auditLogRetrieve,
      addrTest,
      "The retrieved address is not the same as the saved one"
    );
  });
});
