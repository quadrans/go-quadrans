// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"toCheck\",\"type\":\"address\"}],\"name\":\"isAuthority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAuthority\",\"type\":\"address\"},{\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"addAuthority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorities\",\"outputs\":[{\"name\":\"outArray_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\"}],\"name\":\"removeAuthority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160008073f1a565c601ce8c57f7974c7f6fefffc26bffb28c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600080734f0a7b562e8a664f6c9570fbe90530964d36198473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600080738ae646dad76acc29c31d094e42bf2684c734ab9b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160008073be5b88ab4a4d063b6055c3515017e1a1cf972c8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160008073eb453620fb7518c7cc8ad8b7030761b23a59fc7e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600173f1a565c601ce8c57f7974c7f6fefffc26bffb28c90806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506001734f0a7b562e8a664f6c9570fbe90530964d36198490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506001738ae646dad76acc29c31d094e42bf2684c734ab9b90806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600173be5b88ab4a4d063b6055c3515017e1a1cf972c8b90806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600173eb453620fb7518c7cc8ad8b7030761b23a59fc7e90806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506109938061059d6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632330f2471461005c578063893d20e8146100b8578063bbcf1a5d14610102578063c21b48651461017e578063d544e010146101dd575b600080fd5b61009e6004803603602081101561007257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610239565b604051808215151515815260200191505060405180910390f35b6100c06102d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101646004803603604081101561011857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102fa565b604051808215151515815260200191505060405180910390f35b6101866104e3565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101c95780820151818401526020810190506101ae565b505050509050019250505060405180910390f35b61021f600480360360208110156101f357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105c4565b604051808215151515815260200191505060405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff166000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561035857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156103f557600090506104dd565b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600190505b92915050565b60606001805490506040519080825280602002602001820160405280156105195781602001602082028038833980820191505090505b50905060008090505b6001805490508110156105c05760018181548110151561053e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828281518110151561057757fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610522565b5090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561062257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156106be576000905061074b565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061074582610750565b50600190505b919050565b60008061075c83610873565b90506000811415151561076e57600080fd5b60018080549050038110156108155760018080805490500381548110151561079257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001828154811015156107cc57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60018080805490500381548110151561082a57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600180548091906001900361086c9190610916565b5050919050565b600080600090505b60018054905081101561090b578273ffffffffffffffffffffffffffffffffffffffff166001828154811015156108ae57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156108fe5780915050610911565b808060010191505061087b565b50600090505b919050565b81548183558181111561093d5781836000526020600020918201910161093c9190610942565b5b505050565b61096491905b80821115610960576000816000905550600101610948565b5090565b9056fea165627a7a723058209ba81e76a614cf8325b07109063744614c095c84356a43f391c9ed8e0280bbfb0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetAuthorities is a free data retrieval call binding the contract method 0xc21b4865.
//
// Solidity: function getAuthorities() constant returns(address[] outArray_)
func (_Contract *ContractCaller) GetAuthorities(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getAuthorities")
	return *ret0, err
}

// GetAuthorities is a free data retrieval call binding the contract method 0xc21b4865.
//
// Solidity: function getAuthorities() constant returns(address[] outArray_)
func (_Contract *ContractSession) GetAuthorities() ([]common.Address, error) {
	return _Contract.Contract.GetAuthorities(&_Contract.CallOpts)
}

// GetAuthorities is a free data retrieval call binding the contract method 0xc21b4865.
//
// Solidity: function getAuthorities() constant returns(address[] outArray_)
func (_Contract *ContractCallerSession) GetAuthorities() ([]common.Address, error) {
	return _Contract.Contract.GetAuthorities(&_Contract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address owner)
func (_Contract *ContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address owner)
func (_Contract *ContractSession) GetOwner() (common.Address, error) {
	return _Contract.Contract.GetOwner(&_Contract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address owner)
func (_Contract *ContractCallerSession) GetOwner() (common.Address, error) {
	return _Contract.Contract.GetOwner(&_Contract.CallOpts)
}

// IsAuthority is a free data retrieval call binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(address toCheck) constant returns(bool)
func (_Contract *ContractCaller) IsAuthority(opts *bind.CallOpts, toCheck common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isAuthority", toCheck)
	return *ret0, err
}

// IsAuthority is a free data retrieval call binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(address toCheck) constant returns(bool)
func (_Contract *ContractSession) IsAuthority(toCheck common.Address) (bool, error) {
	return _Contract.Contract.IsAuthority(&_Contract.CallOpts, toCheck)
}

// IsAuthority is a free data retrieval call binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(address toCheck) constant returns(bool)
func (_Contract *ContractCallerSession) IsAuthority(toCheck common.Address) (bool, error) {
	return _Contract.Contract.IsAuthority(&_Contract.CallOpts, toCheck)
}

// AddAuthority is a paid mutator transaction binding the contract method 0xbbcf1a5d.
//
// Solidity: function addAuthority(address newAuthority, address ethAddress) returns(bool)
func (_Contract *ContractTransactor) AddAuthority(opts *bind.TransactOpts, newAuthority common.Address, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addAuthority", newAuthority, ethAddress)
}

// AddAuthority is a paid mutator transaction binding the contract method 0xbbcf1a5d.
//
// Solidity: function addAuthority(address newAuthority, address ethAddress) returns(bool)
func (_Contract *ContractSession) AddAuthority(newAuthority common.Address, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuthority(&_Contract.TransactOpts, newAuthority, ethAddress)
}

// AddAuthority is a paid mutator transaction binding the contract method 0xbbcf1a5d.
//
// Solidity: function addAuthority(address newAuthority, address ethAddress) returns(bool)
func (_Contract *ContractTransactorSession) AddAuthority(newAuthority common.Address, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuthority(&_Contract.TransactOpts, newAuthority, ethAddress)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(address oldAuthority) returns(bool)
func (_Contract *ContractTransactor) RemoveAuthority(opts *bind.TransactOpts, oldAuthority common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeAuthority", oldAuthority)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(address oldAuthority) returns(bool)
func (_Contract *ContractSession) RemoveAuthority(oldAuthority common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuthority(&_Contract.TransactOpts, oldAuthority)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(address oldAuthority) returns(bool)
func (_Contract *ContractTransactorSession) RemoveAuthority(oldAuthority common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuthority(&_Contract.TransactOpts, oldAuthority)
}
