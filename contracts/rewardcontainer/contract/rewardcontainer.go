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

// RewardContainerABI is the input ABI used to generate the binding from.
const RewardContainerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTotalRewarded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"changeRewardImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActualReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ChangedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"rewardImplementation\",\"type\":\"address\"}],\"name\":\"ChangedRewardImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mined\",\"type\":\"event\"}]"

// RewardContainerBin is the compiled bytecode used for deploying new contracts.
const RewardContainerBin = `608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061064b806100606000396000f3fe608060405260043610610072576000357c0100000000000000000000000000000000000000000000000000000000900480630e5ee5ee146100ab5780631a957fe3146100d657806383197ef014610127578063a6f9dae11461013e578063d0679d341461018f578063ea6a62f1146101ea575b7f4229d50c63dbdc5551dd68e0a9879b01ac250cb31feaeba7588533462e6c7aaa346040518082815260200191505060405180910390a1005b3480156100b757600080fd5b506100c0610215565b6040518082815260200191505060405180910390f35b3480156100e257600080fd5b50610125600480360360208110156100f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061021f565b005b34801561013357600080fd5b5061013c610301565b005b34801561014a57600080fd5b5061018d6004803603602081101561016157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610396565b005b34801561019b57600080fd5b506101e8600480360360408110156101b257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610477565b005b3480156101f657600080fd5b506101ff610600565b6040518082815260200191505060405180910390f35b6000600254905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561027a57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f2e25b8052a53bf10dbccee778703f1e8356d76602c3c4c0f55261eac3098f9d260405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561035c57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103f157600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fa701229f4b9ddf00aa1c7228d248e6320ee7c581d856ddfba036e73947cd0d1360405160405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156104d557600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561053157600080fd5b60003073ffffffffffffffffffffffffffffffffffffffff163111151561055757600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561059d573d6000803e3d6000fd5b50806002600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff167f510ffb4dcab972ae9d2007a58e13f1b0881776d23cd8f5cc32f8c5be2dbf70d2826040518082815260200191505060405180910390a25050565b60003073ffffffffffffffffffffffffffffffffffffffff163190509056fea165627a7a72305820e6ef4bb28965af8014ad5d1b5a6e2c3016b67bb4d074ccdab2b1e91cfd81c12a0029`

// DeployRewardContainer deploys a new Ethereum contract, binding an instance of RewardContainer to it.
func DeployRewardContainer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RewardContainer, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardContainerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewardContainerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardContainer{RewardContainerCaller: RewardContainerCaller{contract: contract}, RewardContainerTransactor: RewardContainerTransactor{contract: contract}, RewardContainerFilterer: RewardContainerFilterer{contract: contract}}, nil
}

// RewardContainer is an auto generated Go binding around an Ethereum contract.
type RewardContainer struct {
	RewardContainerCaller     // Read-only binding to the contract
	RewardContainerTransactor // Write-only binding to the contract
	RewardContainerFilterer   // Log filterer for contract events
}

// RewardContainerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardContainerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardContainerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardContainerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardContainerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardContainerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardContainerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardContainerSession struct {
	Contract     *RewardContainer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardContainerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardContainerCallerSession struct {
	Contract *RewardContainerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RewardContainerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardContainerTransactorSession struct {
	Contract     *RewardContainerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RewardContainerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardContainerRaw struct {
	Contract *RewardContainer // Generic contract binding to access the raw methods on
}

// RewardContainerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardContainerCallerRaw struct {
	Contract *RewardContainerCaller // Generic read-only contract binding to access the raw methods on
}

// RewardContainerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardContainerTransactorRaw struct {
	Contract *RewardContainerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardContainer creates a new instance of RewardContainer, bound to a specific deployed contract.
func NewRewardContainer(address common.Address, backend bind.ContractBackend) (*RewardContainer, error) {
	contract, err := bindRewardContainer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardContainer{RewardContainerCaller: RewardContainerCaller{contract: contract}, RewardContainerTransactor: RewardContainerTransactor{contract: contract}, RewardContainerFilterer: RewardContainerFilterer{contract: contract}}, nil
}

// NewRewardContainerCaller creates a new read-only instance of RewardContainer, bound to a specific deployed contract.
func NewRewardContainerCaller(address common.Address, caller bind.ContractCaller) (*RewardContainerCaller, error) {
	contract, err := bindRewardContainer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardContainerCaller{contract: contract}, nil
}

// NewRewardContainerTransactor creates a new write-only instance of RewardContainer, bound to a specific deployed contract.
func NewRewardContainerTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardContainerTransactor, error) {
	contract, err := bindRewardContainer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardContainerTransactor{contract: contract}, nil
}

