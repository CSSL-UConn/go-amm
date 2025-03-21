package ammcore

import (
	"ammcore/TokenBank"
	"ammcore/lib"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url string = "https://eth-sepolia.blastapi.io/aaec52d7-8f95-4c33-a82e-04bcc135aaa1"
var chainId *big.Int = big.NewInt(11155111)
var owner string = "" // change

var tokenBankAddress string = "0xC3c5b2B7c3e2eAc6604FaE1B7177eC3f2b5a0489"
var tokenAAddress string = "0x208A653d61f568C19924FE0F404431b62Ba1B82f"
var tokenBAddress string = "0x65d1845521eB8029CB8777f29de5a12DAC34AFbF"

func GetTokenBalance() {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	// old TB 0x9046CeF0e03C386C42B8B248bd6732dC145c296a
	tbAddress := common.HexToAddress(tokenAAddress)
	instanceTB, err := TokenBank.NewTokenBank(tbAddress, client)
	if err != nil {
		fmt.Println(err)
	}
	tokens, err := instanceTB.GetContractBalance(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokens)
}

func Deposit(mainchainKey *ecdsa.PrivateKey, sidechainAddr *big.Int, amount0, amount1 *big.Int, epoch *big.Int) (receipts []*types.Receipt ) {
	// take amount desired of each token as input, and PK of ethereum address
	receipts = make([]*types.Receipt, 0)
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(mainchainKey, chainId)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	tbAddress := common.HexToAddress(tokenBankAddress)
	instanceTB, err := TokenBank.NewTokenBank(tbAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	addressABTX := common.HexToAddress(tokenAAddress)
	instanceABTX, err := TokenBank.NewERC20(addressABTX, client)
	if err != nil {
		fmt.Println(err)
	}

	addressABTY := common.HexToAddress(tokenBAddress)
	instanceABTY, err := TokenBank.NewERC20(addressABTY, client)
	if err != nil {
		fmt.Println(err)
	}

	bigSidechainAddr := sidechainAddr

	if amount0.Cmp(big.NewInt(0)) > 0 && amount0 != nil {
		// calc amount 0 * 10**18
		tmp := lib.Exponentiate(amount0.Int64(), 18)
		// approve TB for tokens
		opts.Nonce = new(big.Int).SetUint64(GetCurrentNonce(crypto.PubkeyToAddress(mainchainKey.PublicKey)))
		tx, err := instanceABTX.Approve(opts, tbAddress, tmp)
		if err != nil {
			fmt.Println(err)
		}
		reciept, err := bind.WaitMined(opts.Context, client, tx)
		if err != nil {
			fmt.Println(err)
		}
		receipts = append(receipts, reciept)
		fmt.Println("Approval has been mined: ", reciept.TxHash.String())

		// deposit tokens to TB
		opts.Nonce = new(big.Int).SetUint64(GetCurrentNonce(crypto.PubkeyToAddress(mainchainKey.PublicKey)))
		tx, err = instanceTB.Deposit(opts, tmp, true, epoch, bigSidechainAddr)
		if err != nil {
			fmt.Println(err)
		}
		reciept, err = bind.WaitMined(opts.Context, client, tx)

		if err != nil {
			fmt.Println(err)
		}
		receipts = append(receipts, reciept)
		fmt.Println("Deposit has been mined: ", reciept.TxHash.String())

	}
	if amount1.Cmp(big.NewInt(0)) > 0 && amount1 != nil {
		// calc amount1 * 10**18
		tmp := lib.Exponentiate(amount1.Int64(), 18)
		// approve TB for tokens
		opts.Nonce = new(big.Int).SetUint64(GetCurrentNonce(crypto.PubkeyToAddress(mainchainKey.PublicKey)))
		tx, err := instanceABTY.Approve(opts, tbAddress, tmp)
		if err != nil {
			fmt.Println(err)
		}
		reciept, err := bind.WaitMined(opts.Context, client, tx)

		if err != nil {
			fmt.Println(err)
		}
		receipts = append(receipts, reciept)
		fmt.Println("Approval has been mined: ", reciept.TxHash.String())
		// deposit tokens to TB
		opts.Nonce = new(big.Int).SetUint64(GetCurrentNonce(crypto.PubkeyToAddress(mainchainKey.PublicKey)))
		tx, err = instanceTB.Deposit(opts, tmp, false, epoch, bigSidechainAddr)
		if err != nil {
			fmt.Println(err)
		}
		reciept, err = bind.WaitMined(opts.Context, client, tx)
		if err != nil {
			fmt.Println(err)
		}
		receipts = append(receipts, reciept)
		fmt.Printf("Deposit for token 1 has been confirmed: 0x%x\n", tx.Hash())
	}

	return receipts
}

func GetApproval(mainchainKey *ecdsa.PrivateKey) (amount0, amount1 *big.Int) {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}

	tbAddress := common.HexToAddress(tokenBankAddress)
	userAddr := crypto.PubkeyToAddress(mainchainKey.PublicKey)

	addressABTX := common.HexToAddress(tokenAAddress)
	instanceABTX, err := TokenBank.NewERC20(addressABTX, client)
	if err != nil {
		fmt.Println(err)
	}

	addressABTY := common.HexToAddress(tokenBAddress)
	instanceABTY, err := TokenBank.NewERC20(addressABTY, client)
	if err != nil {
		fmt.Println(err)
	}

	amount0, err = instanceABTX.Allowance(nil, userAddr, tbAddress)
	if err != nil {
		fmt.Println(err)
	}
	amount1, err = instanceABTY.Allowance(nil, userAddr, tbAddress)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetDepositBalance(addr *big.Int, epoch *big.Int) (token0, token1 *big.Int) {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	tbAddress := common.HexToAddress(tokenBankAddress)
	instanceTB, err := TokenBank.NewTokenBank(tbAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	bigAddr := addr

	mapping, err := instanceTB.GetDepositBalance(nil, bigAddr, epoch)
	if err != nil {
		fmt.Println(err)
	}
	return mapping.TokenAval, mapping.TokenBval
}

func Synchronize(mainchainKey *ecdsa.PrivateKey, syncData []TokenBank.TokenBanksyncStruct, pubkey [4]*big.Int, sig [2]*big.Int, epoch *big.Int, liquidity *big.Int, price *big.Int)  *types.Receipt {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(mainchainKey, chainId)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	tbAddress := common.HexToAddress(tokenBankAddress)
	instanceTB, err := TokenBank.NewTokenBank(tbAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	opts.GasLimit = 3000000
	opts.Nonce = new(big.Int).SetUint64(GetCurrentNonce(crypto.PubkeyToAddress(mainchainKey.PublicKey)))
	tx, err := instanceTB.HandleSync(opts, syncData, epoch, pubkey, sig, liquidity, price)
	if err != nil {
		panic(err)
	}
	reciept, err := bind.WaitMined(opts.Context, client, tx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deposit for token 1 has been confirmed: 0x%x\n", tx.Hash())
	return reciept
}

func MintTokens(mainchainKey *ecdsa.PrivateKey, amount *big.Int) bool {

	destAddr := crypto.PubkeyToAddress(mainchainKey.PublicKey)

	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// owner key: 7c2e78a49c1ce348a4ce4f6fb66caf5e8270bd94f2e3a7e0f02732e53bc62bf0
	// get owner key

	//ownerKeyBytes := []byte("7c2e78a49c1ce348a4ce4f6fb66caf5e8270bd94f2e3a7e0f02732e53bc62bf0")
	ownerKeyBytes, err := hex.DecodeString(owner)
	if err != nil {
		fmt.Println(err)
		return false
	}
	ownerKey, err := crypto.ToECDSA(ownerKeyBytes)
	//ownerKey, err := x509.ParseECPrivateKey(ownerKeyBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	opts, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainId)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return false
	}
	addressABTX := common.HexToAddress(tokenAAddress)
	instanceABTX, err := TokenBank.NewERC20(addressABTX, client)
	if err != nil {
		fmt.Println(err)
		return false
	}

	addressABTY := common.HexToAddress(tokenBAddress)
	instanceABTY, err := TokenBank.NewERC20(addressABTY, client)
	if err != nil {
		fmt.Println(err)
		return false
	}

	tx, err := instanceABTX.Mint(opts, destAddr, amount)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("mint for token 0 has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Mint token 0 has been mined: ", reciept.TxHash.String())

	tx, err = instanceABTY.Mint(opts, destAddr, amount)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("mint for token 1 has been confirmed: 0x%x\n", tx.Hash())

	reciept, err = bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Mint token 1 has been mined: ", reciept.TxHash.String())

	return true
}

func AbiEncode(syncData []TokenBank.TokenBanksyncStruct) ([]byte, error) {
	dataStruct := make([]interface{}, 0)
	for _, syncDatum := range syncData {
		dataStruct = append(dataStruct, syncDatum)
	}
	syncStruct, _ := abi.NewType("tuple", "struct sync_struct", []abi.ArgumentMarshaling{
		{Name: "tx_type_id", Type: "bool"},
		{Name: "sidechainAddr", Type: "uint256"},
		{Name: "position_id", Type: "string"},
		{Name: "amount_tokenA", Type: "int256"},
		{Name: "amount_tokenB", Type: "int256"},
		{Name: "lower_bound", Type: "int32"},
		{Name: "upper_bound", Type: "int32"},
		{Name: "fees_earned_A", Type: "int256"},
		{Name: "fees_earned_B", Type: "int256"},
	})

	args := abi.Arguments{}

	for i, _ := range syncData {
		args = append(args, abi.Argument{Type: syncStruct, Name: fmt.Sprint(i), Indexed: false})
	}
	prefix := hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020")
	data, err := args.PackValues(dataStruct)
	return append(prefix, data...), err
}

func UnterminatedAbiEncode(syncData []TokenBank.TokenBanksyncStruct) ([]byte, error) {
	dataStruct := make([]interface{}, 0)
	for _, syncDatum := range syncData {
		dataStruct = append(dataStruct, syncDatum)
	}
	syncStruct, _ := abi.NewType("tuple", "struct sync_struct", []abi.ArgumentMarshaling{
		{Name: "tx_type_id", Type: "bool"},
		{Name: "sidechainAddr", Type: "uint256"},
		{Name: "position_id", Type: "string"},
		{Name: "amount_tokenA", Type: "int256"},
		{Name: "amount_tokenB", Type: "int256"},
		{Name: "lower_bound", Type: "int32"},
		{Name: "upper_bound", Type: "int32"},
		{Name: "fees_earned_A", Type: "int256"},
		{Name: "fees_earned_B", Type: "int256"},
	})

	args := abi.Arguments{}

	for i, _ := range syncData {
		args = append(args, abi.Argument{Type: syncStruct, Name: fmt.Sprint(i), Indexed: false})
	}

	data, err := args.PackValues(dataStruct)
	return data, err
}
