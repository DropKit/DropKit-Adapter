pragma solidity 0.5.x;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/authRole.sol";

contract TestAuditLog {
    string tableNameTest = "TableName";
    address addrTest = 0xc0ffee254729296a45a3885639AC7E10F9d54979;
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}

    function testAddNewTableContract() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(tableNameTest),
            true,
            "Write a Table with a address to the map"
        );
    }

    function testGetAddr() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(tableNameTest),
            true,
            "Write a Table with a address to the map"
        );
        Assert.notEqual(
            auth.getAddr(tableNameTest),
            address(0),
            "Retrieve the Table Address with Table name"
        );
    }

    function testVerifyPermission() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(tableNameTest),
            true,
            "Write a Table with a address to the map"
        );
        Assert.notEqual(
            auth.getAddr(tableNameTest),
            address(0),
            "Retrieve the Table Address with Table name"
        );
        Assert.equal(
            auth.grantPermission(
                addrTest,
                tableNameTest,
                uint8(ROLE.MAINTAINER)
            ),
            true,
            "Grant user as MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(addrTest, tableNameTest)),
            uint256(ROLE.MAINTAINER),
            "Verify user as MAINTAINER"
        );
    }

    function testRevokePermission() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(tableNameTest),
            true,
            "Write a Table with a address to the map"
        );
        Assert.notEqual(
            auth.getAddr(tableNameTest),
            address(0),
            "Retrieve the Table Address with Table name"
        );
        Assert.equal(
            auth.grantPermission(
                addrTest,
                tableNameTest,
                uint8(ROLE.MAINTAINER)
            ),
            true,
            "Grant user as MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(addrTest, tableNameTest)),
            uint256(ROLE.MAINTAINER),
            "Verify user as MAINTAINER"
        );
        Assert.equal(
            auth.revokePermission(addrTest, tableNameTest),
            true,
            "Revoke user from being MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(addrTest, tableNameTest)),
            uint256(ROLE.NULL),
            "Verify the secret key holds by user as NULL"
        );
    }
}
