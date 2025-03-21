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

// LpManageMetaData contains all meta data concerning the LpManage contract.
var LpManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"contractINonfungiblePositionManager\",\"name\":\"_nonfungiblePositionManager\",\"type\":\"address\"},{\"internalType\":\"contractISwapRouter\",\"name\":\"_swapRouter\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"}],\"name\":\"FreshMint\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ABTX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ABTY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"lower\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"upper\",\"type\":\"uint160\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAdd0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAdd1\",\"type\":\"uint256\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1desired\",\"type\":\"uint256\"}],\"name\":\"mintNewPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonfungiblePositionManager\",\"outputs\":[{\"internalType\":\"contractINonfungiblePositionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"retrieveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LpManageABI is the input ABI used to generate the binding from.
// Deprecated: Use LpManageMetaData.ABI instead.
var LpManageABI = LpManageMetaData.ABI

// LpManage is an auto generated Go binding around an Ethereum contract.
type LpManage struct {
	LpManageCaller     // Read-only binding to the contract
	LpManageTransactor // Write-only binding to the contract
	LpManageFilterer   // Log filterer for contract events
}

// LpManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpManageSession struct {
	Contract     *LpManage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpManageCallerSession struct {
	Contract *LpManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LpManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpManageTransactorSession struct {
	Contract     *LpManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LpManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpManageRaw struct {
	Contract *LpManage // Generic contract binding to access the raw methods on
}

// LpManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpManageCallerRaw struct {
	Contract *LpManageCaller // Generic read-only contract binding to access the raw methods on
}

// LpManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpManageTransactorRaw struct {
	Contract *LpManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLpManage creates a new instance of LpManage, bound to a specific deployed contract.
func NewLpManage(address common.Address, backend bind.ContractBackend) (*LpManage, error) {
	contract, err := bindLpManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LpManage{LpManageCaller: LpManageCaller{contract: contract}, LpManageTransactor: LpManageTransactor{contract: contract}, LpManageFilterer: LpManageFilterer{contract: contract}}, nil
}

// NewLpManageCaller creates a new read-only instance of LpManage, bound to a specific deployed contract.
func NewLpManageCaller(address common.Address, caller bind.ContractCaller) (*LpManageCaller, error) {
	contract, err := bindLpManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpManageCaller{contract: contract}, nil
}

// NewLpManageTransactor creates a new write-only instance of LpManage, bound to a specific deployed contract.
func NewLpManageTransactor(address common.Address, transactor bind.ContractTransactor) (*LpManageTransactor, error) {
	contract, err := bindLpManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpManageTransactor{contract: contract}, nil
}

// NewLpManageFilterer creates a new log filterer instance of LpManage, bound to a specific deployed contract.
func NewLpManageFilterer(address common.Address, filterer bind.ContractFilterer) (*LpManageFilterer, error) {
	contract, err := bindLpManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpManageFilterer{contract: contract}, nil
}

// bindLpManage binds a generic wrapper to an already deployed contract.
func bindLpManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LpManageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpManage *LpManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpManage.Contract.LpManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpManage *LpManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManage.Contract.LpManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpManage *LpManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpManage.Contract.LpManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpManage *LpManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpManage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpManage *LpManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpManage *LpManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpManage.Contract.contract.Transact(opts, method, params...)
}

// ABTX is a free data retrieval call binding the contract method 0x40f0e2f3.
//
// Solidity: function ABTX() view returns(address)
func (_LpManage *LpManageCaller) ABTX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "ABTX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ABTX is a free data retrieval call binding the contract method 0x40f0e2f3.
//
// Solidity: function ABTX() view returns(address)
func (_LpManage *LpManageSession) ABTX() (common.Address, error) {
	return _LpManage.Contract.ABTX(&_LpManage.CallOpts)
}

// ABTX is a free data retrieval call binding the contract method 0x40f0e2f3.
//
// Solidity: function ABTX() view returns(address)
func (_LpManage *LpManageCallerSession) ABTX() (common.Address, error) {
	return _LpManage.Contract.ABTX(&_LpManage.CallOpts)
}

