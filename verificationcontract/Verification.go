// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verification

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
	_ = abi.ConvertType
)

// VerificationPost is an auto generated low-level Go binding around an user-defined struct.
type VerificationPost struct {
	PostNumber *big.Int
	IpfsHash   string
	QrCodeSvg  string
}

// VerificationMetaData contains all meta data concerning the Verification contract.
var VerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllUsersPost\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"postNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qrCodeSvg\",\"type\":\"string\"}],\"internalType\":\"structVerification.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"postNum\",\"type\":\"uint256\"}],\"name\":\"getSpecificPost\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"postNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qrCodeSvg\",\"type\":\"string\"}],\"internalType\":\"structVerification.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_qrcodeSvg\",\"type\":\"string\"}],\"name\":\"makePost\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"postNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qrCodeSvg\",\"type\":\"string\"}],\"internalType\":\"structVerification.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"postNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qrCodeSvg\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationMetaData.ABI instead.
var VerificationABI = VerificationMetaData.ABI

// Verification is an auto generated Go binding around an Ethereum contract.
type Verification struct {
	VerificationCaller     // Read-only binding to the contract
	VerificationTransactor // Write-only binding to the contract
	VerificationFilterer   // Log filterer for contract events
}

// VerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationSession struct {
	Contract     *Verification     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationCallerSession struct {
	Contract *VerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationTransactorSession struct {
	Contract     *VerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationRaw struct {
	Contract *Verification // Generic contract binding to access the raw methods on
}

// VerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationCallerRaw struct {
	Contract *VerificationCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationTransactorRaw struct {
	Contract *VerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerification creates a new instance of Verification, bound to a specific deployed contract.
func NewVerification(address common.Address, backend bind.ContractBackend) (*Verification, error) {
	contract, err := bindVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verification{VerificationCaller: VerificationCaller{contract: contract}, VerificationTransactor: VerificationTransactor{contract: contract}, VerificationFilterer: VerificationFilterer{contract: contract}}, nil
}

// NewVerificationCaller creates a new read-only instance of Verification, bound to a specific deployed contract.
func NewVerificationCaller(address common.Address, caller bind.ContractCaller) (*VerificationCaller, error) {
	contract, err := bindVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationCaller{contract: contract}, nil
}

// NewVerificationTransactor creates a new write-only instance of Verification, bound to a specific deployed contract.
func NewVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationTransactor, error) {
	contract, err := bindVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationTransactor{contract: contract}, nil
}

// NewVerificationFilterer creates a new log filterer instance of Verification, bound to a specific deployed contract.
func NewVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationFilterer, error) {
	contract, err := bindVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationFilterer{contract: contract}, nil
}

// bindVerification binds a generic wrapper to an already deployed contract.
func bindVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verification *VerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verification.Contract.VerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verification *VerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verification.Contract.VerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verification *VerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verification.Contract.VerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verification *VerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verification *VerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verification *VerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verification.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Verification *VerificationCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Verification *VerificationSession) Fee() (*big.Int, error) {
	return _Verification.Contract.Fee(&_Verification.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Verification *VerificationCallerSession) Fee() (*big.Int, error) {
	return _Verification.Contract.Fee(&_Verification.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Verification *VerificationCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Verification *VerificationSession) GetAdmin() (common.Address, error) {
	return _Verification.Contract.GetAdmin(&_Verification.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Verification *VerificationCallerSession) GetAdmin() (common.Address, error) {
	return _Verification.Contract.GetAdmin(&_Verification.CallOpts)
}

// GetAllUsersPost is a free data retrieval call binding the contract method 0xfd922863.
//
// Solidity: function getAllUsersPost(address user) view returns((uint256,string,string)[])
func (_Verification *VerificationCaller) GetAllUsersPost(opts *bind.CallOpts, user common.Address) ([]VerificationPost, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "getAllUsersPost", user)

	if err != nil {
		return *new([]VerificationPost), err
	}

	out0 := *abi.ConvertType(out[0], new([]VerificationPost)).(*[]VerificationPost)

	return out0, err

}

// GetAllUsersPost is a free data retrieval call binding the contract method 0xfd922863.
//
// Solidity: function getAllUsersPost(address user) view returns((uint256,string,string)[])
func (_Verification *VerificationSession) GetAllUsersPost(user common.Address) ([]VerificationPost, error) {
	return _Verification.Contract.GetAllUsersPost(&_Verification.CallOpts, user)
}

// GetAllUsersPost is a free data retrieval call binding the contract method 0xfd922863.
//
// Solidity: function getAllUsersPost(address user) view returns((uint256,string,string)[])
func (_Verification *VerificationCallerSession) GetAllUsersPost(user common.Address) ([]VerificationPost, error) {
	return _Verification.Contract.GetAllUsersPost(&_Verification.CallOpts, user)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Verification *VerificationCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Verification *VerificationSession) GetFee() (*big.Int, error) {
	return _Verification.Contract.GetFee(&_Verification.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Verification *VerificationCallerSession) GetFee() (*big.Int, error) {
	return _Verification.Contract.GetFee(&_Verification.CallOpts)
}

// GetSpecificPost is a free data retrieval call binding the contract method 0x9519a48d.
//
// Solidity: function getSpecificPost(address user, uint256 postNum) view returns((uint256,string,string))
func (_Verification *VerificationCaller) GetSpecificPost(opts *bind.CallOpts, user common.Address, postNum *big.Int) (VerificationPost, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "getSpecificPost", user, postNum)

	if err != nil {
		return *new(VerificationPost), err
	}

	out0 := *abi.ConvertType(out[0], new(VerificationPost)).(*VerificationPost)

	return out0, err

}

// GetSpecificPost is a free data retrieval call binding the contract method 0x9519a48d.
//
// Solidity: function getSpecificPost(address user, uint256 postNum) view returns((uint256,string,string))
func (_Verification *VerificationSession) GetSpecificPost(user common.Address, postNum *big.Int) (VerificationPost, error) {
	return _Verification.Contract.GetSpecificPost(&_Verification.CallOpts, user, postNum)
}

// GetSpecificPost is a free data retrieval call binding the contract method 0x9519a48d.
//
// Solidity: function getSpecificPost(address user, uint256 postNum) view returns((uint256,string,string))
func (_Verification *VerificationCallerSession) GetSpecificPost(user common.Address, postNum *big.Int) (VerificationPost, error) {
	return _Verification.Contract.GetSpecificPost(&_Verification.CallOpts, user, postNum)
}

// UsersPosts is a free data retrieval call binding the contract method 0x4be751a7.
//
// Solidity: function usersPosts(address , uint256 ) view returns(uint256 postNumber, string ipfsHash, string qrCodeSvg)
func (_Verification *VerificationCaller) UsersPosts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	PostNumber *big.Int
	IpfsHash   string
	QrCodeSvg  string
}, error) {
	var out []interface{}
	err := _Verification.contract.Call(opts, &out, "usersPosts", arg0, arg1)

	outstruct := new(struct {
		PostNumber *big.Int
		IpfsHash   string
		QrCodeSvg  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PostNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IpfsHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.QrCodeSvg = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// UsersPosts is a free data retrieval call binding the contract method 0x4be751a7.
//
// Solidity: function usersPosts(address , uint256 ) view returns(uint256 postNumber, string ipfsHash, string qrCodeSvg)
func (_Verification *VerificationSession) UsersPosts(arg0 common.Address, arg1 *big.Int) (struct {
	PostNumber *big.Int
	IpfsHash   string
	QrCodeSvg  string
}, error) {
	return _Verification.Contract.UsersPosts(&_Verification.CallOpts, arg0, arg1)
}

// UsersPosts is a free data retrieval call binding the contract method 0x4be751a7.
//
// Solidity: function usersPosts(address , uint256 ) view returns(uint256 postNumber, string ipfsHash, string qrCodeSvg)
func (_Verification *VerificationCallerSession) UsersPosts(arg0 common.Address, arg1 *big.Int) (struct {
	PostNumber *big.Int
	IpfsHash   string
	QrCodeSvg  string
}, error) {
	return _Verification.Contract.UsersPosts(&_Verification.CallOpts, arg0, arg1)
}

// MakePost is a paid mutator transaction binding the contract method 0x3822ba8d.
//
// Solidity: function makePost(string _ipfsHash, string _qrcodeSvg) payable returns((uint256,string,string))
func (_Verification *VerificationTransactor) MakePost(opts *bind.TransactOpts, _ipfsHash string, _qrcodeSvg string) (*types.Transaction, error) {
	return _Verification.contract.Transact(opts, "makePost", _ipfsHash, _qrcodeSvg)
}

// MakePost is a paid mutator transaction binding the contract method 0x3822ba8d.
//
// Solidity: function makePost(string _ipfsHash, string _qrcodeSvg) payable returns((uint256,string,string))
func (_Verification *VerificationSession) MakePost(_ipfsHash string, _qrcodeSvg string) (*types.Transaction, error) {
	return _Verification.Contract.MakePost(&_Verification.TransactOpts, _ipfsHash, _qrcodeSvg)
}

// MakePost is a paid mutator transaction binding the contract method 0x3822ba8d.
//
// Solidity: function makePost(string _ipfsHash, string _qrcodeSvg) payable returns((uint256,string,string))
func (_Verification *VerificationTransactorSession) MakePost(_ipfsHash string, _qrcodeSvg string) (*types.Transaction, error) {
	return _Verification.Contract.MakePost(&_Verification.TransactOpts, _ipfsHash, _qrcodeSvg)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns(address)
func (_Verification *VerificationTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Verification.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns(address)
func (_Verification *VerificationSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Verification.Contract.SetAdmin(&_Verification.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns(address)
func (_Verification *VerificationTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Verification.Contract.SetAdmin(&_Verification.TransactOpts, newAdmin)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Verification *VerificationTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Verification.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Verification *VerificationSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Verification.Contract.SetFee(&_Verification.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Verification *VerificationTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Verification.Contract.SetFee(&_Verification.TransactOpts, newFee)
}
