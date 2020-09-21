const authRole = artifacts.require("Authentication");

contract("authRole", function (accounts) {
  const tableNameTest = "tableName";
  const addrNullTest = 0x0000000000000000000000000000000000000000;
  const dbStatementTest = "UPDATE employee0 SET name='David' WHERE id = 1";
  const ownerRole = 3;
  const maintainerRole = 2;
  const viewerRole = 1;
  const nullRole = 0;

  it("Write with 'addNewTableContract'", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.call(
      tableNameTest
    );

    assert.equal(
      authRoleAddNewTableContract,
      true,
      "Add Table with addNewTableContract failed"
    );
  });

  it("Retrieve address from mapping", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      tableNameTest
    );
    const authRoleRetrieve = await authRoleInstance.getAddr.call(tableNameTest);
    assert.notEqual(
      authRoleRetrieve,
      addrNullTest,
      "The retrieved address is not the same as the saved one"
    );
  });

  it("Append audit log to Table Contract", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      tableNameTest
    );

    const authRoleAppend = await authRoleInstance.appendLog.call(
      tableNameTest,
      dbStatementTest,
      ownerRole
    );
    assert.equal(
      authRoleAppend,
      true,
      "Failed to append DB statement to Table Contract"
    );
  });

  it("Grant the given address (`accounts[1]`) as MAINTAINER", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.call(
      accounts[1],
      tableNameTest,
      maintainerRole
    );
    assert.equal(
      authGrantPermission,
      true,
      "Assign the given address as a MAINTAINER"
    );
  });

  it("Check whether the given address is a MAINTAINER or not", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.sendTransaction(
      accounts[1],
      tableNameTest,
      maintainerRole
    );
    const authVerifyPermission = await authRoleInstance.verifyPermission.call(
      accounts[1],
      tableNameTest
    );
    assert.equal(
      authVerifyPermission,
      maintainerRole,
      "The given address doesn't refer to a MAINTAINER role"
    );
  });

  it("Revoke the given address from being a MAINTAINER", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.sendTransaction(
      accounts[1],
      tableNameTest,
      maintainerRole
    );
    var authVerifyPermission = await authRoleInstance.verifyPermission.call(
      accounts[1],
      tableNameTest
    );
    assert.equal(
      authVerifyPermission,
      maintainerRole,
      "The given address doesn't refer to a MAINTAINER role"
    );

    // Check the response
    var authRevokePermission = await authRoleInstance.revokePermission.call(
      accounts[1],
      tableNameTest
    );
    assert.equal(
      authRevokePermission,
      true,
      "Failed to revoke the role of the given address from being a MAINTAINER"
    );

    // Revoke again for modifying the data on Blockchain
    authRevokePermission = await authRoleInstance.revokePermission.sendTransaction(
      accounts[1],
      tableNameTest
    );
    authVerifyPermission = await authRoleInstance.verifyPermission.call(
      accounts[1],
      tableNameTest
    );

    assert.equal(
      authVerifyPermission,
      nullRole,
      "Failed to revoke the role of the given address from being a MAINTAINER"
    );
  });
});