// ABTY is a free data retrieval call binding the contract method 0xf8bd37cc.
//
// Solidity: function ABTY() view returns(address)
func (_LpManage *LpManageCaller) ABTY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "ABTY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ABTY is a free data retrieval call binding the contract method 0xf8bd37cc.
//
// Solidity: function ABTY() view returns(address)
func (_LpManage *LpManageSession) ABTY() (common.Address, error) {
	return _LpManage.Contract.ABTY(&_LpManage.CallOpts)
}

// ABTY is a free data retrieval call binding the contract method 0xf8bd37cc.
//
// Solidity: function ABTY() view returns(address)
func (_LpManage *LpManageCallerSession) ABTY() (common.Address, error) {
	return _LpManage.Contract.ABTY(&_LpManage.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, uint128 liquidity, address token0, address token1)
func (_LpManage *LpManageCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner     common.Address
	Liquidity *big.Int
	Token0    common.Address
	Token1    common.Address
}, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Owner     common.Address
		Liquidity *big.Int
		Token0    common.Address
		Token1    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Liquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, uint128 liquidity, address token0, address token1)
func (_LpManage *LpManageSession) Deposits(arg0 *big.Int) (struct {
	Owner     common.Address
	Liquidity *big.Int
	Token0    common.Address
	Token1    common.Address
}, error) {
	return _LpManage.Contract.Deposits(&_LpManage.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, uint128 liquidity, address token0, address token1)
func (_LpManage *LpManageCallerSession) Deposits(arg0 *big.Int) (struct {
	Owner     common.Address
	Liquidity *big.Int
	Token0    common.Address
	Token1    common.Address
}, error) {
	return _LpManage.Contract.Deposits(&_LpManage.CallOpts, arg0)
}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_LpManage *LpManageCaller) NonfungiblePositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "nonfungiblePositionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_LpManage *LpManageSession) NonfungiblePositionManager() (common.Address, error) {
	return _LpManage.Contract.NonfungiblePositionManager(&_LpManage.CallOpts)
}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_LpManage *LpManageCallerSession) NonfungiblePositionManager() (common.Address, error) {
	return _LpManage.Contract.NonfungiblePositionManager(&_LpManage.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LpManage *LpManageCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LpManage *LpManageSession) Pool() (common.Address, error) {
	return _LpManage.Contract.Pool(&_LpManage.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LpManage *LpManageCallerSession) Pool() (common.Address, error) {
	return _LpManage.Contract.Pool(&_LpManage.CallOpts)
}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_LpManage *LpManageCaller) PoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "poolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_LpManage *LpManageSession) PoolFee() (*big.Int, error) {
	return _LpManage.Contract.PoolFee(&_LpManage.CallOpts)
}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_LpManage *LpManageCallerSession) PoolFee() (*big.Int, error) {
	return _LpManage.Contract.PoolFee(&_LpManage.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_LpManage *LpManageCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_LpManage *LpManageSession) SwapRouter() (common.Address, error) {
	return _LpManage.Contract.SwapRouter(&_LpManage.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_LpManage *LpManageCallerSession) SwapRouter() (common.Address, error) {
	return _LpManage.Contract.SwapRouter(&_LpManage.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_LpManage *LpManageCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_LpManage *LpManageSession) TokenA() (common.Address, error) {
	return _LpManage.Contract.TokenA(&_LpManage.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_LpManage *LpManageCallerSession) TokenA() (common.Address, error) {
	return _LpManage.Contract.TokenA(&_LpManage.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_LpManage *LpManageCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManage.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_LpManage *LpManageSession) TokenB() (common.Address, error) {
	return _LpManage.Contract.TokenB(&_LpManage.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_LpManage *LpManageCallerSession) TokenB() (common.Address, error) {
	return _LpManage.Contract.TokenB(&_LpManage.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xce3f865f.
//
// Solidity: function collect(uint256 tokenId) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactor) Collect(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "collect", tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xce3f865f.
//
// Solidity: function collect(uint256 tokenId) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageSession) Collect(tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.Collect(&_LpManage.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xce3f865f.
//
// Solidity: function collect(uint256 tokenId) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactorSession) Collect(tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.Collect(&_LpManage.TransactOpts, tokenId)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x336bab1a.
//
// Solidity: function decreaseLiquidity(uint256 tokenId, uint256 token0Amount, uint256 token1Amount, uint160 lower, uint160 upper) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactor) DecreaseLiquidity(opts *bind.TransactOpts, tokenId *big.Int, token0Amount *big.Int, token1Amount *big.Int, lower *big.Int, upper *big.Int) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "decreaseLiquidity", tokenId, token0Amount, token1Amount, lower, upper)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x336bab1a.
//
// Solidity: function decreaseLiquidity(uint256 tokenId, uint256 token0Amount, uint256 token1Amount, uint160 lower, uint160 upper) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageSession) DecreaseLiquidity(tokenId *big.Int, token0Amount *big.Int, token1Amount *big.Int, lower *big.Int, upper *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.DecreaseLiquidity(&_LpManage.TransactOpts, tokenId, token0Amount, token1Amount, lower, upper)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x336bab1a.
//
// Solidity: function decreaseLiquidity(uint256 tokenId, uint256 token0Amount, uint256 token1Amount, uint160 lower, uint160 upper) returns(uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactorSession) DecreaseLiquidity(tokenId *big.Int, token0Amount *big.Int, token1Amount *big.Int, lower *big.Int, upper *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.DecreaseLiquidity(&_LpManage.TransactOpts, tokenId, token0Amount, token1Amount, lower, upper)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xe115bc5c.
//
// Solidity: function increaseLiquidity(uint256 tokenId, uint256 amountAdd0, uint256 amountAdd1) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactor) IncreaseLiquidity(opts *bind.TransactOpts, tokenId *big.Int, amountAdd0 *big.Int, amountAdd1 *big.Int) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "increaseLiquidity", tokenId, amountAdd0, amountAdd1)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xe115bc5c.
//
// Solidity: function increaseLiquidity(uint256 tokenId, uint256 amountAdd0, uint256 amountAdd1) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageSession) IncreaseLiquidity(tokenId *big.Int, amountAdd0 *big.Int, amountAdd1 *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.IncreaseLiquidity(&_LpManage.TransactOpts, tokenId, amountAdd0, amountAdd1)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xe115bc5c.
//
// Solidity: function increaseLiquidity(uint256 tokenId, uint256 amountAdd0, uint256 amountAdd1) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactorSession) IncreaseLiquidity(tokenId *big.Int, amountAdd0 *big.Int, amountAdd1 *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.IncreaseLiquidity(&_LpManage.TransactOpts, tokenId, amountAdd0, amountAdd1)
}

// MintNewPosition is a paid mutator transaction binding the contract method 0xf058ebee.
//
// Solidity: function mintNewPosition(uint256 amount0desired, uint256 amount1desired) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactor) MintNewPosition(opts *bind.TransactOpts, amount0desired *big.Int, amount1desired *big.Int) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "mintNewPosition", amount0desired, amount1desired)
}

// MintNewPosition is a paid mutator transaction binding the contract method 0xf058ebee.
//
// Solidity: function mintNewPosition(uint256 amount0desired, uint256 amount1desired) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageSession) MintNewPosition(amount0desired *big.Int, amount1desired *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.MintNewPosition(&_LpManage.TransactOpts, amount0desired, amount1desired)
}

// MintNewPosition is a paid mutator transaction binding the contract method 0xf058ebee.
//
// Solidity: function mintNewPosition(uint256 amount0desired, uint256 amount1desired) returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_LpManage *LpManageTransactorSession) MintNewPosition(amount0desired *big.Int, amount1desired *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.MintNewPosition(&_LpManage.TransactOpts, amount0desired, amount1desired)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_LpManage *LpManageTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "onERC721Received", operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_LpManage *LpManageSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LpManage.Contract.OnERC721Received(&_LpManage.TransactOpts, operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes ) returns(bytes4)
func (_LpManage *LpManageTransactorSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LpManage.Contract.OnERC721Received(&_LpManage.TransactOpts, operator, arg1, tokenId, arg3)
}

// RetrieveNFT is a paid mutator transaction binding the contract method 0x0a1d7c5f.
//
// Solidity: function retrieveNFT(uint256 tokenId) returns()
func (_LpManage *LpManageTransactor) RetrieveNFT(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "retrieveNFT", tokenId)
}

// RetrieveNFT is a paid mutator transaction binding the contract method 0x0a1d7c5f.
//
// Solidity: function retrieveNFT(uint256 tokenId) returns()
func (_LpManage *LpManageSession) RetrieveNFT(tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.RetrieveNFT(&_LpManage.TransactOpts, tokenId)
}

// RetrieveNFT is a paid mutator transaction binding the contract method 0x0a1d7c5f.
//
// Solidity: function retrieveNFT(uint256 tokenId) returns()
func (_LpManage *LpManageTransactorSession) RetrieveNFT(tokenId *big.Int) (*types.Transaction, error) {
	return _LpManage.Contract.RetrieveNFT(&_LpManage.TransactOpts, tokenId)
}

// Swap is a paid mutator transaction binding the contract method 0x2aea6605.
//
// Solidity: function swap(uint256 amountIn, bool zeroForOne) returns(uint256 amountOut)
func (_LpManage *LpManageTransactor) Swap(opts *bind.TransactOpts, amountIn *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _LpManage.contract.Transact(opts, "swap", amountIn, zeroForOne)
}

// Swap is a paid mutator transaction binding the contract method 0x2aea6605.
//
// Solidity: function swap(uint256 amountIn, bool zeroForOne) returns(uint256 amountOut)
func (_LpManage *LpManageSession) Swap(amountIn *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _LpManage.Contract.Swap(&_LpManage.TransactOpts, amountIn, zeroForOne)
}

// Swap is a paid mutator transaction binding the contract method 0x2aea6605.
//
// Solidity: function swap(uint256 amountIn, bool zeroForOne) returns(uint256 amountOut)
func (_LpManage *LpManageTransactorSession) Swap(amountIn *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _LpManage.Contract.Swap(&_LpManage.TransactOpts, amountIn, zeroForOne)
}

// LpManageFreshMintIterator is returned from FilterFreshMint and is used to iterate over the raw logs and unpacked data for FreshMint events raised by the LpManage contract.
type LpManageFreshMintIterator struct {
	Event *LpManageFreshMint // Event containing the contract specifics and raw log

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
func (it *LpManageFreshMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManageFreshMint)
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
		it.Event = new(LpManageFreshMint)
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
func (it *LpManageFreshMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManageFreshMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManageFreshMint represents a FreshMint event raised by the LpManage contract.
type LpManageFreshMint struct {
	TokenId   *big.Int
	OwnerAddr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreshMint is a free log retrieval operation binding the contract event 0x278f9319ed4ed7688a354c86c0ff85ce4bc17dd9f80050a50613898c0a62b221.
//
// Solidity: event FreshMint(uint256 tokenId, address ownerAddr)
func (_LpManage *LpManageFilterer) FilterFreshMint(opts *bind.FilterOpts) (*LpManageFreshMintIterator, error) {

	logs, sub, err := _LpManage.contract.FilterLogs(opts, "FreshMint")
	if err != nil {
		return nil, err
	}
	return &LpManageFreshMintIterator{contract: _LpManage.contract, event: "FreshMint", logs: logs, sub: sub}, nil
}

// WatchFreshMint is a free log subscription operation binding the contract event 0x278f9319ed4ed7688a354c86c0ff85ce4bc17dd9f80050a50613898c0a62b221.
//
// Solidity: event FreshMint(uint256 tokenId, address ownerAddr)
func (_LpManage *LpManageFilterer) WatchFreshMint(opts *bind.WatchOpts, sink chan<- *LpManageFreshMint) (event.Subscription, error) {

	logs, sub, err := _LpManage.contract.WatchLogs(opts, "FreshMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManageFreshMint)
				if err := _LpManage.contract.UnpackLog(event, "FreshMint", log); err != nil {
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

// ParseFreshMint is a log parse operation binding the contract event 0x278f9319ed4ed7688a354c86c0ff85ce4bc17dd9f80050a50613898c0a62b221.
//
// Solidity: event FreshMint(uint256 tokenId, address ownerAddr)
func (_LpManage *LpManageFilterer) ParseFreshMint(log types.Log) (*LpManageFreshMint, error) {
	event := new(LpManageFreshMint)
	if err := _LpManage.contract.UnpackLog(event, "FreshMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
