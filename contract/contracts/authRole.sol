// SPDX-License-Identifier: MIT
pragma solidity 0.5.x;

import "./auditLog.sol";

contract Authentication {
    enum OP {SELECT, DELETE, UPDATE, CREATE, INSERT}
    /*
     * OWNER: Owners can take all the CRUD operations.
     * MAINTAINER: Maintainers are only allowed to execute INSERT, UPDATE and SELECT.
     * VIEWER: Viewers can only execute SELECT operation.
     */
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}

    // The table for mapping DB Table Name and address which save the the audit log of DB
    // operations.
    mapping(string => address) tableToAddrMap;

    /**
     * @dev Add a new DB table name to the Table Name to Contract Address Map and deploy the Table Contract
     */
    function addNewTableContract(string calldata tableName)
        external
        returns (bool)
    {
        if (tableToAddrMap[tableName] == address(0)) {
            SK2AddrTable tableContract = new SK2AddrTable(msg.sender);
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
        string calldata tableName,
        string calldata statement,
        uint8 operation
    ) external returns (bool) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);

        uint8 role = conv_table.roleMap(msg.sender);
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
        return conv_table.appendLog(msg.sender, statement);
    }

    /**
     * @dev Grant a address (Ethereum address) with a certain permission
     * @return whether the operation success
     */
    function grantPermission(
        address granteeAddr,
        string calldata tableName,
        uint8 role
    ) external returns (bool) {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);
        uint8 granterRole = conv_table.roleMap(msg.sender);

        // Allow the granter and grantee both as OWNER
        if (granterRole < role) {
            return false;
        }
        return conv_table.grantPermission(granteeAddr, role);
    }

    /**
     * @dev Verify the permission of a address (Ethereum address)
     * @return whether the operation success
     */
    function verifyPermission(address verifiedAddr, string calldata tableName)
        external
        view
        returns (uint8)
    {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);

        // TODO Check the permission of the two participants

        return conv_table.verifyPermission(verifiedAddr);
    }

    /**
     * @dev Revoke a address (Ethereum address) from being a certain permission
     * @return whether the operation success
     */
    function revokePermission(address revokeeAddr, string calldata tableName)
        external
        returns (bool)
    {
        address addr = getAddr(tableName);
        SK2AddrTable conv_table = SK2AddrTable(addr);
        uint8 revokerRole = conv_table.roleMap(msg.sender);
        uint8 revokeeRole = conv_table.roleMap(revokeeAddr);

        // Allow the revoker and revokee both as OWNER
        if (revokerRole < revokeeRole) {
            return false;
        }
        return conv_table.revokePermission(revokeeAddr);
    }
}
