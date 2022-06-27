// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ServiceABI is the input ABI used to generate the binding from.
const ServiceABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"NewTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenUnlock\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_transaction\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"deployContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isProccessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_transaction\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeTokenContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signaturesForTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_transaction\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Service is an auto generated Go binding around an Ethereum contract.
type Service struct {
	ServiceCaller     // Read-only binding to the contract
	ServiceTransactor // Write-only binding to the contract
	ServiceFilterer   // Log filterer for contract events
}

// ServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceSession struct {
	Contract     *Service          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceCallerSession struct {
	Contract *ServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceTransactorSession struct {
	Contract     *ServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceRaw struct {
	Contract *Service // Generic contract binding to access the raw methods on
}

// ServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceCallerRaw struct {
	Contract *ServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceTransactorRaw struct {
	Contract *ServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewService creates a new instance of Service, bound to a specific deployed contract.
func NewService(address common.Address, backend bind.ContractBackend) (*Service, error) {
	contract, err := bindService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Service{ServiceCaller: ServiceCaller{contract: contract}, ServiceTransactor: ServiceTransactor{contract: contract}, ServiceFilterer: ServiceFilterer{contract: contract}}, nil
}

// NewServiceCaller creates a new read-only instance of Service, bound to a specific deployed contract.
func NewServiceCaller(address common.Address, caller bind.ContractCaller) (*ServiceCaller, error) {
	contract, err := bindService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceCaller{contract: contract}, nil
}

// NewServiceTransactor creates a new write-only instance of Service, bound to a specific deployed contract.
func NewServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceTransactor, error) {
	contract, err := bindService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceTransactor{contract: contract}, nil
}

// NewServiceFilterer creates a new log filterer instance of Service, bound to a specific deployed contract.
func NewServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceFilterer, error) {
	contract, err := bindService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceFilterer{contract: contract}, nil
}

// bindService binds a generic wrapper to an already deployed contract.
func bindService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Service.Contract.ServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Service.Contract.ServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Service.Contract.ServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Service.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Service.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Service.Contract.contract.Transact(opts, method, params...)
}

