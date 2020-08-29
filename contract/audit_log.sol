// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.7.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */
contract audit_log {
    // The table for mapping DB table name and address which save the the audit log of DB
    // operations.
    mapping(string => address) auditLogMap;

    /**
     * @dev Add a new DB table name and corresponding ethereum public key
     */
    // FIXME We need to implement a mechanism to avoid tableName conflicting among different organizations
    function addTable(string calldata tableName, address dbAccount)
        external
        returns (bool)
    {
        if (auditLogMap[tableName] == address(0)) {
            auditLogMap[tableName] = dbAccount;
            return true;
        }
        return false;
    }

    /**
     * @dev Retrieve address with given tableName
     * @return table Ethereum address
     */
    function retrieve(string calldata tableName)
        external
        view
        returns (address)
    {
        return auditLogMap[tableName];
    }
}
