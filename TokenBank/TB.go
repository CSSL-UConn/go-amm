// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenBank

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

// TokenBankpositionStuct is an auto generated low-level Go binding around an user-defined struct.
type TokenBankpositionStuct struct {
	AmountTokenA *big.Int
	AmountTokenB *big.Int
	LowerBound   int32
	UpperBound   int32
	FeesEarnedA  *big.Int
	FeesEarnedB  *big.Int
}

// TokenBanksyncStruct is an auto generated low-level Go binding around an user-defined struct.
type TokenBanksyncStruct struct {
	TxTypeId      bool
	SidechainAddr *big.Int
	PositionId    string
	AmountTokenA  *big.Int
	AmountTokenB  *big.Int
	LowerBound    int32
	UpperBound    int32
	FeesEarnedA   *big.Int
	FeesEarnedB   *big.Int
}

// TokenBankMetaData contains all meta data concerning the TokenBank contract.
var TokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_amount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"_type\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sidechainAddr\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposit_map\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sepoliaAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBval\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_addr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"tokenAval\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"tokenBval\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"}],\"name\":\"getPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"amount_tokenA\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount_tokenB\",\"type\":\"int256\"},{\"internalType\":\"int32\",\"name\":\"lower_bound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"upper_bound\",\"type\":\"int32\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_A\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_B\",\"type\":\"int256\"}],\"internalType\":\"structTokenBank.position_stuct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"tx_type_id\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"sidechainAddr\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"position_id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"amount_tokenA\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount_tokenB\",\"type\":\"int256\"},{\"internalType\":\"int32\",\"name\":\"lower_bound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"upper_bound\",\"type\":\"int32\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_A\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_B\",\"type\":\"int256\"}],\"internalType\":\"structTokenBank.sync_struct[]\",\"name\":\"syncTx\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"nextPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sqrtPrice\",\"type\":\"uint256\"}],\"name\":\"handle_sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"position_map\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount_tokenA\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount_tokenB\",\"type\":\"int256\"},{\"internalType\":\"int32\",\"name\":\"lower_bound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"upper_bound\",\"type\":\"int32\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_A\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"fees_earned_B\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBankMetaData.ABI instead.
var TokenBankABI = TokenBankMetaData.ABI

// TokenBank is an auto generated Go binding around an Ethereum contract.
type TokenBank struct {
	TokenBankCaller     // Read-only binding to the contract
	TokenBankTransactor // Write-only binding to the contract
	TokenBankFilterer   // Log filterer for contract events
}

// TokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBankSession struct {
	Contract     *TokenBank        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBankCallerSession struct {
	Contract *TokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBankTransactorSession struct {
	Contract     *TokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBankRaw struct {
	Contract *TokenBank // Generic contract binding to access the raw methods on
}

// TokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBankCallerRaw struct {
	Contract *TokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBankTransactorRaw struct {
	Contract *TokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBank creates a new instance of TokenBank, bound to a specific deployed contract.
func NewTokenBank(address common.Address, backend bind.ContractBackend) (*TokenBank, error) {
	contract, err := bindTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBank{TokenBankCaller: TokenBankCaller{contract: contract}, TokenBankTransactor: TokenBankTransactor{contract: contract}, TokenBankFilterer: TokenBankFilterer{contract: contract}}, nil
}

// NewTokenBankCaller creates a new read-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankCaller(address common.Address, caller bind.ContractCaller) (*TokenBankCaller, error) {
	contract, err := bindTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankCaller{contract: contract}, nil
}

// NewTokenBankTransactor creates a new write-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBankTransactor, error) {
	contract, err := bindTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankTransactor{contract: contract}, nil
}

// NewTokenBankFilterer creates a new log filterer instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBankFilterer, error) {
	contract, err := bindTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBankFilterer{contract: contract}, nil
}

// bindTokenBank binds a generic wrapper to an already deployed contract.
func bindTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.TokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transact(opts, method, params...)
}