// IsProccessed is a free data retrieval call binding the contract method 0xe2ed78bb.
//
// Solidity: function isProccessed(string ) view returns(bool)
func (_Service *ServiceCaller) IsProccessed(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Service.contract.Call(opts, &out, "isProccessed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProccessed is a free data retrieval call binding the contract method 0xe2ed78bb.
//
// Solidity: function isProccessed(string ) view returns(bool)
func (_Service *ServiceSession) IsProccessed(arg0 string) (bool, error) {
	return _Service.Contract.IsProccessed(&_Service.CallOpts, arg0)
}

// IsProccessed is a free data retrieval call binding the contract method 0xe2ed78bb.
//
// Solidity: function isProccessed(string ) view returns(bool)
func (_Service *ServiceCallerSession) IsProccessed(arg0 string) (bool, error) {
	return _Service.Contract.IsProccessed(&_Service.CallOpts, arg0)
}

// NativeTokenContracts is a free data retrieval call binding the contract method 0xe92ea15e.
//
// Solidity: function nativeTokenContracts(address ) view returns(address)
func (_Service *ServiceCaller) NativeTokenContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Service.contract.Call(opts, &out, "nativeTokenContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenContracts is a free data retrieval call binding the contract method 0xe92ea15e.
//
// Solidity: function nativeTokenContracts(address ) view returns(address)
func (_Service *ServiceSession) NativeTokenContracts(arg0 common.Address) (common.Address, error) {
	return _Service.Contract.NativeTokenContracts(&_Service.CallOpts, arg0)
}

// NativeTokenContracts is a free data retrieval call binding the contract method 0xe92ea15e.
//
// Solidity: function nativeTokenContracts(address ) view returns(address)
func (_Service *ServiceCallerSession) NativeTokenContracts(arg0 common.Address) (common.Address, error) {
	return _Service.Contract.NativeTokenContracts(&_Service.CallOpts, arg0)
}

// SignaturesForTransaction is a free data retrieval call binding the contract method 0x5e9c1863.
//
// Solidity: function signaturesForTransaction(string , uint256 ) view returns(address)
func (_Service *ServiceCaller) SignaturesForTransaction(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Service.contract.Call(opts, &out, "signaturesForTransaction", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignaturesForTransaction is a free data retrieval call binding the contract method 0x5e9c1863.
//
// Solidity: function signaturesForTransaction(string , uint256 ) view returns(address)
func (_Service *ServiceSession) SignaturesForTransaction(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _Service.Contract.SignaturesForTransaction(&_Service.CallOpts, arg0, arg1)
}

// SignaturesForTransaction is a free data retrieval call binding the contract method 0x5e9c1863.
//
// Solidity: function signaturesForTransaction(string , uint256 ) view returns(address)
func (_Service *ServiceCallerSession) SignaturesForTransaction(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _Service.Contract.SignaturesForTransaction(&_Service.CallOpts, arg0, arg1)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(address)
func (_Service *ServiceCaller) WrappedTokenContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Service.contract.Call(opts, &out, "wrappedTokenContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(address)
func (_Service *ServiceSession) WrappedTokenContracts(arg0 common.Address) (common.Address, error) {
	return _Service.Contract.WrappedTokenContracts(&_Service.CallOpts, arg0)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(address)
func (_Service *ServiceCallerSession) WrappedTokenContracts(arg0 common.Address) (common.Address, error) {
	return _Service.Contract.WrappedTokenContracts(&_Service.CallOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0xa0510a26.
//
// Solidity: function burn(address _owner, address _token, uint256 _amount, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceTransactor) Burn(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "burn", _owner, _token, _amount, _v, _r, _s)
}

// Burn is a paid mutator transaction binding the contract method 0xa0510a26.
//
// Solidity: function burn(address _owner, address _token, uint256 _amount, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceSession) Burn(_owner common.Address, _token common.Address, _amount *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.Contract.Burn(&_Service.TransactOpts, _owner, _token, _amount, _v, _r, _s)
}

// Burn is a paid mutator transaction binding the contract method 0xa0510a26.
//
// Solidity: function burn(address _owner, address _token, uint256 _amount, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceTransactorSession) Burn(_owner common.Address, _token common.Address, _amount *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.Contract.Burn(&_Service.TransactOpts, _owner, _token, _amount, _v, _r, _s)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9c5e3bdf.
//
// Solidity: function claimTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "claimTokens", _token, _amount, _transaction, _v, _r, _s)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9c5e3bdf.
//
// Solidity: function claimTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceSession) ClaimTokens(_token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.ClaimTokens(&_Service.TransactOpts, _token, _amount, _transaction, _v, _r, _s)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9c5e3bdf.
//
// Solidity: function claimTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactorSession) ClaimTokens(_token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.ClaimTokens(&_Service.TransactOpts, _token, _amount, _transaction, _v, _r, _s)
}

// DeployContract is a paid mutator transaction binding the contract method 0xf5fe572c.
//
// Solidity: function deployContract(address _nativeTokenAddress, string _name, string _symbol) returns()
func (_Service *ServiceTransactor) DeployContract(opts *bind.TransactOpts, _nativeTokenAddress common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "deployContract", _nativeTokenAddress, _name, _symbol)
}

// DeployContract is a paid mutator transaction binding the contract method 0xf5fe572c.
//
// Solidity: function deployContract(address _nativeTokenAddress, string _name, string _symbol) returns()
func (_Service *ServiceSession) DeployContract(_nativeTokenAddress common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Service.Contract.DeployContract(&_Service.TransactOpts, _nativeTokenAddress, _name, _symbol)
}

// DeployContract is a paid mutator transaction binding the contract method 0xf5fe572c.
//
// Solidity: function deployContract(address _nativeTokenAddress, string _name, string _symbol) returns()
func (_Service *ServiceTransactorSession) DeployContract(_nativeTokenAddress common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Service.Contract.DeployContract(&_Service.TransactOpts, _nativeTokenAddress, _name, _symbol)
}

// Lock is a paid mutator transaction binding the contract method 0xb058c837.
//
// Solidity: function lock(address _owner, address _token, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceTransactor) Lock(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "lock", _owner, _token, _amount, _deadline, _v, _r, _s)
}

// Lock is a paid mutator transaction binding the contract method 0xb058c837.
//
// Solidity: function lock(address _owner, address _token, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceSession) Lock(_owner common.Address, _token common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.Contract.Lock(&_Service.TransactOpts, _owner, _token, _amount, _deadline, _v, _r, _s)
}

// Lock is a paid mutator transaction binding the contract method 0xb058c837.
//
// Solidity: function lock(address _owner, address _token, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) payable returns()
func (_Service *ServiceTransactorSession) Lock(_owner common.Address, _token common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Service.Contract.Lock(&_Service.TransactOpts, _owner, _token, _amount, _deadline, _v, _r, _s)
}

// MintTokens is a paid mutator transaction binding the contract method 0xdf7af17b.
//
// Solidity: function mintTokens(address _nativeTokenAddress, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactor) MintTokens(opts *bind.TransactOpts, _nativeTokenAddress common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "mintTokens", _nativeTokenAddress, _amount, _transaction, _v, _r, _s)
}

// MintTokens is a paid mutator transaction binding the contract method 0xdf7af17b.
//
// Solidity: function mintTokens(address _nativeTokenAddress, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceSession) MintTokens(_nativeTokenAddress common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.MintTokens(&_Service.TransactOpts, _nativeTokenAddress, _amount, _transaction, _v, _r, _s)
}

// MintTokens is a paid mutator transaction binding the contract method 0xdf7af17b.
//
// Solidity: function mintTokens(address _nativeTokenAddress, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactorSession) MintTokens(_nativeTokenAddress common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.MintTokens(&_Service.TransactOpts, _nativeTokenAddress, _amount, _transaction, _v, _r, _s)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xf7aaab22.
//
// Solidity: function unlockTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactor) UnlockTokens(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.contract.Transact(opts, "unlockTokens", _token, _amount, _transaction, _v, _r, _s)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xf7aaab22.
//
// Solidity: function unlockTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceSession) UnlockTokens(_token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.UnlockTokens(&_Service.TransactOpts, _token, _amount, _transaction, _v, _r, _s)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xf7aaab22.
//
// Solidity: function unlockTokens(address _token, uint256 _amount, string _transaction, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Service *ServiceTransactorSession) UnlockTokens(_token common.Address, _amount *big.Int, _transaction string, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Service.Contract.UnlockTokens(&_Service.TransactOpts, _token, _amount, _transaction, _v, _r, _s)
}

// ServiceNewTokenDeployedIterator is returned from FilterNewTokenDeployed and is used to iterate over the raw logs and unpacked data for NewTokenDeployed events raised by the Service contract.
type ServiceNewTokenDeployedIterator struct {
	Event *ServiceNewTokenDeployed // Event containing the contract specifics and raw log

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
func (it *ServiceNewTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceNewTokenDeployed)
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
		it.Event = new(ServiceNewTokenDeployed)
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
func (it *ServiceNewTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceNewTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceNewTokenDeployed represents a NewTokenDeployed event raised by the Service contract.
type ServiceNewTokenDeployed struct {
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewTokenDeployed is a free log retrieval operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address indexed tokenContract)
func (_Service *ServiceFilterer) FilterNewTokenDeployed(opts *bind.FilterOpts, tokenContract []common.Address) (*ServiceNewTokenDeployedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Service.contract.FilterLogs(opts, "NewTokenDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &ServiceNewTokenDeployedIterator{contract: _Service.contract, event: "NewTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchNewTokenDeployed is a free log subscription operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address indexed tokenContract)
func (_Service *ServiceFilterer) WatchNewTokenDeployed(opts *bind.WatchOpts, sink chan<- *ServiceNewTokenDeployed, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Service.contract.WatchLogs(opts, "NewTokenDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceNewTokenDeployed)
				if err := _Service.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
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

// ParseNewTokenDeployed is a log parse operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address indexed tokenContract)
func (_Service *ServiceFilterer) ParseNewTokenDeployed(log types.Log) (*ServiceNewTokenDeployed, error) {
	event := new(ServiceNewTokenDeployed)
	if err := _Service.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceTokenLockIterator is returned from FilterTokenLock and is used to iterate over the raw logs and unpacked data for TokenLock events raised by the Service contract.
type ServiceTokenLockIterator struct {
	Event *ServiceTokenLock // Event containing the contract specifics and raw log

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
func (it *ServiceTokenLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceTokenLock)
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
		it.Event = new(ServiceTokenLock)
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
func (it *ServiceTokenLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceTokenLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceTokenLock represents a TokenLock event raised by the Service contract.
type ServiceTokenLock struct {
	From               common.Address
	SourceTokenAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenLock is a free log retrieval operation binding the contract event 0x2cae3a6b058bfe7efa080e1ee24ad4df747ad82e5c9113ae9018018531aec2b3.
//
// Solidity: event TokenLock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) FilterTokenLock(opts *bind.FilterOpts, from []common.Address, sourceTokenAddress []common.Address) (*ServiceTokenLockIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var sourceTokenAddressRule []interface{}
	for _, sourceTokenAddressItem := range sourceTokenAddress {
		sourceTokenAddressRule = append(sourceTokenAddressRule, sourceTokenAddressItem)
	}

	logs, sub, err := _Service.contract.FilterLogs(opts, "TokenLock", fromRule, sourceTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ServiceTokenLockIterator{contract: _Service.contract, event: "TokenLock", logs: logs, sub: sub}, nil
}

// WatchTokenLock is a free log subscription operation binding the contract event 0x2cae3a6b058bfe7efa080e1ee24ad4df747ad82e5c9113ae9018018531aec2b3.
//
// Solidity: event TokenLock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) WatchTokenLock(opts *bind.WatchOpts, sink chan<- *ServiceTokenLock, from []common.Address, sourceTokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var sourceTokenAddressRule []interface{}
	for _, sourceTokenAddressItem := range sourceTokenAddress {
		sourceTokenAddressRule = append(sourceTokenAddressRule, sourceTokenAddressItem)
	}

	logs, sub, err := _Service.contract.WatchLogs(opts, "TokenLock", fromRule, sourceTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceTokenLock)
				if err := _Service.contract.UnpackLog(event, "TokenLock", log); err != nil {
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

// ParseTokenLock is a log parse operation binding the contract event 0x2cae3a6b058bfe7efa080e1ee24ad4df747ad82e5c9113ae9018018531aec2b3.
//
// Solidity: event TokenLock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) ParseTokenLock(log types.Log) (*ServiceTokenLock, error) {
	event := new(ServiceTokenLock)
	if err := _Service.contract.UnpackLog(event, "TokenLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceTokenUnlockIterator is returned from FilterTokenUnlock and is used to iterate over the raw logs and unpacked data for TokenUnlock events raised by the Service contract.
type ServiceTokenUnlockIterator struct {
	Event *ServiceTokenUnlock // Event containing the contract specifics and raw log

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
func (it *ServiceTokenUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceTokenUnlock)
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
		it.Event = new(ServiceTokenUnlock)
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
func (it *ServiceTokenUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceTokenUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceTokenUnlock represents a TokenUnlock event raised by the Service contract.
type ServiceTokenUnlock struct {
	From               common.Address
	SourceTokenAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenUnlock is a free log retrieval operation binding the contract event 0x1f473bec585a8dd5a8df4e222f3139f825cacc28a63d17dc8a81ffdf1c48bdc7.
//
// Solidity: event TokenUnlock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) FilterTokenUnlock(opts *bind.FilterOpts, from []common.Address, sourceTokenAddress []common.Address) (*ServiceTokenUnlockIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var sourceTokenAddressRule []interface{}
	for _, sourceTokenAddressItem := range sourceTokenAddress {
		sourceTokenAddressRule = append(sourceTokenAddressRule, sourceTokenAddressItem)
	}

	logs, sub, err := _Service.contract.FilterLogs(opts, "TokenUnlock", fromRule, sourceTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ServiceTokenUnlockIterator{contract: _Service.contract, event: "TokenUnlock", logs: logs, sub: sub}, nil
}

// WatchTokenUnlock is a free log subscription operation binding the contract event 0x1f473bec585a8dd5a8df4e222f3139f825cacc28a63d17dc8a81ffdf1c48bdc7.
//
// Solidity: event TokenUnlock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) WatchTokenUnlock(opts *bind.WatchOpts, sink chan<- *ServiceTokenUnlock, from []common.Address, sourceTokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var sourceTokenAddressRule []interface{}
	for _, sourceTokenAddressItem := range sourceTokenAddress {
		sourceTokenAddressRule = append(sourceTokenAddressRule, sourceTokenAddressItem)
	}

	logs, sub, err := _Service.contract.WatchLogs(opts, "TokenUnlock", fromRule, sourceTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceTokenUnlock)
				if err := _Service.contract.UnpackLog(event, "TokenUnlock", log); err != nil {
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

// ParseTokenUnlock is a log parse operation binding the contract event 0x1f473bec585a8dd5a8df4e222f3139f825cacc28a63d17dc8a81ffdf1c48bdc7.
//
// Solidity: event TokenUnlock(address indexed from, address indexed sourceTokenAddress, uint256 amount)
func (_Service *ServiceFilterer) ParseTokenUnlock(log types.Log) (*ServiceTokenUnlock, error) {
	event := new(ServiceTokenUnlock)
	if err := _Service.contract.UnpackLog(event, "TokenUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
