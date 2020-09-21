// SPDX-License-Identifier: MIT
pragma solidity 0.5.x;

contract SK2AddrTable {
    address internal deployerAddr;
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}
    mapping(address => uint8) public roleMap;
    event statementLog(address addr, string statement);

    modifier ownership {
        assert(deployerAddr == msg.sender);
        _;
    }

    // Initailize default state
    constructor(address dbTableOwnerAddr) public {
        // Currently only the deployer can interact with Table Contract
        deployerAddr = msg.sender;
        roleMap[dbTableOwnerAddr] = uint8(ROLE.OWNER);
    }

    function appendLog(address addr, string calldata statement)
        external
        ownership
        returns (bool)
    {
        // Apppend the statement into the contract with event log (addr, and statement)
        emit statementLog(addr, statement);
        return true;
    }

    function grantPermission(address addr, uint8 role)
        external
        ownership
        returns (bool)
    {
        roleMap[addr] = role;
        return true;
    }

    function verifyPermission(address addr)
        external
        view
        ownership
        returns (uint8)
    {
        return roleMap[addr];
    }

    function revokePermission(address addr) external ownership returns (bool) {
        roleMap[addr] = uint8(ROLE.NULL);
        return true;
    }
}