// DepositMap is a free data retrieval call binding the contract method 0x63edec25.
//
// Solidity: function deposit_map(uint256 ) view returns(address sepoliaAddr, bool initialized)
func (_TokenBank *TokenBankCaller) DepositMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SepoliaAddr common.Address
	Initialized bool
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "deposit_map", arg0)

	outstruct := new(struct {
		SepoliaAddr common.Address
		Initialized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SepoliaAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Initialized = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// DepositMap is a free data retrieval call binding the contract method 0x63edec25.
//
// Solidity: function deposit_map(uint256 ) view returns(address sepoliaAddr, bool initialized)
func (_TokenBank *TokenBankSession) DepositMap(arg0 *big.Int) (struct {
	SepoliaAddr common.Address
	Initialized bool
}, error) {
	return _TokenBank.Contract.DepositMap(&_TokenBank.CallOpts, arg0)
}

// DepositMap is a free data retrieval call binding the contract method 0x63edec25.
//
// Solidity: function deposit_map(uint256 ) view returns(address sepoliaAddr, bool initialized)
func (_TokenBank *TokenBankCallerSession) DepositMap(arg0 *big.Int) (struct {
	SepoliaAddr common.Address
	Initialized bool
}, error) {
	return _TokenBank.Contract.DepositMap(&_TokenBank.CallOpts, arg0)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 tokenAval, uint256 tokenBval)
func (_TokenBank *TokenBankCaller) GetContractBalance(opts *bind.CallOpts) (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getContractBalance")

	outstruct := new(struct {
		TokenAval *big.Int
		TokenBval *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenBval = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 tokenAval, uint256 tokenBval)
func (_TokenBank *TokenBankSession) GetContractBalance() (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	return _TokenBank.Contract.GetContractBalance(&_TokenBank.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 tokenAval, uint256 tokenBval)
func (_TokenBank *TokenBankCallerSession) GetContractBalance() (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	return _TokenBank.Contract.GetContractBalance(&_TokenBank.CallOpts)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xf41032d2.
//
// Solidity: function getDepositBalance(uint256 _addr, uint256 epoch) view returns(int256 tokenAval, int256 tokenBval)
func (_TokenBank *TokenBankCaller) GetDepositBalance(opts *bind.CallOpts, _addr *big.Int, epoch *big.Int) (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getDepositBalance", _addr, epoch)

	outstruct := new(struct {
		TokenAval *big.Int
		TokenBval *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenBval = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xf41032d2.
//
// Solidity: function getDepositBalance(uint256 _addr, uint256 epoch) view returns(int256 tokenAval, int256 tokenBval)
func (_TokenBank *TokenBankSession) GetDepositBalance(_addr *big.Int, epoch *big.Int) (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	return _TokenBank.Contract.GetDepositBalance(&_TokenBank.CallOpts, _addr, epoch)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xf41032d2.
//
// Solidity: function getDepositBalance(uint256 _addr, uint256 epoch) view returns(int256 tokenAval, int256 tokenBval)
func (_TokenBank *TokenBankCallerSession) GetDepositBalance(_addr *big.Int, epoch *big.Int) (struct {
	TokenAval *big.Int
	TokenBval *big.Int
}, error) {
	return _TokenBank.Contract.GetDepositBalance(&_TokenBank.CallOpts, _addr, epoch)
}

// GetPosition is a free data retrieval call binding the contract method 0xb1ad0fa0.
//
// Solidity: function getPosition(address _addr, string _uid) view returns((int256,int256,int32,int32,int256,int256))
func (_TokenBank *TokenBankCaller) GetPosition(opts *bind.CallOpts, _addr common.Address, _uid string) (TokenBankpositionStuct, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getPosition", _addr, _uid)

	if err != nil {
		return *new(TokenBankpositionStuct), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenBankpositionStuct)).(*TokenBankpositionStuct)

	return out0, err

}

// GetPosition is a free data retrieval call binding the contract method 0xb1ad0fa0.
//
// Solidity: function getPosition(address _addr, string _uid) view returns((int256,int256,int32,int32,int256,int256))
func (_TokenBank *TokenBankSession) GetPosition(_addr common.Address, _uid string) (TokenBankpositionStuct, error) {
	return _TokenBank.Contract.GetPosition(&_TokenBank.CallOpts, _addr, _uid)
}

// GetPosition is a free data retrieval call binding the contract method 0xb1ad0fa0.
//
// Solidity: function getPosition(address _addr, string _uid) view returns((int256,int256,int32,int32,int256,int256))
func (_TokenBank *TokenBankCallerSession) GetPosition(_addr common.Address, _uid string) (TokenBankpositionStuct, error) {
	return _TokenBank.Contract.GetPosition(&_TokenBank.CallOpts, _addr, _uid)
}

// PositionMap is a free data retrieval call binding the contract method 0x2055c42b.
//
// Solidity: function position_map(address , string ) view returns(int256 amount_tokenA, int256 amount_tokenB, int32 lower_bound, int32 upper_bound, int256 fees_earned_A, int256 fees_earned_B)
func (_TokenBank *TokenBankCaller) PositionMap(opts *bind.CallOpts, arg0 common.Address, arg1 string) (struct {
	AmountTokenA *big.Int
	AmountTokenB *big.Int
	LowerBound   int32
	UpperBound   int32
	FeesEarnedA  *big.Int
	FeesEarnedB  *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "position_map", arg0, arg1)

	outstruct := new(struct {
		AmountTokenA *big.Int
		AmountTokenB *big.Int
		LowerBound   int32
		UpperBound   int32
		FeesEarnedA  *big.Int
		FeesEarnedB  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountTokenA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountTokenB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LowerBound = *abi.ConvertType(out[2], new(int32)).(*int32)
	outstruct.UpperBound = *abi.ConvertType(out[3], new(int32)).(*int32)
	outstruct.FeesEarnedA = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FeesEarnedB = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PositionMap is a free data retrieval call binding the contract method 0x2055c42b.
//
// Solidity: function position_map(address , string ) view returns(int256 amount_tokenA, int256 amount_tokenB, int32 lower_bound, int32 upper_bound, int256 fees_earned_A, int256 fees_earned_B)
func (_TokenBank *TokenBankSession) PositionMap(arg0 common.Address, arg1 string) (struct {
	AmountTokenA *big.Int
	AmountTokenB *big.Int
	LowerBound   int32
	UpperBound   int32
	FeesEarnedA  *big.Int
	FeesEarnedB  *big.Int
}, error) {
	return _TokenBank.Contract.PositionMap(&_TokenBank.CallOpts, arg0, arg1)
}

// PositionMap is a free data retrieval call binding the contract method 0x2055c42b.
//
// Solidity: function position_map(address , string ) view returns(int256 amount_tokenA, int256 amount_tokenB, int32 lower_bound, int32 upper_bound, int256 fees_earned_A, int256 fees_earned_B)
func (_TokenBank *TokenBankCallerSession) PositionMap(arg0 common.Address, arg1 string) (struct {
	AmountTokenA *big.Int
	AmountTokenB *big.Int
	LowerBound   int32
	UpperBound   int32
	FeesEarnedA  *big.Int
	FeesEarnedB  *big.Int
}, error) {
	return _TokenBank.Contract.PositionMap(&_TokenBank.CallOpts, arg0, arg1)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_TokenBank *TokenBankCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_TokenBank *TokenBankSession) TokenA() (common.Address, error) {
	return _TokenBank.Contract.TokenA(&_TokenBank.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_TokenBank *TokenBankCallerSession) TokenA() (common.Address, error) {
	return _TokenBank.Contract.TokenA(&_TokenBank.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_TokenBank *TokenBankCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_TokenBank *TokenBankSession) TokenB() (common.Address, error) {
	return _TokenBank.Contract.TokenB(&_TokenBank.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_TokenBank *TokenBankCallerSession) TokenB() (common.Address, error) {
	return _TokenBank.Contract.TokenB(&_TokenBank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0705eb01.
//
// Solidity: function deposit(int256 _amount, bool _type, uint256 epoch, uint256 sidechainAddr) payable returns()
func (_TokenBank *TokenBankTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _type bool, epoch *big.Int, sidechainAddr *big.Int) (*types.Transaction, error) {
	return _TokenBank.contract.Transact(opts, "deposit", _amount, _type, epoch, sidechainAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0x0705eb01.
//
// Solidity: function deposit(int256 _amount, bool _type, uint256 epoch, uint256 sidechainAddr) payable returns()
func (_TokenBank *TokenBankSession) Deposit(_amount *big.Int, _type bool, epoch *big.Int, sidechainAddr *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.Deposit(&_TokenBank.TransactOpts, _amount, _type, epoch, sidechainAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0x0705eb01.
//
// Solidity: function deposit(int256 _amount, bool _type, uint256 epoch, uint256 sidechainAddr) payable returns()
func (_TokenBank *TokenBankTransactorSession) Deposit(_amount *big.Int, _type bool, epoch *big.Int, sidechainAddr *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.Deposit(&_TokenBank.TransactOpts, _amount, _type, epoch, sidechainAddr)
}

// HandleSync is a paid mutator transaction binding the contract method 0x0db3397b.
//
// Solidity: function handle_sync((bool,uint256,string,int256,int256,int32,int32,int256,int256)[] syncTx, uint256 epoch, uint256[4] nextPubkey, uint256[2] signature, uint256 liquidity, uint256 sqrtPrice) returns()
func (_TokenBank *TokenBankTransactor) HandleSync(opts *bind.TransactOpts, syncTx []TokenBanksyncStruct, epoch *big.Int, nextPubkey [4]*big.Int, signature [2]*big.Int, liquidity *big.Int, sqrtPrice *big.Int) (*types.Transaction, error) {
	return _TokenBank.contract.Transact(opts, "handle_sync", syncTx, epoch, nextPubkey, signature, liquidity, sqrtPrice)
}

// HandleSync is a paid mutator transaction binding the contract method 0x0db3397b.
//
// Solidity: function handle_sync((bool,uint256,string,int256,int256,int32,int32,int256,int256)[] syncTx, uint256 epoch, uint256[4] nextPubkey, uint256[2] signature, uint256 liquidity, uint256 sqrtPrice) returns()
func (_TokenBank *TokenBankSession) HandleSync(syncTx []TokenBanksyncStruct, epoch *big.Int, nextPubkey [4]*big.Int, signature [2]*big.Int, liquidity *big.Int, sqrtPrice *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.HandleSync(&_TokenBank.TransactOpts, syncTx, epoch, nextPubkey, signature, liquidity, sqrtPrice)
}

// HandleSync is a paid mutator transaction binding the contract method 0x0db3397b.
//
// Solidity: function handle_sync((bool,uint256,string,int256,int256,int32,int32,int256,int256)[] syncTx, uint256 epoch, uint256[4] nextPubkey, uint256[2] signature, uint256 liquidity, uint256 sqrtPrice) returns()
func (_TokenBank *TokenBankTransactorSession) HandleSync(syncTx []TokenBanksyncStruct, epoch *big.Int, nextPubkey [4]*big.Int, signature [2]*big.Int, liquidity *big.Int, sqrtPrice *big.Int) (*types.Transaction, error) {
	return _TokenBank.Contract.HandleSync(&_TokenBank.TransactOpts, syncTx, epoch, nextPubkey, signature, liquidity, sqrtPrice)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenBank *TokenBankTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenBank *TokenBankSession) WithdrawAll() (*types.Transaction, error) {
	return _TokenBank.Contract.WithdrawAll(&_TokenBank.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenBank *TokenBankTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _TokenBank.Contract.WithdrawAll(&_TokenBank.TransactOpts)
}
