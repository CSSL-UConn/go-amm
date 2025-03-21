package lib

import (
	"math/big"
)

type FeeAmount uint64

const (
	FeeLowest FeeAmount = 100
	FeeLow    FeeAmount = 500
	FeeMedium FeeAmount = 3000
	FeeHigh   FeeAmount = 10000

	FeeMax FeeAmount = 1000000
)

// The default factory tick spacings by fee amount.
var TickSpacings = map[FeeAmount]int{
	FeeLowest: 1,
	FeeLow:    10,
	FeeMedium: 60,
	FeeHigh:   200,
}

var (
	NegativeOne   = big.NewInt(-1)
	Zero          = big.NewInt(0)
	One           = big.NewInt(1)
	MaxUint256, _ = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	MaxQ128, _    = new(big.Int).SetString("340282366920938463463374607431768211456", 10)
	MaxUint128, _ = new(big.Int).SetString("340282366920938463463374607431768211455", 10)
	MaxUint64, _  = new(big.Int).SetString("18446744073709551615", 10)
	MaxUint32, _  = new(big.Int).SetString("4294967295", 10)
	MaxUint16, _  = new(big.Int).SetString("65535", 10)
	MaxUint8, _   = new(big.Int).SetString("255", 10)
	MaxUint8Int32 = int32(255)
	MaxUint160    = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), One)

	// used in liquidity amount math
	Q96  = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
	Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)

	PercentZero = NewFraction(big.NewInt(0), big.NewInt(1))
)
