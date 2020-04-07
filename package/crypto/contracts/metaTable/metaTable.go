// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metaTable

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

// MetaTableABI is the input ABI used to generate the binding from.
const MetaTableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"newTable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_tableAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tableIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tableAddress\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTable\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tableRegistryCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MetaTableBin is the compiled bytecode used for deploying new contracts.
var MetaTableBin = "0x608060405234801561001057600080fd5b50610e03806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630d0d141d1461006757806359ca98db14610086578063693ec85e146100b6578063a4a2b687146100e6578063a90797dc14610104578063ebdf86ca14610134575b600080fd5b61006f610164565b60405161007d929190610b4f565b60405180910390f35b6100a0600480360361009b91908101906107f9565b6103a8565b6040516100ad9190610ba1565b60405180910390f35b6100d060048036036100cb919081019061083a565b61046e565b6040516100dd9190610bc3565b60405180910390f35b6100ee61052e565b6040516100fb9190610c1c565b60405180910390f35b61011e600480360361011991908101906108e7565b610534565b60405161012b9190610ba1565b60405180910390f35b61014e6004803603610149919081019061087b565b6105e4565b60405161015b9190610b86565b60405180910390f35b606080606060025460405190808252806020026020018201604052801561019f57816020015b606081526020019060019003908161018a5790505b50905060606002546040519080825280602002602001820160405280156101da57816020015b60608152602001906001900390816101c55790505b50905060008090505b60025481101561039b576000808281526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102935780601f1061026857610100808354040283529160200191610293565b820191906000526020600020905b81548152906001019060200180831161027657829003601f168201915b50505050508282815181106102a457fe5b602002602001018190525060016000808381526020019081526020016000206040516102d09190610b38565b90815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103725780601f1061034757610100808354040283529160200191610372565b820191906000526020600020905b81548152906001019060200180831161035557829003601f168201915b505050505083828151811061038357fe5b602002602001018190525080806001019150506101e3565b5080829350935050509091565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104665780601f1061043b57610100808354040283529160200191610466565b820191906000526020600020905b81548152906001019060200180831161044957829003601f168201915b505050505081565b60606001826040516104809190610b21565b90815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105225780601f106104f757610100808354040283529160200191610522565b820191906000526020600020905b81548152906001019060200180831161050557829003601f168201915b50505050509050919050565b60025481565b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105dc5780601f106105b1576101008083540402835291602001916105dc565b820191906000526020600020905b8154815290600101906020018083116105bf57829003601f168201915b505050505081565b6000826000806002548152602001908152602001600020908051906020019061060e929190610697565b50816001846040516106209190610b21565b90815260200160405180910390209080519060200190610641929190610697565b506002600081548092919060010191905055507f18e448d301b2feb12847d36273372ac332ce9fc0de8d1137c1d9a4e96b569bd98383604051610685929190610be5565b60405180910390a16001905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106d857805160ff1916838001178555610706565b82800160010185558215610706579182015b828111156107055782518255916020019190600101906106ea565b5b5090506107139190610717565b5090565b61073991905b8082111561073557600081600090555060010161071d565b5090565b90565b600082601f83011261074d57600080fd5b813561076061075b82610c64565b610c37565b9150808252602083016020830185838301111561077c57600080fd5b610787838284610d63565b50505092915050565b600082601f8301126107a157600080fd5b81356107b46107af82610c90565b610c37565b915080825260208301602083018583830111156107d057600080fd5b6107db838284610d63565b50505092915050565b6000813590506107f381610db6565b92915050565b60006020828403121561080b57600080fd5b600082013567ffffffffffffffff81111561082557600080fd5b6108318482850161073c565b91505092915050565b60006020828403121561084c57600080fd5b600082013567ffffffffffffffff81111561086657600080fd5b61087284828501610790565b91505092915050565b6000806040838503121561088e57600080fd5b600083013567ffffffffffffffff8111156108a857600080fd5b6108b485828601610790565b925050602083013567ffffffffffffffff8111156108d157600080fd5b6108dd85828601610790565b9150509250929050565b6000602082840312156108f957600080fd5b6000610907848285016107e4565b91505092915050565b600061091c8383610a12565b905092915050565b600061092f82610ce1565b6109398185610d0f565b93508360208202850161094b85610cbc565b8060005b8581101561098757848403895281516109688582610910565b945061097383610d02565b925060208a0199505060018101905061094f565b50829750879550505050505092915050565b6109a281610d4d565b82525050565b60006109b382610cf7565b6109bd8185610d31565b93506109cd818560208601610d72565b6109d681610da5565b840191505092915050565b60006109ec82610cf7565b6109f68185610d42565b9350610a06818560208601610d72565b80840191505092915050565b6000610a1d82610cec565b610a278185610d20565b9350610a37818560208601610d72565b610a4081610da5565b840191505092915050565b6000610a5682610cec565b610a608185610d31565b9350610a70818560208601610d72565b610a7981610da5565b840191505092915050565b600081546001811660008114610aa15760018114610ac657610b0a565b607f6002830416610ab28187610d42565b955060ff1983168652808601935050610b0a565b60028204610ad48187610d42565b9550610adf85610ccc565b60005b82811015610b0157815481890152600182019150602081019050610ae2565b82880195505050505b505092915050565b610b1b81610d59565b82525050565b6000610b2d82846109e1565b915081905092915050565b6000610b448284610a84565b915081905092915050565b60006040820190508181036000830152610b698185610924565b90508181036020830152610b7d8184610924565b90509392505050565b6000602082019050610b9b6000830184610999565b92915050565b60006020820190508181036000830152610bbb8184610a4b565b905092915050565b60006020820190508181036000830152610bdd81846109a8565b905092915050565b60006040820190508181036000830152610bff81856109a8565b90508181036020830152610c1381846109a8565b90509392505050565b6000602082019050610c316000830184610b12565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610c5a57600080fd5b8060405250919050565b600067ffffffffffffffff821115610c7b57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115610ca757600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610d90578082015181840152602081019050610d75565b83811115610d9f576000848401525b50505050565b6000601f19601f8301169050919050565b610dbf81610d59565b8114610dca57600080fd5b5056fea264697066735822122022136240f78abab0c61d8de83a9067a72f655b40e9dbe8bcc9fcc690e251e0f464736f6c63430006020033"