// NewRewardContainerFilterer creates a new log filterer instance of RewardContainer, bound to a specific deployed contract.
func NewRewardContainerFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardContainerFilterer, error) {
	contract, err := bindRewardContainer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardContainerFilterer{contract: contract}, nil
}

// bindRewardContainer binds a generic wrapper to an already deployed contract.
func bindRewardContainer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardContainerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardContainer *RewardContainerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardContainer.Contract.RewardContainerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardContainer *RewardContainerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardContainer.Contract.RewardContainerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardContainer *RewardContainerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardContainer.Contract.RewardContainerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardContainer *RewardContainerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardContainer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardContainer *RewardContainerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardContainer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardContainer *RewardContainerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardContainer.Contract.contract.Transact(opts, method, params...)
}

// GetActualReward is a free data retrieval call binding the contract method 0xea6a62f1.
//
// Solidity: function getActualReward() constant returns(uint256)
func (_RewardContainer *RewardContainerCaller) GetActualReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RewardContainer.contract.Call(opts, out, "getActualReward")
	return *ret0, err
}

// GetActualReward is a free data retrieval call binding the contract method 0xea6a62f1.
//
// Solidity: function getActualReward() constant returns(uint256)
func (_RewardContainer *RewardContainerSession) GetActualReward() (*big.Int, error) {
	return _RewardContainer.Contract.GetActualReward(&_RewardContainer.CallOpts)
}

// GetActualReward is a free data retrieval call binding the contract method 0xea6a62f1.
//
// Solidity: function getActualReward() constant returns(uint256)
func (_RewardContainer *RewardContainerCallerSession) GetActualReward() (*big.Int, error) {
	return _RewardContainer.Contract.GetActualReward(&_RewardContainer.CallOpts)
}

// GetTotalRewarded is a free data retrieval call binding the contract method 0x0e5ee5ee.
//
// Solidity: function getTotalRewarded() constant returns(uint256)
func (_RewardContainer *RewardContainerCaller) GetTotalRewarded(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RewardContainer.contract.Call(opts, out, "getTotalRewarded")
	return *ret0, err
}

// GetTotalRewarded is a free data retrieval call binding the contract method 0x0e5ee5ee.
//
// Solidity: function getTotalRewarded() constant returns(uint256)
func (_RewardContainer *RewardContainerSession) GetTotalRewarded() (*big.Int, error) {
	return _RewardContainer.Contract.GetTotalRewarded(&_RewardContainer.CallOpts)
}

