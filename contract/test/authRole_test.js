const authRole = artifacts.require("Authentication");

contract("authRole", function (accounts) {
  const skOwnerTest =
    "e0302f799ee2e6bcff366cf395dda8225e0a3ae9250740aeabf8e174a8d55c03";
  const skMaintainerTest =
    "aaf4f5s4df5sd47h2ns4i54yk24a65sg7jt7iagj8lo4lj55wfwefjjjukaewqq4";
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
      skOwnerTest,
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
      skOwnerTest,
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
      skOwnerTest,
      tableNameTest
    );

    const authRoleAppend = await authRoleInstance.appendLog.call(
      skOwnerTest,
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

  it("Make a user as MAINTAINER", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      skOwnerTest,
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.call(
      skOwnerTest,
      skMaintainerTest,
      tableNameTest,
      maintainerRole
    );
    assert.equal(authGrantPermission, true, "Assign skTest as MAINTAINER");
  });

  it("Check whether skTest is a MAINTAINER", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      skOwnerTest,
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.sendTransaction(
      skOwnerTest,
      skMaintainerTest,
      tableNameTest,
      maintainerRole
    );
    const authVerifyPermission = await authRoleInstance.verifyPermission.call(
      skMaintainerTest,
      tableNameTest
    );
    assert.equal(
      authVerifyPermission,
      maintainerRole,
      "skTest doesn't refer to a MAINTAINER role"
    );
  });

  it("Revoke skTest from being a MAINTAINER", async () => {
    const authRoleInstance = await authRole.deployed();
    const authRoleAddNewTableContract = await authRoleInstance.addNewTableContract.sendTransaction(
      skOwnerTest,
      tableNameTest
    );

    const authGrantPermission = await authRoleInstance.grantPermission.sendTransaction(
      skOwnerTest,
      skMaintainerTest,
      tableNameTest,
      maintainerRole
    );
    var authVerifyPermission = await authRoleInstance.verifyPermission.call(
      skMaintainerTest,
      tableNameTest
    );
    assert.equal(
      authVerifyPermission,
      maintainerRole,
      "skTest doesn't refer to a MAINTAINER role"
    );

    // Check the response
    var authRevokePermission = await authRoleInstance.revokePermission.call(
      skOwnerTest,
      skMaintainerTest,
      tableNameTest
    );
    assert.equal(
      authRevokePermission,
      true,
      "Failed to revoke the role of skTest from being a MAINTAINER"
    );

    // Revoke again for modifying the data on Blockchain
    authRevokePermission = await authRoleInstance.revokePermission.sendTransaction(
      skOwnerTest,
      skMaintainerTest,
      tableNameTest
    );
    authVerifyPermission = await authRoleInstance.verifyPermission.call(
      skMaintainerTest,
      tableNameTest
    );

    assert.equal(
      authVerifyPermission,
      nullRole,
      "Failed to revoke the role of skTest from being a MAINTAINER"
    );
  });
});
