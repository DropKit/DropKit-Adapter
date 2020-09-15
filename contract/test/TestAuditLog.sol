pragma solidity 0.5.x;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/authRole.sol";

contract TestAuditLog {
    string skOwnerTest = "e0302f799ee2e6bcff366cf395dda8225e0a3ae9250740aeabf8e174a8d55c03";
    string skMaintainerTest = "aaf4f5s4df5sd47h2ns4i54yk24a65sg7jt7iagj8lo4lj55wfwefjjjukaewqq4";
    string tableNameTest = "TableName";
    address addrTest = 0xc0ffee254729296a45a3885639AC7E10F9d54979;
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}

    function testAddNewTableContract() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(skOwnerTest, tableNameTest),
            true,
            "Write a Table with a address to the map"
        );
    }

    function testGetAddr() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(skOwnerTest, tableNameTest),
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
            auth.addNewTableContract(skOwnerTest, tableNameTest),
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
                skOwnerTest,
                skMaintainerTest,
                tableNameTest,
                uint8(ROLE.MAINTAINER)
            ),
            true,
            "Grant user as MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(skMaintainerTest, tableNameTest)),
            uint256(ROLE.MAINTAINER),
            "Verify user as MAINTAINER"
        );
    }

    function testRevokePermission() public {
        Authentication auth = new Authentication();
        Assert.equal(
            auth.addNewTableContract(skOwnerTest, tableNameTest),
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
                skOwnerTest,
                skMaintainerTest,
                tableNameTest,
                uint8(ROLE.MAINTAINER)
            ),
            true,
            "Grant user as MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(skMaintainerTest, tableNameTest)),
            uint256(ROLE.MAINTAINER),
            "Verify user as MAINTAINER"
        );
        Assert.equal(
            auth.revokePermission(skOwnerTest, skMaintainerTest, tableNameTest),
            true,
            "Revoke user from being MAINTAINER"
        );
        Assert.equal(
            uint256(auth.verifyPermission(skMaintainerTest, tableNameTest)),
            uint256(ROLE.NULL),
            "Verify the secret key holds by user as NULL"
        );
    }
}