// GetTotalRewarded is a free data retrieval call binding the contract method 0x0e5ee5ee.
//
// Solidity: function getTotalRewarded() constant returns(uint256)
func (_RewardContainer *RewardContainerCallerSession) GetTotalRewarded() (*big.Int, error) {
	return _RewardContainer.Contract.GetTotalRewarded(&_RewardContainer.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_RewardContainer *RewardContainerTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewardContainer.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_RewardContainer *RewardContainerSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RewardContainer.Contract.ChangeOwner(&_RewardContainer.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_RewardContainer *RewardContainerTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RewardContainer.Contract.ChangeOwner(&_RewardContainer.TransactOpts, newOwner)
}

// ChangeRewardImplementation is a paid mutator transaction binding the contract method 0x1a957fe3.
//
// Solidity: function changeRewardImplementation(address newImpl) returns()
func (_RewardContainer *RewardContainerTransactor) ChangeRewardImplementation(opts *bind.TransactOpts, newImpl common.Address) (*types.Transaction, error) {
	return _RewardContainer.contract.Transact(opts, "changeRewardImplementation", newImpl)
}

// ChangeRewardImplementation is a paid mutator transaction binding the contract method 0x1a957fe3.
//
// Solidity: function changeRewardImplementation(address newImpl) returns()
func (_RewardContainer *RewardContainerSession) ChangeRewardImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _RewardContainer.Contract.ChangeRewardImplementation(&_RewardContainer.TransactOpts, newImpl)
}

// ChangeRewardImplementation is a paid mutator transaction binding the contract method 0x1a957fe3.
//
// Solidity: function changeRewardImplementation(address newImpl) returns()
func (_RewardContainer *RewardContainerTransactorSession) ChangeRewardImplementation(newImpl common.Address) (*types.Transaction, error) {
	return _RewardContainer.Contract.ChangeRewardImplementation(&_RewardContainer.TransactOpts, newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RewardContainer *RewardContainerTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardContainer.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RewardContainer *RewardContainerSession) Destroy() (*types.Transaction, error) {
	return _RewardContainer.Contract.Destroy(&_RewardContainer.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RewardContainer *RewardContainerTransactorSession) Destroy() (*types.Transaction, error) {
	return _RewardContainer.Contract.Destroy(&_RewardContainer.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address to, uint256 amount) returns()
func (_RewardContainer *RewardContainerTransactor) Send(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardContainer.contract.Transact(opts, "send", to, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address to, uint256 amount) returns()
func (_RewardContainer *RewardContainerSession) Send(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardContainer.Contract.Send(&_RewardContainer.TransactOpts, to, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address to, uint256 amount) returns()
func (_RewardContainer *RewardContainerTransactorSession) Send(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardContainer.Contract.Send(&_RewardContainer.TransactOpts, to, amount)
}

// RewardContainerChangedOwnerIterator is returned from FilterChangedOwner and is used to iterate over the raw logs and unpacked data for ChangedOwner events raised by the RewardContainer contract.
type RewardContainerChangedOwnerIterator struct {
	Event *RewardContainerChangedOwner // Event containing the contract specifics and raw log

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
func (it *RewardContainerChangedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardContainerChangedOwner)
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
		it.Event = new(RewardContainerChangedOwner)
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
func (it *RewardContainerChangedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardContainerChangedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardContainerChangedOwner represents a ChangedOwner event raised by the RewardContainer contract.
type RewardContainerChangedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedOwner is a free log retrieval operation binding the contract event 0xa701229f4b9ddf00aa1c7228d248e6320ee7c581d856ddfba036e73947cd0d13.
//
// Solidity: event ChangedOwner(address indexed owner)
func (_RewardContainer *RewardContainerFilterer) FilterChangedOwner(opts *bind.FilterOpts, owner []common.Address) (*RewardContainerChangedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RewardContainer.contract.FilterLogs(opts, "ChangedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &RewardContainerChangedOwnerIterator{contract: _RewardContainer.contract, event: "ChangedOwner", logs: logs, sub: sub}, nil
}

// WatchChangedOwner is a free log subscription operation binding the contract event 0xa701229f4b9ddf00aa1c7228d248e6320ee7c581d856ddfba036e73947cd0d13.
//
// Solidity: event ChangedOwner(address indexed owner)
func (_RewardContainer *RewardContainerFilterer) WatchChangedOwner(opts *bind.WatchOpts, sink chan<- *RewardContainerChangedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RewardContainer.contract.WatchLogs(opts, "ChangedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardContainerChangedOwner)
				if err := _RewardContainer.contract.UnpackLog(event, "ChangedOwner", log); err != nil {
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

// RewardContainerChangedRewardImplementationIterator is returned from FilterChangedRewardImplementation and is used to iterate over the raw logs and unpacked data for ChangedRewardImplementation events raised by the RewardContainer contract.
type RewardContainerChangedRewardImplementationIterator struct {
	Event *RewardContainerChangedRewardImplementation // Event containing the contract specifics and raw log

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
func (it *RewardContainerChangedRewardImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardContainerChangedRewardImplementation)
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
		it.Event = new(RewardContainerChangedRewardImplementation)
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
func (it *RewardContainerChangedRewardImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardContainerChangedRewardImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardContainerChangedRewardImplementation represents a ChangedRewardImplementation event raised by the RewardContainer contract.
type RewardContainerChangedRewardImplementation struct {
	RewardImplementation common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterChangedRewardImplementation is a free log retrieval operation binding the contract event 0x2e25b8052a53bf10dbccee778703f1e8356d76602c3c4c0f55261eac3098f9d2.
//
// Solidity: event ChangedRewardImplementation(address indexed rewardImplementation)
func (_RewardContainer *RewardContainerFilterer) FilterChangedRewardImplementation(opts *bind.FilterOpts, rewardImplementation []common.Address) (*RewardContainerChangedRewardImplementationIterator, error) {

	var rewardImplementationRule []interface{}
	for _, rewardImplementationItem := range rewardImplementation {
		rewardImplementationRule = append(rewardImplementationRule, rewardImplementationItem)
	}

	logs, sub, err := _RewardContainer.contract.FilterLogs(opts, "ChangedRewardImplementation", rewardImplementationRule)
	if err != nil {
		return nil, err
	}
	return &RewardContainerChangedRewardImplementationIterator{contract: _RewardContainer.contract, event: "ChangedRewardImplementation", logs: logs, sub: sub}, nil
}

// WatchChangedRewardImplementation is a free log subscription operation binding the contract event 0x2e25b8052a53bf10dbccee778703f1e8356d76602c3c4c0f55261eac3098f9d2.
//
// Solidity: event ChangedRewardImplementation(address indexed rewardImplementation)
func (_RewardContainer *RewardContainerFilterer) WatchChangedRewardImplementation(opts *bind.WatchOpts, sink chan<- *RewardContainerChangedRewardImplementation, rewardImplementation []common.Address) (event.Subscription, error) {

	var rewardImplementationRule []interface{}
	for _, rewardImplementationItem := range rewardImplementation {
		rewardImplementationRule = append(rewardImplementationRule, rewardImplementationItem)
	}

	logs, sub, err := _RewardContainer.contract.WatchLogs(opts, "ChangedRewardImplementation", rewardImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardContainerChangedRewardImplementation)
				if err := _RewardContainer.contract.UnpackLog(event, "ChangedRewardImplementation", log); err != nil {
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

// RewardContainerMinedIterator is returned from FilterMined and is used to iterate over the raw logs and unpacked data for Mined events raised by the RewardContainer contract.
type RewardContainerMinedIterator struct {
	Event *RewardContainerMined // Event containing the contract specifics and raw log

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
func (it *RewardContainerMinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardContainerMined)
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
		it.Event = new(RewardContainerMined)
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
func (it *RewardContainerMinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardContainerMinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardContainerMined represents a Mined event raised by the RewardContainer contract.
type RewardContainerMined struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMined is a free log retrieval operation binding the contract event 0x4229d50c63dbdc5551dd68e0a9879b01ac250cb31feaeba7588533462e6c7aaa.
//
// Solidity: event Mined(uint256 amount)
func (_RewardContainer *RewardContainerFilterer) FilterMined(opts *bind.FilterOpts) (*RewardContainerMinedIterator, error) {

	logs, sub, err := _RewardContainer.contract.FilterLogs(opts, "Mined")
	if err != nil {
		return nil, err
	}
	return &RewardContainerMinedIterator{contract: _RewardContainer.contract, event: "Mined", logs: logs, sub: sub}, nil
}

// WatchMined is a free log subscription operation binding the contract event 0x4229d50c63dbdc5551dd68e0a9879b01ac250cb31feaeba7588533462e6c7aaa.
//
// Solidity: event Mined(uint256 amount)
func (_RewardContainer *RewardContainerFilterer) WatchMined(opts *bind.WatchOpts, sink chan<- *RewardContainerMined) (event.Subscription, error) {

	logs, sub, err := _RewardContainer.contract.WatchLogs(opts, "Mined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardContainerMined)
				if err := _RewardContainer.contract.UnpackLog(event, "Mined", log); err != nil {
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

// RewardContainerSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the RewardContainer contract.
type RewardContainerSentIterator struct {
	Event *RewardContainerSent // Event containing the contract specifics and raw log

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
func (it *RewardContainerSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardContainerSent)
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
		it.Event = new(RewardContainerSent)
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
func (it *RewardContainerSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardContainerSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardContainerSent represents a Sent event raised by the RewardContainer contract.
type RewardContainerSent struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x510ffb4dcab972ae9d2007a58e13f1b0881776d23cd8f5cc32f8c5be2dbf70d2.
//
// Solidity: event Sent(address indexed to, uint256 amount)
func (_RewardContainer *RewardContainerFilterer) FilterSent(opts *bind.FilterOpts, to []common.Address) (*RewardContainerSentIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RewardContainer.contract.FilterLogs(opts, "Sent", toRule)
	if err != nil {
		return nil, err
	}
	return &RewardContainerSentIterator{contract: _RewardContainer.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x510ffb4dcab972ae9d2007a58e13f1b0881776d23cd8f5cc32f8c5be2dbf70d2.
//
// Solidity: event Sent(address indexed to, uint256 amount)
func (_RewardContainer *RewardContainerFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *RewardContainerSent, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RewardContainer.contract.WatchLogs(opts, "Sent", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardContainerSent)
				if err := _RewardContainer.contract.UnpackLog(event, "Sent", log); err != nil {
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