// DeployMetaTable deploys a new Ethereum contract, binding an instance of MetaTable to it.
func DeployMetaTable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MetaTable, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaTableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MetaTableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaTable{MetaTableCaller: MetaTableCaller{contract: contract}, MetaTableTransactor: MetaTableTransactor{contract: contract}, MetaTableFilterer: MetaTableFilterer{contract: contract}}, nil
}

// MetaTable is an auto generated Go binding around an Ethereum contract.
type MetaTable struct {
	MetaTableCaller     // Read-only binding to the contract
	MetaTableTransactor // Write-only binding to the contract
	MetaTableFilterer   // Log filterer for contract events
}

// MetaTableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaTableSession struct {
	Contract     *MetaTable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaTableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaTableCallerSession struct {
	Contract *MetaTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MetaTableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaTableTransactorSession struct {
	Contract     *MetaTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MetaTableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaTableRaw struct {
	Contract *MetaTable // Generic contract binding to access the raw methods on
}

// MetaTableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaTableCallerRaw struct {
	Contract *MetaTableCaller // Generic read-only contract binding to access the raw methods on
}

// MetaTableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaTableTransactorRaw struct {
	Contract *MetaTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaTable creates a new instance of MetaTable, bound to a specific deployed contract.
func NewMetaTable(address common.Address, backend bind.ContractBackend) (*MetaTable, error) {
	contract, err := bindMetaTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaTable{MetaTableCaller: MetaTableCaller{contract: contract}, MetaTableTransactor: MetaTableTransactor{contract: contract}, MetaTableFilterer: MetaTableFilterer{contract: contract}}, nil
}

// NewMetaTableCaller creates a new read-only instance of MetaTable, bound to a specific deployed contract.
func NewMetaTableCaller(address common.Address, caller bind.ContractCaller) (*MetaTableCaller, error) {
	contract, err := bindMetaTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTableCaller{contract: contract}, nil
}

// NewMetaTableTransactor creates a new write-only instance of MetaTable, bound to a specific deployed contract.
func NewMetaTableTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaTableTransactor, error) {
	contract, err := bindMetaTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTableTransactor{contract: contract}, nil
}

