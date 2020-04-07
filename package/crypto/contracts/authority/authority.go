// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package authority

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AuthorityABI is the input ABI used to generate the binding from.
const AuthorityABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"has\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AuthorityBin is the compiled bytecode used for deploying new contracts.
var AuthorityBin = "0x608060405234801561001057600080fd5b5061074a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063308714911461004657806344590a7e146101b0578063ebdf86ca14610302575b600080fd5b6101966004803603604081101561005c57600080fd5b810190808035906020019064010000000081111561007957600080fd5b82018360208201111561008b57600080fd5b803590602001918460018302840111640100000000831117156100ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561011057600080fd5b82018360208201111561012257600080fd5b8035906020019184600183028401116401000000008311171561014457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610454565b604051808215151515815260200191505060405180910390f35b610300600480360360408110156101c657600080fd5b81019080803590602001906401000000008111156101e357600080fd5b8201836020820111156101f557600080fd5b8035906020019184600183028401116401000000008311171561021757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561027a57600080fd5b82018360208201111561028c57600080fd5b803590602001918460018302840111640100000000831117156102ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061053b565b005b6104526004803603604081101561031857600080fd5b810190808035906020019064010000000081111561033557600080fd5b82018360208201111561034757600080fd5b8035906020019184600183028401116401000000008311171561036957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156103cc57600080fd5b8201836020820111156103de57600080fd5b8035906020019184600183028401116401000000008311171561040057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610627565b005b600080836040518082805190602001908083835b6020831061048b5780518252602082019150602081019050602083039250610468565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020826040518082805190602001908083835b602083106104f257805182526020820191506020810190506020830392506104cf565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff16905092915050565b600080836040518082805190602001908083835b60208310610572578051825260208201915060208101905060208303925061054f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020826040518082805190602001908083835b602083106105d957805182526020820191506020810190506020830392506105b6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff0219169083151502179055505050565b60016000836040518082805190602001908083835b6020831061065f578051825260208201915060208101905060208303925061063c565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020826040518082805190602001908083835b602083106106c657805182526020820191506020810190506020830392506106a3565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550505056fea2646970667358221220caf4197f13b01873c25f49ef0667e79717256f5d1d907ff4d1c6e6e6251700af64736f6c63430006020033"

// DeployAuthority deploys a new Ethereum contract, binding an instance of Authority to it.
func DeployAuthority(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authority, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthorityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}, AuthorityFilterer: AuthorityFilterer{contract: contract}}, nil
}

// Authority is an auto generated Go binding around an Ethereum contract.
type Authority struct {
	AuthorityCaller     // Read-only binding to the contract
	AuthorityTransactor // Write-only binding to the contract
	AuthorityFilterer   // Log filterer for contract events
}

// AuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthoritySession struct {
	Contract     *Authority        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityCallerSession struct {
	Contract *AuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityTransactorSession struct {
	Contract     *AuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityRaw struct {
	Contract *Authority // Generic contract binding to access the raw methods on
}

// AuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityCallerRaw struct {
	Contract *AuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityTransactorRaw struct {
	Contract *AuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthority creates a new instance of Authority, bound to a specific deployed contract.
func NewAuthority(address common.Address, backend bind.ContractBackend) (*Authority, error) {
	contract, err := bindAuthority(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}, AuthorityFilterer: AuthorityFilterer{contract: contract}}, nil
}

// NewAuthorityCaller creates a new read-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityCaller(address common.Address, caller bind.ContractCaller) (*AuthorityCaller, error) {
	contract, err := bindAuthority(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityCaller{contract: contract}, nil
}

// NewAuthorityTransactor creates a new write-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityTransactor, error) {
	contract, err := bindAuthority(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityTransactor{contract: contract}, nil
}

// NewAuthorityFilterer creates a new log filterer instance of Authority, bound to a specific deployed contract.
func NewAuthorityFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorityFilterer, error) {
	contract, err := bindAuthority(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorityFilterer{contract: contract}, nil
}

// bindAuthority binds a generic wrapper to an already deployed contract.
func bindAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.AuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transact(opts, method, params...)
}

// Has is a free data retrieval call binding the contract method 0x30871491.
//
// Solidity: function has(string tableName, string user) constant returns(bool)
func (_Authority *AuthorityCaller) Has(opts *bind.CallOpts, tableName string, user string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "has", tableName, user)
	return *ret0, err
}

// Has is a free data retrieval call binding the contract method 0x30871491.
//
// Solidity: function has(string tableName, string user) constant returns(bool)
func (_Authority *AuthoritySession) Has(tableName string, user string) (bool, error) {
	return _Authority.Contract.Has(&_Authority.CallOpts, tableName, user)
}

// Has is a free data retrieval call binding the contract method 0x30871491.
//
// Solidity: function has(string tableName, string user) constant returns(bool)
func (_Authority *AuthorityCallerSession) Has(tableName string, user string) (bool, error) {
	return _Authority.Contract.Has(&_Authority.CallOpts, tableName, user)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string user) returns()
func (_Authority *AuthorityTransactor) Add(opts *bind.TransactOpts, tableName string, user string) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "add", tableName, user)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string user) returns()
func (_Authority *AuthoritySession) Add(tableName string, user string) (*types.Transaction, error) {
	return _Authority.Contract.Add(&_Authority.TransactOpts, tableName, user)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string user) returns()
func (_Authority *AuthorityTransactorSession) Add(tableName string, user string) (*types.Transaction, error) {
	return _Authority.Contract.Add(&_Authority.TransactOpts, tableName, user)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string tableName, string user) returns()
func (_Authority *AuthorityTransactor) Remove(opts *bind.TransactOpts, tableName string, user string) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "remove", tableName, user)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string tableName, string user) returns()
func (_Authority *AuthoritySession) Remove(tableName string, user string) (*types.Transaction, error) {
	return _Authority.Contract.Remove(&_Authority.TransactOpts, tableName, user)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string tableName, string user) returns()
func (_Authority *AuthorityTransactorSession) Remove(tableName string, user string) (*types.Transaction, error) {
	return _Authority.Contract.Remove(&_Authority.TransactOpts, tableName, user)
}
