package ammcore

import (
	"ammcore/TokenBank"
	"ammcore/lib"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var NonceSyncMap sync.Map
var OtherSyncMap sync.Map
var urlBaseline string = "https://eth-sepolia.blastapi.io/aaec52d7-8f95-4c33-a82e-04bcc135aaa1"
var lpManagerADDR string = "0xb39b2A8a4dcD1B33D512c94dC6E54Fe336AB2f06"

func GetNonceFromSyncMap(key common.Address) uint64 {
	keyString := key.Hex()
	mutexObject, _ := OtherSyncMap.Load(key.Hex())
	mutex := mutexObject.(sync.Mutex)
	mutex.Lock()
	// read current nonce value, return it, increment by 1 and save it for next call made
	value, _ := NonceSyncMap.Load(keyString)
	currentVal, _ := value.(uint64)
	NonceSyncMap.Store(keyString, currentVal+1)
	mutex.Unlock()
	return currentVal

}

func CallMint(nonce uint64, privateKey *ecdsa.PrivateKey, amount0, amount1 *big.Int) ([][]interface{}, []*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	tx, err := lpManagerInstance.MintNewPosition(opts, amount0, amount1)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	fmt.Printf("Mint has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	fmt.Println("Mint has been mined: ", reciept.TxHash.String())

	transactions = append(transactions, reciept)

	query := ethereum.FilterQuery{
		FromBlock: reciept.BlockNumber,
		ToBlock:   reciept.BlockNumber,
		Addresses: []common.Address{lpManagerAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(TokenBank.LpManageABI)))
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	var inter [][]interface{}

	for _, log := range logs {
		cur, err := contractAbi.Unpack("FreshMint", log.Data)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		inter = append(inter, cur)
	}

	return inter, transactions, nil
}

func CallIncreaseLiquidity(nonce uint64, privateKey *ecdsa.PrivateKey, amount0, amount1, tokenID *big.Int) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	tx, err := lpManagerInstance.IncreaseLiquidity(opts, tokenID, amount0, amount1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Increase Liquidity has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Increase Liquidity has been mined: ", reciept.TxHash.String())

	transactions = append(transactions, reciept)
	return transactions, nil
}

func CallDecreaseLiquidity(nonce uint64, privateKey *ecdsa.PrivateKey, amount0, amount1, tokenID *big.Int) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	lower, _ := lib.GetSqrtRatioAtTick(-887220)
	upper, _ := lib.GetSqrtRatioAtTick(887220)

	tx, err := lpManagerInstance.DecreaseLiquidity(opts, tokenID, amount0, amount1, lower, upper)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Decrease Liquidity has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Decrease Liquidity has been mined: ", reciept.TxHash.String())

	transactions = append(transactions, reciept)

	return transactions, nil
}

func CheckDeposits(tokenID *big.Int) {
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	output, err := lpManagerInstance.Deposits(nil, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.Liquidity)
	fmt.Println(output.Owner)

}

func CallCollect(nonce uint64, privateKey *ecdsa.PrivateKey, tokenID *big.Int) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	tx, err := lpManagerInstance.Collect(opts, tokenID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Collect has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Collect has been mined: ", reciept.TxHash.String())
	transactions = append(transactions, reciept)
	return transactions, nil
}

func CallRetrNFT(privateKey *ecdsa.PrivateKey, tokenID *big.Int) {
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	tx, err := lpManagerInstance.RetrieveNFT(opts, tokenID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Retrieve has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Retrieve has been mined: ", reciept.TxHash.String())

}

func CallSwap(nonce uint64, privateKey *ecdsa.PrivateKey, amount0 *big.Int, zeroForOne bool) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)
	lpManagerInstance, err := TokenBank.NewLpManage(lpManagerAddress, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	tx, err := lpManagerInstance.Swap(opts, amount0, zeroForOne)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Swap has been confirmed: 0x%x\n", tx.Hash())

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Swap has been mined: ", reciept.TxHash.String())
	transactions = append(transactions, reciept)

	return transactions, nil
}

func CallApproveABTX(nonce uint64, privateKey *ecdsa.PrivateKey, amount0 *big.Int) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	addressABTX := common.HexToAddress("0x208A653d61f568C19924FE0F404431b62Ba1B82f")
	instanceABTX, err := TokenBank.NewERC20(addressABTX, client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tx, err := instanceABTX.Approve(opts, lpManagerAddress, new(big.Int).Mul(amount0, big.NewInt(2)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Approval has been mined: ", reciept.TxHash.String())
	transactions = append(transactions, reciept)
	return transactions, nil
}

func CallApproveABTY(nonce uint64, privateKey *ecdsa.PrivateKey, amount0 *big.Int) ([]*types.Receipt, error) {
	var transactions []*types.Receipt
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lpManagerAddress := common.HexToAddress(lpManagerADDR)

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	addressABTY := common.HexToAddress("0x65d1845521eB8029CB8777f29de5a12DAC34AFbF")
	instanceABTY, err := TokenBank.NewERC20(addressABTY, client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tx, err := instanceABTY.Approve(opts, lpManagerAddress, new(big.Int).Mul(amount0, big.NewInt(2)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reciept, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Approval has been mined: ", reciept.TxHash.String())
	transactions = append(transactions, reciept)
	return transactions, nil
}

func GetCurrentNonce(addr common.Address) uint64 {
	client, err := ethclient.Dial(urlBaseline)
	if err != nil {
		fmt.Println(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		fmt.Println(err)
	}
	return nonce
}

func CheckCurrentGasPrice(upperGasLimit *big.Int, sleepDuration int) {
	var acceptable_gas_price = false
	for !acceptable_gas_price {
		suggestedGasPrice := GetCurrentGasPrice()
		fmt.Println("Current gas price: ", suggestedGasPrice)
		if suggestedGasPrice.Cmp(upperGasLimit) <= 0 {
			acceptable_gas_price = true
		} else {
			fmt.Println("gas price too high")
			time.Sleep(time.Minute * time.Duration(sleepDuration))
		}
	}
}

func GetCurrentGasPrice() *big.Int {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return suggestedGasPrice
}
