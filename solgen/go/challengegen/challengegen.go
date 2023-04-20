// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengegen

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// OldChallengeLibChallenge is an auto generated low-level Go binding around an user-defined struct.
type OldChallengeLibChallenge struct {
	Current            OldChallengeLibParticipant
	Next               OldChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}

// OldChallengeLibParticipant is an auto generated low-level Go binding around an user-defined struct.
type OldChallengeLibParticipant struct {
	Addr     common.Address
	TimeLeft *big.Int
}

// OldChallengeLibSegmentSelection is an auto generated low-level Go binding around an user-defined struct.
type OldChallengeLibSegmentSelection struct {
	OldSegmentsStart  *big.Int
	OldSegmentsLength *big.Int
	OldSegments       [][32]byte
	ChallengePosition *big.Int
}

// IOldChallengeManagerMetaData contains all meta data concerning the IOldChallengeManager contract.
var IOldChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIOldChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumOldChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structOldChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOldChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOldChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IOldChallengeManagerMetaData.ABI instead.
var IOldChallengeManagerABI = IOldChallengeManagerMetaData.ABI

// IOldChallengeManager is an auto generated Go binding around an Ethereum contract.
type IOldChallengeManager struct {
	IOldChallengeManagerCaller     // Read-only binding to the contract
	IOldChallengeManagerTransactor // Write-only binding to the contract
	IOldChallengeManagerFilterer   // Log filterer for contract events
}

// IOldChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOldChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOldChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOldChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOldChallengeManagerSession struct {
	Contract     *IOldChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IOldChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOldChallengeManagerCallerSession struct {
	Contract *IOldChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IOldChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOldChallengeManagerTransactorSession struct {
	Contract     *IOldChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IOldChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOldChallengeManagerRaw struct {
	Contract *IOldChallengeManager // Generic contract binding to access the raw methods on
}

// IOldChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOldChallengeManagerCallerRaw struct {
	Contract *IOldChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IOldChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOldChallengeManagerTransactorRaw struct {
	Contract *IOldChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOldChallengeManager creates a new instance of IOldChallengeManager, bound to a specific deployed contract.
func NewIOldChallengeManager(address common.Address, backend bind.ContractBackend) (*IOldChallengeManager, error) {
	contract, err := bindIOldChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManager{IOldChallengeManagerCaller: IOldChallengeManagerCaller{contract: contract}, IOldChallengeManagerTransactor: IOldChallengeManagerTransactor{contract: contract}, IOldChallengeManagerFilterer: IOldChallengeManagerFilterer{contract: contract}}, nil
}

// NewIOldChallengeManagerCaller creates a new read-only instance of IOldChallengeManager, bound to a specific deployed contract.
func NewIOldChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IOldChallengeManagerCaller, error) {
	contract, err := bindIOldChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerCaller{contract: contract}, nil
}

// NewIOldChallengeManagerTransactor creates a new write-only instance of IOldChallengeManager, bound to a specific deployed contract.
func NewIOldChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOldChallengeManagerTransactor, error) {
	contract, err := bindIOldChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerTransactor{contract: contract}, nil
}

// NewIOldChallengeManagerFilterer creates a new log filterer instance of IOldChallengeManager, bound to a specific deployed contract.
func NewIOldChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOldChallengeManagerFilterer, error) {
	contract, err := bindIOldChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerFilterer{contract: contract}, nil
}

