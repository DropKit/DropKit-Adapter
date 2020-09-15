// SPDX-License-Identifier: MIT
pragma solidity 0.5.x;

import "./auditLog.sol";

contract Authentication {
    enum OP {SELECT, DELETE, UPDATE, CREATE, INSERT}
    // An owner can take all the CRUD operations.
    // A maintainer is only allowed to do INSERT, UPDATE and SELECT.
    // A viewer can only take SELECT operation.
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}

    // The table for mapping DB Table Name and address which save the the audit log of DB
    // operations.
    mapping(string => address) tableToAddrMap;

    /**
     * @dev Add a new DB table name to the Table Name to Contract Address Map and deploy the Table Contract
     */
    function addNewTableContract(string calldata sk, string calldata tableName)
        external
        returns (bool)
    {
        if (tableToAddrMap[tableName] == address(0)) {
            SK2AddrTable tableContract = new SK2AddrTable(sk);
            tableToAddrMap[tableName] = address(tableContract);

            return true;
        }
        return false;
    }

    /**
     * @dev Retrieve address with given tableName
     * @return table Ethereum address
     */
    function getAddr(string memory tableName) public view returns (address) {
        return tableToAddrMap[tableName];
    }

    /**
     * @dev Append the DB statement to Table Contract.
     * @return whether the operation success
     */
    function appendLog(
        string calldata sk,
        string calldata tableName,
        string calldata statement,
        uint8 operation
    ) external returns (bool) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);

        uint8 role = conv_table.roleMap(sk);
        if (role == uint8(ROLE.NULL)) {
            return false;
        }

        // Check if the user has valid permission
        if (role == uint8(ROLE.MAINTAINER)) {
            if (operation == uint8(OP.CREATE)) {
                return false;
            }
        } else if (role == uint8(ROLE.VIEWER)) {
            if (operation != uint8(OP.SELECT)) {
                return false;
            }
        }

        // Apppend the statement into the contract
        return conv_table.appendLog(sk, statement);
    }

    /**
     * @dev Grant a sk (secret key) as a certain permission
     * @return whether the operation success
     */
    function grantPermission(
        string calldata granterSK,
        string calldata granteeSK,
        string calldata tableName,
        uint8 role
    ) external returns (bool) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);
        uint8 granterRole = conv_table.roleMap(granterSK);

        // Allow the granter and grantee both as OWNER
        if (granterRole < role) {
            return false;
        }
        return conv_table.grantPermission(granteeSK, role);
    }

    /**
     * @dev Verify the permission of a sk (secret key)
     * @return whether the operation success
     */
    function verifyPermission(
        string calldata verifiedSK,
        string calldata tableName
    ) external view returns (uint8) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);

        // TODO Check the permission of the two participants

        return conv_table.verifyPermission(verifiedSK);
    }

    /**
     * @dev Revoke a sk (secret key) from a certain permission
     * @return whether the operation success
     */
    function revokePermission(
        string calldata revokerSK,
        string calldata revokeeSK,
        string calldata tableName
    ) external returns (bool) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);
        uint8 revokerRole = conv_table.roleMap(revokerSK);
        uint8 revokeeRole = conv_table.roleMap(revokeeSK);

        // Allow the revoker and revokee both as OWNER
        if (revokerRole < revokeeRole) {
            return false;
        }
        return conv_table.revokePermission(revokeeSK);
    }
}
