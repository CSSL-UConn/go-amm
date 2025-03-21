package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//7c2e78a49c1ce348a4ce4f6fb66caf5e8270bd94f2e3a7e0f02732e53bc62bf0
	//f064b1f3eefc925be42db085a597d5e8967ade6876873cba1070d0ef0c288fee
	hexString := "7c2e78a49c1ce348a4ce4f6fb66caf5e8270bd94f2e3a7e0f02732e53bc62bf0"
	ownerKeyBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
	}
	userPK, err := crypto.ToECDSA(ownerKeyBytes)
	if err != nil {
		fmt.Println(err)
	}
	tokenID := big.NewInt(15621)
	_ = tokenID
	_ = userPK

	//inter, _, _ := ammcore.CallMint(nil, userPK, big.NewInt(2), big.NewInt(2))
	//fmt.Println(inter)
	//TokenBank.RunBaseline()
}