// NewMetaTableFilterer creates a new log filterer instance of MetaTable, bound to a specific deployed contract.
func NewMetaTableFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaTableFilterer, error) {
	contract, err := bindMetaTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaTableFilterer{contract: contract}, nil
}

// bindMetaTable binds a generic wrapper to an already deployed contract.
func bindMetaTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaTableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTable *MetaTableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaTable.Contract.MetaTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTable *MetaTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTable.Contract.MetaTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTable *MetaTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTable.Contract.MetaTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTable *MetaTableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTable *MetaTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTable *MetaTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTable.Contract.contract.Transact(opts, method, params...)
}

// TableAddresses is a free data retrieval call binding the contract method 0x59ca98db.
//
// Solidity: function _tableAddresses(string ) constant returns(string)
func (_MetaTable *MetaTableCaller) TableAddresses(opts *bind.CallOpts, arg0 string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MetaTable.contract.Call(opts, out, "_tableAddresses", arg0)
	return *ret0, err
}

// TableAddresses is a free data retrieval call binding the contract method 0x59ca98db.
//
// Solidity: function _tableAddresses(string ) constant returns(string)
func (_MetaTable *MetaTableSession) TableAddresses(arg0 string) (string, error) {
	return _MetaTable.Contract.TableAddresses(&_MetaTable.CallOpts, arg0)
}

// TableAddresses is a free data retrieval call binding the contract method 0x59ca98db.
//
// Solidity: function _tableAddresses(string ) constant returns(string)
func (_MetaTable *MetaTableCallerSession) TableAddresses(arg0 string) (string, error) {
	return _MetaTable.Contract.TableAddresses(&_MetaTable.CallOpts, arg0)
}

// TableIndex is a free data retrieval call binding the contract method 0xa90797dc.
//
// Solidity: function _tableIndex(uint256 ) constant returns(string)
func (_MetaTable *MetaTableCaller) TableIndex(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MetaTable.contract.Call(opts, out, "_tableIndex", arg0)
	return *ret0, err
}

// TableIndex is a free data retrieval call binding the contract method 0xa90797dc.
//
// Solidity: function _tableIndex(uint256 ) constant returns(string)
func (_MetaTable *MetaTableSession) TableIndex(arg0 *big.Int) (string, error) {
	return _MetaTable.Contract.TableIndex(&_MetaTable.CallOpts, arg0)
}

// TableIndex is a free data retrieval call binding the contract method 0xa90797dc.
//
// Solidity: function _tableIndex(uint256 ) constant returns(string)
func (_MetaTable *MetaTableCallerSession) TableIndex(arg0 *big.Int) (string, error) {
	return _MetaTable.Contract.TableIndex(&_MetaTable.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string tableName) constant returns(string)
func (_MetaTable *MetaTableCaller) Get(opts *bind.CallOpts, tableName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MetaTable.contract.Call(opts, out, "get", tableName)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string tableName) constant returns(string)
func (_MetaTable *MetaTableSession) Get(tableName string) (string, error) {
	return _MetaTable.Contract.Get(&_MetaTable.CallOpts, tableName)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string tableName) constant returns(string)
func (_MetaTable *MetaTableCallerSession) Get(tableName string) (string, error) {
	return _MetaTable.Contract.Get(&_MetaTable.CallOpts, tableName)
}

// GetAllTable is a free data retrieval call binding the contract method 0x0d0d141d.
//
// Solidity: function getAllTable() constant returns(string[], string[])
func (_MetaTable *MetaTableCaller) GetAllTable(opts *bind.CallOpts) ([]string, []string, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MetaTable.contract.Call(opts, out, "getAllTable")
	return *ret0, *ret1, err
}

// GetAllTable is a free data retrieval call binding the contract method 0x0d0d141d.
//
// Solidity: function getAllTable() constant returns(string[], string[])
func (_MetaTable *MetaTableSession) GetAllTable() ([]string, []string, error) {
	return _MetaTable.Contract.GetAllTable(&_MetaTable.CallOpts)
}

// GetAllTable is a free data retrieval call binding the contract method 0x0d0d141d.
//
// Solidity: function getAllTable() constant returns(string[], string[])
func (_MetaTable *MetaTableCallerSession) GetAllTable() ([]string, []string, error) {
	return _MetaTable.Contract.GetAllTable(&_MetaTable.CallOpts)
}

// TableRegistryCount is a free data retrieval call binding the contract method 0xa4a2b687.
//
// Solidity: function tableRegistryCount() constant returns(uint256)
func (_MetaTable *MetaTableCaller) TableRegistryCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaTable.contract.Call(opts, out, "tableRegistryCount")
	return *ret0, err
}

