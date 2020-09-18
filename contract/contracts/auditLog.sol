// SPDX-License-Identifier: MIT
pragma solidity 0.5.x;

contract SK2AddrTable {
    address internal deployerAddr;
    string ownerSK;
    enum ROLE {NULL, VIEWER, MAINTAINER, OWNER}
    mapping(string => uint8) public roleMap;
    event statementLog(string sk, string statement);

    modifier ownership {
        assert(deployerAddr == msg.sender);
        _;
    }

    // Initailize default state
    constructor(string memory initOwnerSK) public {
        // Currently only the deployer can interact with Table Contract
        deployerAddr = msg.sender;
        ownerSK = initOwnerSK;
        roleMap[initOwnerSK] = uint8(ROLE.OWNER);
    }

    function appendLog(string calldata sk, string calldata statement)
        external
        ownership
        returns (bool)
    {
        // Apppend the statement into the contract with event log (sk, and statement)
        emit statementLog(sk, statement);
        return true;
    }

    function grantPermission(string calldata sk, uint8 role)
        external
        ownership
        returns (bool)
    {
        roleMap[sk] = role;
        return true;
    }

    function verifyPermission(string calldata sk)
        external
        view
        ownership
        returns (uint8)
    {
        return roleMap[sk];
    }

    function revokePermission(string calldata sk)
        external
        ownership
        returns (bool)
    {
        roleMap[sk] = uint8(ROLE.NULL);
        return true;
    }
}
