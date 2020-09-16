pragma solidity 0.5.x;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/auditLog.sol";

contract TestAuditLog {
    string tableNameTest = "TableName";
    address addrTest = 0xc0ffee254729296a45a3885639AC7E10F9d54979;

    function testAddTable() public {
        auditLog audit = new auditLog();
        Assert.equal(
            audit.addTable(tableNameTest, addrTest),
            true,
            "Write a Table with a address to the map"
        );
    }

    function testRetrieve() public {
        auditLog audit = new auditLog();
        Assert.equal(
            audit.addTable(tableNameTest, addrTest),
            true,
            "Write a Table with a address to the map"
        );
        Assert.equal(
            audit.retrieve(tableNameTest),
            addrTest,
            "Retrieve the Table Address with Table name"
        );
    }
}