// bindIOldChallengeManager binds a generic wrapper to an already deployed contract.
func bindIOldChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOldChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOldChallengeManager *IOldChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOldChallengeManager.Contract.IOldChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOldChallengeManager *IOldChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.IOldChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOldChallengeManager *IOldChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.IOldChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOldChallengeManager *IOldChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOldChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOldChallengeManager *IOldChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOldChallengeManager *IOldChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IOldChallengeManager *IOldChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex_ uint64) (OldChallengeLibChallenge, error) {
	var out []interface{}
	err := _IOldChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex_)

	if err != nil {
		return *new(OldChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(OldChallengeLibChallenge)).(*OldChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IOldChallengeManager *IOldChallengeManagerSession) ChallengeInfo(challengeIndex_ uint64) (OldChallengeLibChallenge, error) {
	return _IOldChallengeManager.Contract.ChallengeInfo(&_IOldChallengeManager.CallOpts, challengeIndex_)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IOldChallengeManager *IOldChallengeManagerCallerSession) ChallengeInfo(challengeIndex_ uint64) (OldChallengeLibChallenge, error) {
	return _IOldChallengeManager.Contract.ChallengeInfo(&_IOldChallengeManager.CallOpts, challengeIndex_)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IOldChallengeManager *IOldChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _IOldChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IOldChallengeManager *IOldChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IOldChallengeManager.Contract.CurrentResponder(&_IOldChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IOldChallengeManager *IOldChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IOldChallengeManager.Contract.CurrentResponder(&_IOldChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IOldChallengeManager *IOldChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _IOldChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IOldChallengeManager *IOldChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IOldChallengeManager.Contract.IsTimedOut(&_IOldChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IOldChallengeManager *IOldChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IOldChallengeManager.Contract.IsTimedOut(&_IOldChallengeManager.CallOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.ClearChallenge(&_IOldChallengeManager.TransactOpts, challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactorSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.ClearChallenge(&_IOldChallengeManager.TransactOpts, challengeIndex_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IOldChallengeManager *IOldChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IOldChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IOldChallengeManager *IOldChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.CreateChallenge(&_IOldChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IOldChallengeManager *IOldChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.CreateChallenge(&_IOldChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IOldChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IOldChallengeManager *IOldChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.Initialize(&_IOldChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.Initialize(&_IOldChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.contract.Transact(opts, "timeout", challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.Timeout(&_IOldChallengeManager.TransactOpts, challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IOldChallengeManager *IOldChallengeManagerTransactorSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IOldChallengeManager.Contract.Timeout(&_IOldChallengeManager.TransactOpts, challengeIndex_)
}

// IOldChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the IOldChallengeManager contract.
type IOldChallengeManagerBisectedIterator struct {
	Event *IOldChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *IOldChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOldChallengeManagerBisected)
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
		it.Event = new(IOldChallengeManagerBisected)
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
func (it *IOldChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOldChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOldChallengeManagerBisected represents a Bisected event raised by the IOldChallengeManager contract.
type IOldChallengeManagerBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*IOldChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerBisectedIterator{contract: _IOldChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *IOldChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOldChallengeManagerBisected)
				if err := _IOldChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) ParseBisected(log types.Log) (*IOldChallengeManagerBisected, error) {
	event := new(IOldChallengeManagerBisected)
	if err := _IOldChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOldChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the IOldChallengeManager contract.
type IOldChallengeManagerChallengeEndedIterator struct {
	Event *IOldChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *IOldChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOldChallengeManagerChallengeEnded)
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
		it.Event = new(IOldChallengeManagerChallengeEnded)
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
func (it *IOldChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOldChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOldChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the IOldChallengeManager contract.
type IOldChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*IOldChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerChallengeEndedIterator{contract: _IOldChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *IOldChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOldChallengeManagerChallengeEnded)
				if err := _IOldChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*IOldChallengeManagerChallengeEnded, error) {
	event := new(IOldChallengeManagerChallengeEnded)
	if err := _IOldChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOldChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the IOldChallengeManager contract.
type IOldChallengeManagerExecutionChallengeBegunIterator struct {
	Event *IOldChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *IOldChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOldChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(IOldChallengeManagerExecutionChallengeBegun)
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
func (it *IOldChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOldChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOldChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the IOldChallengeManager contract.
type IOldChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*IOldChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerExecutionChallengeBegunIterator{contract: _IOldChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *IOldChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOldChallengeManagerExecutionChallengeBegun)
				if err := _IOldChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*IOldChallengeManagerExecutionChallengeBegun, error) {
	event := new(IOldChallengeManagerExecutionChallengeBegun)
	if err := _IOldChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOldChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the IOldChallengeManager contract.
type IOldChallengeManagerInitiatedChallengeIterator struct {
	Event *IOldChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *IOldChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOldChallengeManagerInitiatedChallenge)
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
		it.Event = new(IOldChallengeManagerInitiatedChallenge)
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
func (it *IOldChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOldChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOldChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the IOldChallengeManager contract.
type IOldChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*IOldChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerInitiatedChallengeIterator{contract: _IOldChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *IOldChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOldChallengeManagerInitiatedChallenge)
				if err := _IOldChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*IOldChallengeManagerInitiatedChallenge, error) {
	event := new(IOldChallengeManagerInitiatedChallenge)
	if err := _IOldChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOldChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the IOldChallengeManager contract.
type IOldChallengeManagerOneStepProofCompletedIterator struct {
	Event *IOldChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *IOldChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOldChallengeManagerOneStepProofCompleted)
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
		it.Event = new(IOldChallengeManagerOneStepProofCompleted)
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
func (it *IOldChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOldChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOldChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the IOldChallengeManager contract.
type IOldChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*IOldChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeManagerOneStepProofCompletedIterator{contract: _IOldChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *IOldChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IOldChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOldChallengeManagerOneStepProofCompleted)
				if err := _IOldChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IOldChallengeManager *IOldChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*IOldChallengeManagerOneStepProofCompleted, error) {
	event := new(IOldChallengeManagerOneStepProofCompleted)
	if err := _IOldChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOldChallengeResultReceiverMetaData contains all meta data concerning the IOldChallengeResultReceiver contract.
var IOldChallengeResultReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOldChallengeResultReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOldChallengeResultReceiverMetaData.ABI instead.
var IOldChallengeResultReceiverABI = IOldChallengeResultReceiverMetaData.ABI

// IOldChallengeResultReceiver is an auto generated Go binding around an Ethereum contract.
type IOldChallengeResultReceiver struct {
	IOldChallengeResultReceiverCaller     // Read-only binding to the contract
	IOldChallengeResultReceiverTransactor // Write-only binding to the contract
	IOldChallengeResultReceiverFilterer   // Log filterer for contract events
}

// IOldChallengeResultReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOldChallengeResultReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeResultReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOldChallengeResultReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeResultReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOldChallengeResultReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOldChallengeResultReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOldChallengeResultReceiverSession struct {
	Contract     *IOldChallengeResultReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IOldChallengeResultReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOldChallengeResultReceiverCallerSession struct {
	Contract *IOldChallengeResultReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IOldChallengeResultReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOldChallengeResultReceiverTransactorSession struct {
	Contract     *IOldChallengeResultReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IOldChallengeResultReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOldChallengeResultReceiverRaw struct {
	Contract *IOldChallengeResultReceiver // Generic contract binding to access the raw methods on
}

// IOldChallengeResultReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOldChallengeResultReceiverCallerRaw struct {
	Contract *IOldChallengeResultReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IOldChallengeResultReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOldChallengeResultReceiverTransactorRaw struct {
	Contract *IOldChallengeResultReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOldChallengeResultReceiver creates a new instance of IOldChallengeResultReceiver, bound to a specific deployed contract.
func NewIOldChallengeResultReceiver(address common.Address, backend bind.ContractBackend) (*IOldChallengeResultReceiver, error) {
	contract, err := bindIOldChallengeResultReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeResultReceiver{IOldChallengeResultReceiverCaller: IOldChallengeResultReceiverCaller{contract: contract}, IOldChallengeResultReceiverTransactor: IOldChallengeResultReceiverTransactor{contract: contract}, IOldChallengeResultReceiverFilterer: IOldChallengeResultReceiverFilterer{contract: contract}}, nil
}

// NewIOldChallengeResultReceiverCaller creates a new read-only instance of IOldChallengeResultReceiver, bound to a specific deployed contract.
func NewIOldChallengeResultReceiverCaller(address common.Address, caller bind.ContractCaller) (*IOldChallengeResultReceiverCaller, error) {
	contract, err := bindIOldChallengeResultReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeResultReceiverCaller{contract: contract}, nil
}

// NewIOldChallengeResultReceiverTransactor creates a new write-only instance of IOldChallengeResultReceiver, bound to a specific deployed contract.
func NewIOldChallengeResultReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOldChallengeResultReceiverTransactor, error) {
	contract, err := bindIOldChallengeResultReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeResultReceiverTransactor{contract: contract}, nil
}

// NewIOldChallengeResultReceiverFilterer creates a new log filterer instance of IOldChallengeResultReceiver, bound to a specific deployed contract.
func NewIOldChallengeResultReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOldChallengeResultReceiverFilterer, error) {
	contract, err := bindIOldChallengeResultReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOldChallengeResultReceiverFilterer{contract: contract}, nil
}

// bindIOldChallengeResultReceiver binds a generic wrapper to an already deployed contract.
func bindIOldChallengeResultReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOldChallengeResultReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOldChallengeResultReceiver.Contract.IOldChallengeResultReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.IOldChallengeResultReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.IOldChallengeResultReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOldChallengeResultReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverTransactor) CompleteChallenge(opts *bind.TransactOpts, challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.contract.Transact(opts, "completeChallenge", challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.CompleteChallenge(&_IOldChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IOldChallengeResultReceiver *IOldChallengeResultReceiverTransactorSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IOldChallengeResultReceiver.Contract.CompleteChallenge(&_IOldChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}

// OldChallengeLibMetaData contains all meta data concerning the OldChallengeLib contract.
var OldChallengeLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200bd6c3669342a661f37cd2646c328f8cee80be5b501b4111d6cbf813c8f5a4c064736f6c63430008110033",
}

// OldChallengeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OldChallengeLibMetaData.ABI instead.
var OldChallengeLibABI = OldChallengeLibMetaData.ABI

// OldChallengeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OldChallengeLibMetaData.Bin instead.
var OldChallengeLibBin = OldChallengeLibMetaData.Bin

// DeployOldChallengeLib deploys a new Ethereum contract, binding an instance of OldChallengeLib to it.
func DeployOldChallengeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldChallengeLib, error) {
	parsed, err := OldChallengeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OldChallengeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldChallengeLib{OldChallengeLibCaller: OldChallengeLibCaller{contract: contract}, OldChallengeLibTransactor: OldChallengeLibTransactor{contract: contract}, OldChallengeLibFilterer: OldChallengeLibFilterer{contract: contract}}, nil
}

// OldChallengeLib is an auto generated Go binding around an Ethereum contract.
type OldChallengeLib struct {
	OldChallengeLibCaller     // Read-only binding to the contract
	OldChallengeLibTransactor // Write-only binding to the contract
	OldChallengeLibFilterer   // Log filterer for contract events
}

// OldChallengeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldChallengeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldChallengeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldChallengeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldChallengeLibSession struct {
	Contract     *OldChallengeLib  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OldChallengeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldChallengeLibCallerSession struct {
	Contract *OldChallengeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OldChallengeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldChallengeLibTransactorSession struct {
	Contract     *OldChallengeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OldChallengeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldChallengeLibRaw struct {
	Contract *OldChallengeLib // Generic contract binding to access the raw methods on
}

// OldChallengeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldChallengeLibCallerRaw struct {
	Contract *OldChallengeLibCaller // Generic read-only contract binding to access the raw methods on
}

// OldChallengeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldChallengeLibTransactorRaw struct {
	Contract *OldChallengeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldChallengeLib creates a new instance of OldChallengeLib, bound to a specific deployed contract.
func NewOldChallengeLib(address common.Address, backend bind.ContractBackend) (*OldChallengeLib, error) {
	contract, err := bindOldChallengeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldChallengeLib{OldChallengeLibCaller: OldChallengeLibCaller{contract: contract}, OldChallengeLibTransactor: OldChallengeLibTransactor{contract: contract}, OldChallengeLibFilterer: OldChallengeLibFilterer{contract: contract}}, nil
}

// NewOldChallengeLibCaller creates a new read-only instance of OldChallengeLib, bound to a specific deployed contract.
func NewOldChallengeLibCaller(address common.Address, caller bind.ContractCaller) (*OldChallengeLibCaller, error) {
	contract, err := bindOldChallengeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldChallengeLibCaller{contract: contract}, nil
}

// NewOldChallengeLibTransactor creates a new write-only instance of OldChallengeLib, bound to a specific deployed contract.
func NewOldChallengeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OldChallengeLibTransactor, error) {
	contract, err := bindOldChallengeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldChallengeLibTransactor{contract: contract}, nil
}

// NewOldChallengeLibFilterer creates a new log filterer instance of OldChallengeLib, bound to a specific deployed contract.
func NewOldChallengeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OldChallengeLibFilterer, error) {
	contract, err := bindOldChallengeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldChallengeLibFilterer{contract: contract}, nil
}

// bindOldChallengeLib binds a generic wrapper to an already deployed contract.
func bindOldChallengeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldChallengeLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldChallengeLib *OldChallengeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldChallengeLib.Contract.OldChallengeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldChallengeLib *OldChallengeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldChallengeLib.Contract.OldChallengeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldChallengeLib *OldChallengeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldChallengeLib.Contract.OldChallengeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldChallengeLib *OldChallengeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldChallengeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldChallengeLib *OldChallengeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldChallengeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldChallengeLib *OldChallengeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldChallengeLib.Contract.contract.Transact(opts, method, params...)
}

// OldChallengeManagerMetaData contains all meta data concerning the OldChallengeManager contract.
var OldChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIOldChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumOldChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structOldChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumOldChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOldChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structOldChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIOldChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051612fc461003060003960006111f50152612fc46000f3fe608060405234801561001057600080fd5b50600436106100e05760003560e01c80639ede42b9116100875780639ede42b914610251578063a521b03214610274578063d248d12414610287578063e78cea921461029a578063ee35f327146102ad578063f26a62c6146102c0578063f8c8765e146102d3578063fb7be0a1146102e657600080fd5b806314eab5e7146100e55780631b45c86a1461011557806323a9ef231461012a5780633504f1d71461015557806356e9df97146101685780635ef489e61461017b5780637fd07a9c1461018e5780638f1d3776146101ae575b600080fd5b6100f86100f33660046125b7565b6102f9565b6040516001600160401b0390911681526020015b60405180910390f35b61012861012336600461264a565b6105e9565b005b61013d61013836600461264a565b6106b9565b6040516001600160a01b03909116815260200161010c565b60025461013d906001600160a01b031681565b61012861017636600461264a565b6106dd565b6000546100f8906001600160401b031681565b6101a161019c36600461264a565b61084b565b60405161010c91906126a7565b61023e6101bc366004612719565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b60405161010c9796959493929190612732565b61026461025f36600461264a565b610924565b604051901515815260200161010c565b61012861028236600461278f565b61094b565b610128610295366004612833565b610dbd565b60045461013d906001600160a01b031681565b60035461013d906001600160a01b031681565b60055461013d906001600160a01b031681565b6101286102e13660046128c5565b6111eb565b6101286102f4366004612921565b61135b565b6002546000906001600160a01b0316331461034e5760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506103a461037f60208b018b6129c5565b61039f8a60005b6080020180360381019061039a9190612a84565b6119c6565b611a47565b816000815181106103b7576103b76129af565b60209081029190910101526103e68960016020020160208101906103db91906129c5565b61039f8a6001610386565b816001815181106103f9576103f96129af565b6020908102919091010152600080548190819061041e906001600160401b0316612b32565b82546001600160401b038083166101009490940a8481029102199091161790925590915061044e5761044e612b58565b6001600160401b0381166000908152600160205260408120600581018d905590610488610483368d90038d0160808e01612a84565b611b68565b9050600261049c60408e0160208f016129c5565b60038111156104ad576104ad61267d565b14806104db575060006104d06104cb368e90038e0160808f01612a84565b611b7d565b6001600160401b0316115b156104ee57806104ea81612b32565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a906105b8908e906080820190612bb5565b60405180910390a26105d68360008c6001600160401b031687611b8c565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff1660028111156106225761062261267d565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906106625760405162461bcd60e51b81526004016103459190612bd1565b5061066c81610924565b6106ab5760405162461bcd60e51b815260206004820152601060248201526f54494d454f55545f444541444c494e4560801b6044820152606401610345565b6106b6816000611c22565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6002546001600160a01b0316331461072a5760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b6044820152606401610345565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff1660028111156107635761076361267d565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906107a35760405162461bcd60e51b81526004016103459190612bd1565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f409161084091612c1f565b60405180910390a250565b610853612512565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff169081111561090a5761090a61267d565b600281111561091b5761091b61267d565b90525092915050565b6001600160401b038116600090815260016020526040812061094590611d50565b92915050565b6001600160401b038416600090815260016020526040812085918591610970846106b9565b6001600160a01b0316336001600160a01b0316146109a05760405162461bcd60e51b815260040161034590612c39565b6109a984610924565b156109c65760405162461bcd60e51b815260040161034590612c5e565b60008260028111156109da576109da61267d565b03610a475760006007820154600160401b900460ff166002811115610a0157610a0161267d565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610a415760405162461bcd60e51b81526004016103459190612bd1565b50610b04565b6001826002811115610a5b57610a5b61267d565b03610aa45760016007820154600160401b900460ff166002811115610a8257610a8261267d565b14610a9f5760405162461bcd60e51b815260040161034590612c85565b610b04565b6002826002811115610ab857610ab861267d565b03610afc5760026007820154600160401b900460ff166002811115610adf57610adf61267d565b14610a9f5760405162461bcd60e51b815260040161034590612cad565b610b04612b58565b610b5283356020850135610b1b6040870187612cd9565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611d6892505050565b816006015414610b745760405162461bcd60e51b815260040161034590612d29565b6002610b836040850185612cd9565b90501080610bae57506001610b9b6040850185612cd9565b610ba6929150612d4c565b836060013510155b15610bcb5760405162461bcd60e51b815260040161034590612d5f565b600080610bd789611d9f565b9150915060018111610c175760405162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b6044820152606401610345565b806028811115610c25575060285b610c30816001612d8a565b8814610c6d5760405162461bcd60e51b815260206004820152600c60248201526b57524f4e475f44454752454560a01b6044820152606401610345565b50610cb78989896000818110610c8557610c856129af565b602002919091013590508a8a610c9c600182612d4c565b818110610cab57610cab6129af565b90506020020135611e2f565b610cf68a83838b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611b8c92505050565b50600090505b6007820154600160401b900460ff166002811115610d1c57610d1c61267d565b03610d275750610db4565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154610d579042612d4c565b81602001818151610d689190612d4c565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b50505050505050565b6001600160401b038416600090815260016020526040902084908490600290610de5846106b9565b6001600160a01b0316336001600160a01b031614610e155760405162461bcd60e51b815260040161034590612c39565b610e1e84610924565b15610e3b5760405162461bcd60e51b815260040161034590612c5e565b6000826002811115610e4f57610e4f61267d565b03610ebc5760006007820154600160401b900460ff166002811115610e7657610e7661267d565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610eb65760405162461bcd60e51b81526004016103459190612bd1565b50610f79565b6001826002811115610ed057610ed061267d565b03610f195760016007820154600160401b900460ff166002811115610ef757610ef761267d565b14610f145760405162461bcd60e51b815260040161034590612c85565b610f79565b6002826002811115610f2d57610f2d61267d565b03610f715760026007820154600160401b900460ff166002811115610f5457610f5461267d565b14610f145760405162461bcd60e51b815260040161034590612cad565b610f79612b58565b610f9083356020850135610b1b6040870187612cd9565b816006015414610fb25760405162461bcd60e51b815260040161034590612d29565b6002610fc16040850185612cd9565b90501080610fec57506001610fd96040850185612cd9565b610fe4929150612d4c565b836060013510155b156110095760405162461bcd60e51b815260040161034590612d5f565b6001600160401b0388166000908152600160205260408120908061102c8a611d9f565b9092509050600181146110515760405162461bcd60e51b815260040161034590612d9d565b5060055460408051808201825260078501546001600160401b031681526004546001600160a01b0390811660208301526000931691635d3adcfb9190859061109b908f018f612cd9565b8f606001358181106110af576110af6129af565b905060200201358d8d6040518663ffffffff1660e01b81526004016110d8959493929190612dbf565b602060405180830381865afa1580156110f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111199190612e16565b905061112860408b018b612cd9565b61113760608d01356001612d8a565b818110611146576111466129af565b90506020020135810361118a5760405162461bcd60e51b815260206004820152600c60248201526b14d0535157d3d4d417d1539160a21b6044820152606401610345565b6040516001600160401b038c16907fc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e90600090a26111df8b6001600160401b0316600090815260016020526040812060060155565b5060009150610cfc9050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036112785760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610345565b6002546001600160a01b0316156112c05760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610345565b6001600160a01b03841661130b5760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b6044820152606401610345565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b038516600090815260016020819052604090912086918691611383846106b9565b6001600160a01b0316336001600160a01b0316146113b35760405162461bcd60e51b815260040161034590612c39565b6113bc84610924565b156113d95760405162461bcd60e51b815260040161034590612c5e565b60008260028111156113ed576113ed61267d565b0361145a5760006007820154600160401b900460ff1660028111156114145761141461267d565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906114545760405162461bcd60e51b81526004016103459190612bd1565b50611517565b600182600281111561146e5761146e61267d565b036114b75760016007820154600160401b900460ff1660028111156114955761149561267d565b146114b25760405162461bcd60e51b815260040161034590612c85565b611517565b60028260028111156114cb576114cb61267d565b0361150f5760026007820154600160401b900460ff1660028111156114f2576114f261267d565b146114b25760405162461bcd60e51b815260040161034590612cad565b611517612b58565b61152e83356020850135610b1b6040870187612cd9565b8160060154146115505760405162461bcd60e51b815260040161034590612d29565b600261155f6040850185612cd9565b9050108061158a575060016115776040850185612cd9565b611582929150612d4c565b836060013510155b156115a75760405162461bcd60e51b815260040161034590612d5f565b60018510156115ee5760405162461bcd60e51b815260206004820152601360248201527210d2105313115391d157d513d3d7d4d213d495606a1b6044820152606401610345565b650800000000008511156116395760405162461bcd60e51b81526020600482015260126024820152714348414c4c454e47455f544f4f5f4c4f4e4760701b6044820152606401610345565b61167b8861165b61164d60208b018b6129c5565b8960005b6020020135611a47565b61167661166e60408c0160208d016129c5565b8a6001611651565b611e2f565b6001600160401b0389166000908152600160205260408120908061169e8b611d9f565b91509150806001146116c25760405162461bcd60e51b815260040161034590612d9d565b60016116d160208c018c6129c5565b60038111156116e2576116e261267d565b1461179c576116f760408b0160208c016129c5565b60038111156117085761170861267d565b61171560208c018c6129c5565b60038111156117265761172661267d565b1480156117375750883560208a0135145b6117735760405162461bcd60e51b815260206004820152600d60248201526c48414c5445445f4348414e474560981b6044820152606401610345565b6117948c6001600160401b0316600090815260016020526040812060060155565b505050611902565b60026117ae60408c0160208d016129c5565b60038111156117bf576117bf61267d565b0361180757883560208a0135146118075760405162461bcd60e51b815260206004820152600c60248201526b4552524f525f4348414e474560a01b6044820152606401610345565b6040805160028082526060820183526000926020830190803683375050506005850154909150611839908b3590611f03565b8160008151811061184c5761184c6129af565b602090810291909101015261187a8b600160200201602081019061187091906129c5565b60208c013561208f565b8160018151811061188d5761188d6129af565b602090810291909101015260078401805460ff60401b1916600160411b1790556118ba8d60008b84611b8c565b8c6001600160401b03167f24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db846040516118f591815260200190565b60405180910390a2505050505b60006007820154600160401b900460ff1660028111156119245761192461267d565b0361192f57506119bc565b6040805180820190915281546001600160a01b0316815260018201546020820152600482015461195f9042612d4c565b816020018181516119709190612d4c565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b5050505050505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b60006001836003811115611a5d57611a5d61267d565b03611aa2576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b604051602081830303815290604052805190602001209050610945565b6002836003811115611ab657611ab661267d565b03611aeb5760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501611a85565b6003836003811115611aff57611aff61267d565b03611b2d5760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501611a85565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606401610345565b6020810151600090815b602002015192915050565b60208101516000906001611b72565b6001821015611b9d57611b9d612b58565b600281511015611baf57611baf612b58565b6000611bbc848484611d68565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890611c1390889088908890612e2f565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015611cf057600080fd5b505af1158015611d04573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051611d419190612c1f565b60405180910390a25050505050565b6001810154600090611d6183612135565b1192915050565b6000838383604051602001611d7f93929190612e84565b6040516020818303038152906040528051906020012090505b9392505050565b600080806001611db26040860186612cd9565b611dbd929150612d4c565b9050611dcd816020860135612edc565b9150611ddd606085013583612ef0565b611de8908535612d8a565b92506002611df96040860186612cd9565b611e04929150612d4c565b846060013503611e2957611e1c816020860135612f07565b611e269083612d8a565b91505b50915091565b81611e3d6040850185612cd9565b8560600135818110611e5157611e516129af565b9050602002013514611e935760405162461bcd60e51b815260206004820152600b60248201526a15d493d391d7d4d510549560aa1b6044820152606401610345565b80611ea16040850185612cd9565b611eb060608701356001612d8a565b818110611ebf57611ebf6129af565b9050602002013503611efe5760405162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b6044820152606401610345565b505050565b60408051600380825260808201909252600091829190816020015b6040805180820190915260008082526020820152815260200190600190039081611f1e575050604080518082018252600080825260209182018190528251808401909352600483529082015290915081600081518110611f8057611f806129af565b6020026020010181905250611f956000612147565b81600181518110611fa857611fa86129af565b6020026020010181905250611fbd6000612147565b81600281518110611fd057611fd06129af565b602090810291909101810191909152604080518083018252838152815180830190925280825260009282019290925261202060408051606080820183529181019182529081526000602082015290565b604080518082018252606080825260006020808401829052845161012081018652828152908101879052938401859052908301829052608083018a905260a0830181905260c0830181905260e08301526101008201889052906120828161217a565b9998505050505050505050565b600060018360038111156120a5576120a561267d565b036120bb5781604051602001611a859190612f1b565b60028360038111156120cf576120cf61267d565b036120f8576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611a85565b600383600381111561210c5761210c61267d565b03611b2d576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611a85565b60008160040154426109459190612d4c565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b600080825160038111156121905761219061267d565b03612245576121a2826020015161232f565b6121af836040015161232f565b6121bc84606001516123b4565b608085015160a086015160c087015160e0808901516101008a01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991831b821660b0870152821b811660b486015291901b1660b883015260bc82015260dc01611a2a565b60018251600381111561225a5761225a61267d565b03612274578160800151604051602001611a2a9190612f1b565b6002825160038111156122895761228961267d565b036122b2576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611a2a565b6003825160038111156122c7576122c761267d565b036122f0576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611a2a565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610345565b919050565b60208101518151515160005b818110156123ad57835161235890612353908361244d565b612485565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806123a590612f40565b91505061233b565b5050919050565b602081015160005b825151811015612447576123ec836000015182815181106123df576123df6129af565b60200260200101516124a2565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b6020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061243f90612f40565b9150506123bc565b50919050565b60408051808201909152600080825260208201528251805183908110612475576124756129af565b6020026020010151905092915050565b600081600001518260200151604051602001611a2a929190612f59565b60006124b18260000151612485565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401611a2a565b604080516101208101909152600060e0820181815261010083019190915281908152602001612551604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b806040810183101561094557600080fd5b80356001600160401b038116811461232a57600080fd5b6001600160a01b03811681146106b657600080fd5b600080600080600080600080610200898b0312156125d457600080fd5b883597506125e58a60208b0161257a565b965061016089018a8111156125f957600080fd5b60608a0196506126088161258b565b95505061018089013561261a816125a2565b93506101a089013561262b816125a2565b979a96995094979396929592945050506101c0820135916101e0013590565b60006020828403121561265c57600080fd5b611d988261258b565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b600381106126a3576126a361267d565b9052565b6000610120820190506126bb828451612665565b60208301516126cd6040840182612665565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c0830151612712610100840182612693565b5092915050565b60006020828403121561272b57600080fd5b5035919050565b6101208101612741828a612665565b61274e6040830189612665565b8660808301528560a08301528460c08301526001600160401b03841660e08301526105dd610100830184612693565b60006080828403121561244757600080fd5b600080600080606085870312156127a557600080fd5b6127ae8561258b565b935060208501356001600160401b03808211156127ca57600080fd5b6127d68883890161277d565b945060408701359150808211156127ec57600080fd5b818701915087601f83011261280057600080fd5b81358181111561280f57600080fd5b8860208260051b850101111561282457600080fd5b95989497505060200194505050565b6000806000806060858703121561284957600080fd5b6128528561258b565b935060208501356001600160401b038082111561286e57600080fd5b61287a8883890161277d565b9450604087013591508082111561289057600080fd5b818701915087601f8301126128a457600080fd5b8135818111156128b357600080fd5b88602082850101111561282457600080fd5b600080600080608085870312156128db57600080fd5b84356128e6816125a2565b935060208501356128f6816125a2565b92506040850135612906816125a2565b91506060850135612916816125a2565b939692955090935050565b600080600080600060e0868803121561293957600080fd5b6129428661258b565b945060208601356001600160401b0381111561295d57600080fd5b6129698882890161277d565b945050612979876040880161257a565b9250612988876080880161257a565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000602082840312156129d757600080fd5b813560048110611d9857600080fd5b604080519081016001600160401b0381118282101715612a0857612a08612999565b60405290565b600082601f830112612a1f57600080fd5b604051604081018181106001600160401b0382111715612a4157612a41612999565b8060405250806040840185811115612a5857600080fd5b845b81811015612a7957612a6b8161258b565b835260209283019201612a5a565b509195945050505050565b600060808284031215612a9657600080fd5b604051604081018181106001600160401b0382111715612ab857612ab8612999565b604052601f83018413612aca57600080fd5b612ad26129e6565b806040850186811115612ae457600080fd5b855b81811015612afe578035845260209384019301612ae6565b50818452612b0c8782612a0e565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808316818103612b4e57612b4e612b1c565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b6040818337604082016040820160005b6002811015612bae576001600160401b03612b988361258b565b1683526020928301929190910190600101612b7e565b5050505050565b6101008101612bc48285612b6e565b611d986080830184612b6e565b600060208083528351808285015260005b81811015612bfe57858101830151858201604001528201612be2565b506000604082860101526040601f19601f8301168501019250505092915050565b6020810160048310612c3357612c3361267d565b91905290565b6020808252600b908201526a21a420a62fa9a2a72222a960a91b604082015260600190565b6020808252600d908201526c4348414c5f444541444c494e4560981b604082015260600190565b6020808252600e908201526d4348414c5f4e4f545f424c4f434b60901b604082015260600190565b60208082526012908201527121a420a62fa727aa2fa2ac22a1aaaa24a7a760711b604082015260600190565b6000808335601e19843603018112612cf057600080fd5b8301803591506001600160401b03821115612d0a57600080fd5b6020019150600581901b3603821315612d2257600080fd5b9250929050565b6020808252600990820152684249535f535441544560b81b604082015260600190565b8181038181111561094557610945612b1c565b6020808252601190820152704241445f4348414c4c454e47455f504f5360781b604082015260600190565b8082018082111561094557610945612b1c565b602080825260089082015267544f4f5f4c4f4e4760c01b604082015260600190565b8551815260018060a01b03602087015116602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f19160101949350505050565b600060208284031215612e2857600080fd5b5051919050565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015612e7657845183529383019391830191600101612e5a565b509098975050505050505050565b83815260006020848184015260408301845182860160005b82811015612eb857815184529284019290840190600101612e9c565b509198975050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082612eeb57612eeb612ec6565b500490565b808202811582820484141761094557610945612b1c565b600082612f1657612f16612ec6565b500690565b7026b0b1b434b732903334b734b9b432b21d60791b8152601181019190915260310190565b600060018201612f5257612f52612b1c565b5060010190565b652b30b63ab29d60d11b8152600060078410612f7757612f7761267d565b5060f89290921b600683015260078201526027019056fea264697066735822122067a773d4d93ed175eca41913ffbf9c60d1086d428687f06f494b34ce6d944d8964736f6c63430008110033",
}

// OldChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OldChallengeManagerMetaData.ABI instead.
var OldChallengeManagerABI = OldChallengeManagerMetaData.ABI

// OldChallengeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OldChallengeManagerMetaData.Bin instead.
var OldChallengeManagerBin = OldChallengeManagerMetaData.Bin

// DeployOldChallengeManager deploys a new Ethereum contract, binding an instance of OldChallengeManager to it.
func DeployOldChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldChallengeManager, error) {
	parsed, err := OldChallengeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OldChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldChallengeManager{OldChallengeManagerCaller: OldChallengeManagerCaller{contract: contract}, OldChallengeManagerTransactor: OldChallengeManagerTransactor{contract: contract}, OldChallengeManagerFilterer: OldChallengeManagerFilterer{contract: contract}}, nil
}

// OldChallengeManager is an auto generated Go binding around an Ethereum contract.
type OldChallengeManager struct {
	OldChallengeManagerCaller     // Read-only binding to the contract
	OldChallengeManagerTransactor // Write-only binding to the contract
	OldChallengeManagerFilterer   // Log filterer for contract events
}

// OldChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldChallengeManagerSession struct {
	Contract     *OldChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OldChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldChallengeManagerCallerSession struct {
	Contract *OldChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OldChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldChallengeManagerTransactorSession struct {
	Contract     *OldChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OldChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldChallengeManagerRaw struct {
	Contract *OldChallengeManager // Generic contract binding to access the raw methods on
}

// OldChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldChallengeManagerCallerRaw struct {
	Contract *OldChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OldChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldChallengeManagerTransactorRaw struct {
	Contract *OldChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldChallengeManager creates a new instance of OldChallengeManager, bound to a specific deployed contract.
func NewOldChallengeManager(address common.Address, backend bind.ContractBackend) (*OldChallengeManager, error) {
	contract, err := bindOldChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManager{OldChallengeManagerCaller: OldChallengeManagerCaller{contract: contract}, OldChallengeManagerTransactor: OldChallengeManagerTransactor{contract: contract}, OldChallengeManagerFilterer: OldChallengeManagerFilterer{contract: contract}}, nil
}

// NewOldChallengeManagerCaller creates a new read-only instance of OldChallengeManager, bound to a specific deployed contract.
func NewOldChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*OldChallengeManagerCaller, error) {
	contract, err := bindOldChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerCaller{contract: contract}, nil
}

// NewOldChallengeManagerTransactor creates a new write-only instance of OldChallengeManager, bound to a specific deployed contract.
func NewOldChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OldChallengeManagerTransactor, error) {
	contract, err := bindOldChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerTransactor{contract: contract}, nil
}

// NewOldChallengeManagerFilterer creates a new log filterer instance of OldChallengeManager, bound to a specific deployed contract.
func NewOldChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OldChallengeManagerFilterer, error) {
	contract, err := bindOldChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerFilterer{contract: contract}, nil
}

// bindOldChallengeManager binds a generic wrapper to an already deployed contract.
func bindOldChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldChallengeManager *OldChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldChallengeManager.Contract.OldChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldChallengeManager *OldChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.OldChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldChallengeManager *OldChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.OldChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldChallengeManager *OldChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldChallengeManager *OldChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldChallengeManager *OldChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OldChallengeManager *OldChallengeManagerSession) Bridge() (common.Address, error) {
	return _OldChallengeManager.Contract.Bridge(&_OldChallengeManager.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCallerSession) Bridge() (common.Address, error) {
	return _OldChallengeManager.Contract.Bridge(&_OldChallengeManager.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_OldChallengeManager *OldChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (OldChallengeLibChallenge, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(OldChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(OldChallengeLibChallenge)).(*OldChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_OldChallengeManager *OldChallengeManagerSession) ChallengeInfo(challengeIndex uint64) (OldChallengeLibChallenge, error) {
	return _OldChallengeManager.Contract.ChallengeInfo(&_OldChallengeManager.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_OldChallengeManager *OldChallengeManagerCallerSession) ChallengeInfo(challengeIndex uint64) (OldChallengeLibChallenge, error) {
	return _OldChallengeManager.Contract.ChallengeInfo(&_OldChallengeManager.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_OldChallengeManager *OldChallengeManagerCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            OldChallengeLibParticipant
	Next               OldChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		Current            OldChallengeLibParticipant
		Next               OldChallengeLibParticipant
		LastMoveTimestamp  *big.Int
		WasmModuleRoot     [32]byte
		ChallengeStateHash [32]byte
		MaxInboxMessages   uint64
		Mode               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(OldChallengeLibParticipant)).(*OldChallengeLibParticipant)
	outstruct.Next = *abi.ConvertType(out[1], new(OldChallengeLibParticipant)).(*OldChallengeLibParticipant)
	outstruct.LastMoveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WasmModuleRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ChallengeStateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MaxInboxMessages = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Mode = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_OldChallengeManager *OldChallengeManagerSession) Challenges(arg0 *big.Int) (struct {
	Current            OldChallengeLibParticipant
	Next               OldChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _OldChallengeManager.Contract.Challenges(&_OldChallengeManager.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_OldChallengeManager *OldChallengeManagerCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            OldChallengeLibParticipant
	Next               OldChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _OldChallengeManager.Contract.Challenges(&_OldChallengeManager.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_OldChallengeManager *OldChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_OldChallengeManager *OldChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _OldChallengeManager.Contract.CurrentResponder(&_OldChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_OldChallengeManager *OldChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _OldChallengeManager.Contract.CurrentResponder(&_OldChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_OldChallengeManager *OldChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_OldChallengeManager *OldChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _OldChallengeManager.Contract.IsTimedOut(&_OldChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_OldChallengeManager *OldChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _OldChallengeManager.Contract.IsTimedOut(&_OldChallengeManager.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_OldChallengeManager *OldChallengeManagerSession) Osp() (common.Address, error) {
	return _OldChallengeManager.Contract.Osp(&_OldChallengeManager.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCallerSession) Osp() (common.Address, error) {
	return _OldChallengeManager.Contract.Osp(&_OldChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_OldChallengeManager *OldChallengeManagerSession) ResultReceiver() (common.Address, error) {
	return _OldChallengeManager.Contract.ResultReceiver(&_OldChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCallerSession) ResultReceiver() (common.Address, error) {
	return _OldChallengeManager.Contract.ResultReceiver(&_OldChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_OldChallengeManager *OldChallengeManagerSession) SequencerInbox() (common.Address, error) {
	return _OldChallengeManager.Contract.SequencerInbox(&_OldChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_OldChallengeManager *OldChallengeManagerCallerSession) SequencerInbox() (common.Address, error) {
	return _OldChallengeManager.Contract.SequencerInbox(&_OldChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_OldChallengeManager *OldChallengeManagerCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OldChallengeManager.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_OldChallengeManager *OldChallengeManagerSession) TotalChallengesCreated() (uint64, error) {
	return _OldChallengeManager.Contract.TotalChallengesCreated(&_OldChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_OldChallengeManager *OldChallengeManagerCallerSession) TotalChallengesCreated() (uint64, error) {
	return _OldChallengeManager.Contract.TotalChallengesCreated(&_OldChallengeManager.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection OldChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_OldChallengeManager *OldChallengeManagerSession) BisectExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.BisectExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) BisectExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.BisectExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection OldChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_OldChallengeManager *OldChallengeManagerSession) ChallengeExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.ChallengeExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) ChallengeExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.ChallengeExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.ClearChallenge(&_OldChallengeManager.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.ClearChallenge(&_OldChallengeManager.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_OldChallengeManager *OldChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_OldChallengeManager *OldChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.CreateChallenge(&_OldChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_OldChallengeManager *OldChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.CreateChallenge(&_OldChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_OldChallengeManager *OldChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.Initialize(&_OldChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.Initialize(&_OldChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection OldChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_OldChallengeManager *OldChallengeManagerSession) OneStepProveExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.OneStepProveExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) OneStepProveExecution(challengeIndex uint64, selection OldChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.OneStepProveExecution(&_OldChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.Timeout(&_OldChallengeManager.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_OldChallengeManager *OldChallengeManagerTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _OldChallengeManager.Contract.Timeout(&_OldChallengeManager.TransactOpts, challengeIndex)
}

// OldChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the OldChallengeManager contract.
type OldChallengeManagerBisectedIterator struct {
	Event *OldChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *OldChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldChallengeManagerBisected)
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
		it.Event = new(OldChallengeManagerBisected)
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
func (it *OldChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldChallengeManagerBisected represents a Bisected event raised by the OldChallengeManager contract.
type OldChallengeManagerBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_OldChallengeManager *OldChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*OldChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _OldChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerBisectedIterator{contract: _OldChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_OldChallengeManager *OldChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *OldChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _OldChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldChallengeManagerBisected)
				if err := _OldChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_OldChallengeManager *OldChallengeManagerFilterer) ParseBisected(log types.Log) (*OldChallengeManagerBisected, error) {
	event := new(OldChallengeManagerBisected)
	if err := _OldChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OldChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the OldChallengeManager contract.
type OldChallengeManagerChallengeEndedIterator struct {
	Event *OldChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *OldChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldChallengeManagerChallengeEnded)
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
		it.Event = new(OldChallengeManagerChallengeEnded)
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
func (it *OldChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the OldChallengeManager contract.
type OldChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_OldChallengeManager *OldChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*OldChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerChallengeEndedIterator{contract: _OldChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_OldChallengeManager *OldChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *OldChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldChallengeManagerChallengeEnded)
				if err := _OldChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_OldChallengeManager *OldChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*OldChallengeManagerChallengeEnded, error) {
	event := new(OldChallengeManagerChallengeEnded)
	if err := _OldChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OldChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the OldChallengeManager contract.
type OldChallengeManagerExecutionChallengeBegunIterator struct {
	Event *OldChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *OldChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(OldChallengeManagerExecutionChallengeBegun)
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
func (it *OldChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the OldChallengeManager contract.
type OldChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_OldChallengeManager *OldChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*OldChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerExecutionChallengeBegunIterator{contract: _OldChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_OldChallengeManager *OldChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *OldChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldChallengeManagerExecutionChallengeBegun)
				if err := _OldChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_OldChallengeManager *OldChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*OldChallengeManagerExecutionChallengeBegun, error) {
	event := new(OldChallengeManagerExecutionChallengeBegun)
	if err := _OldChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OldChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the OldChallengeManager contract.
type OldChallengeManagerInitiatedChallengeIterator struct {
	Event *OldChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *OldChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldChallengeManagerInitiatedChallenge)
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
		it.Event = new(OldChallengeManagerInitiatedChallenge)
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
func (it *OldChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the OldChallengeManager contract.
type OldChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_OldChallengeManager *OldChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*OldChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerInitiatedChallengeIterator{contract: _OldChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_OldChallengeManager *OldChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *OldChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldChallengeManagerInitiatedChallenge)
				if err := _OldChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_OldChallengeManager *OldChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*OldChallengeManagerInitiatedChallenge, error) {
	event := new(OldChallengeManagerInitiatedChallenge)
	if err := _OldChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OldChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the OldChallengeManager contract.
type OldChallengeManagerOneStepProofCompletedIterator struct {
	Event *OldChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *OldChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldChallengeManagerOneStepProofCompleted)
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
		it.Event = new(OldChallengeManagerOneStepProofCompleted)
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
func (it *OldChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the OldChallengeManager contract.
type OldChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_OldChallengeManager *OldChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*OldChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &OldChallengeManagerOneStepProofCompletedIterator{contract: _OldChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_OldChallengeManager *OldChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *OldChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _OldChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldChallengeManagerOneStepProofCompleted)
				if err := _OldChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_OldChallengeManager *OldChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*OldChallengeManagerOneStepProofCompleted, error) {
	event := new(OldChallengeManagerOneStepProofCompleted)
	if err := _OldChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