// TableRegistryCount is a free data retrieval call binding the contract method 0xa4a2b687.
//
// Solidity: function tableRegistryCount() constant returns(uint256)
func (_MetaTable *MetaTableSession) TableRegistryCount() (*big.Int, error) {
	return _MetaTable.Contract.TableRegistryCount(&_MetaTable.CallOpts)
}

// TableRegistryCount is a free data retrieval call binding the contract method 0xa4a2b687.
//
// Solidity: function tableRegistryCount() constant returns(uint256)
func (_MetaTable *MetaTableCallerSession) TableRegistryCount() (*big.Int, error) {
	return _MetaTable.Contract.TableRegistryCount(&_MetaTable.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string tableAddress) returns(bool)
func (_MetaTable *MetaTableTransactor) Add(opts *bind.TransactOpts, tableName string, tableAddress string) (*types.Transaction, error) {
	return _MetaTable.contract.Transact(opts, "add", tableName, tableAddress)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string tableAddress) returns(bool)
func (_MetaTable *MetaTableSession) Add(tableName string, tableAddress string) (*types.Transaction, error) {
	return _MetaTable.Contract.Add(&_MetaTable.TransactOpts, tableName, tableAddress)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string tableName, string tableAddress) returns(bool)
func (_MetaTable *MetaTableTransactorSession) Add(tableName string, tableAddress string) (*types.Transaction, error) {
	return _MetaTable.Contract.Add(&_MetaTable.TransactOpts, tableName, tableAddress)
}

// MetaTableNewTableIterator is returned from FilterNewTable and is used to iterate over the raw logs and unpacked data for NewTable events raised by the MetaTable contract.
type MetaTableNewTableIterator struct {
	Event *MetaTableNewTable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaTableNewTableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaTableNewTable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaTableNewTable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaTableNewTableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaTableNewTableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaTableNewTable represents a NewTable event raised by the MetaTable contract.
type MetaTableNewTable struct {
	Name string
	Addr string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTable is a free log retrieval operation binding the contract event 0x18e448d301b2feb12847d36273372ac332ce9fc0de8d1137c1d9a4e96b569bd9.
//
// Solidity: event newTable(string name, string addr)
func (_MetaTable *MetaTableFilterer) FilterNewTable(opts *bind.FilterOpts) (*MetaTableNewTableIterator, error) {

	logs, sub, err := _MetaTable.contract.FilterLogs(opts, "newTable")
	if err != nil {
		return nil, err
	}
	return &MetaTableNewTableIterator{contract: _MetaTable.contract, event: "newTable", logs: logs, sub: sub}, nil
}

// WatchNewTable is a free log subscription operation binding the contract event 0x18e448d301b2feb12847d36273372ac332ce9fc0de8d1137c1d9a4e96b569bd9.
//
// Solidity: event newTable(string name, string addr)
func (_MetaTable *MetaTableFilterer) WatchNewTable(opts *bind.WatchOpts, sink chan<- *MetaTableNewTable) (event.Subscription, error) {

	logs, sub, err := _MetaTable.contract.WatchLogs(opts, "newTable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaTableNewTable)
				if err := _MetaTable.contract.UnpackLog(event, "newTable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTable is a log parse operation binding the contract event 0x18e448d301b2feb12847d36273372ac332ce9fc0de8d1137c1d9a4e96b569bd9.
//
// Solidity: event newTable(string name, string addr)
func (_MetaTable *MetaTableFilterer) ParseNewTable(log types.Log) (*MetaTableNewTable, error) {
	event := new(MetaTableNewTable)
	if err := _MetaTable.contract.UnpackLog(event, "newTable", log); err != nil {
		return nil, err
	}
	return event, nil
}
