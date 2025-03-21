package lib

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

// calcAmount1Delta calculates the change in token1 amount based on the provided square root prices and liquidity.
func CalcAmount1Delta(sqrtPriceAX96, sqrtPriceBX96, liquidity *big.Int) *big.Int {
	// If sqrtPriceAX96 > sqrtPriceBX96 then swap their values
	if sqrtPriceAX96.Cmp(sqrtPriceBX96) > 0 {
		sqrtPriceAX96, sqrtPriceBX96 = sqrtPriceBX96, sqrtPriceAX96
	}

	// Create a new variable amount1
	amount1 := new(big.Int)

	// Calculate amount1 using mulDivRoundingUp function
	amount1 = MulDivRoundingUp(liquidity, new(big.Int).Sub(sqrtPriceBX96, sqrtPriceAX96), new(big.Int).Exp(big.NewInt(2), big.NewInt(84), nil))

	return amount1
}

// calcAmount0Delta calculates the amount0 delta based on input parameters.
func CalcAmount0Delta(sqrtPriceAX96, sqrtPriceBX96, liquidity *big.Int) *big.Int {
	// If sqrtPriceAX96 > sqrtPriceBX96, swap their values
	if sqrtPriceAX96.Cmp(sqrtPriceBX96) > 0 {
		sqrtPriceAX96, sqrtPriceBX96 = sqrtPriceBX96, sqrtPriceAX96
	}

	// Check if sqrtPriceAX96 > 0
	if sqrtPriceAX96.Sign() <= 0 {
		panic("sqrtPriceAX96 must be greater than zero")
	}

	// Calculate amount0
	amount0 := new(big.Int).Set(liquidity)
	amount0.Lsh(amount0, 96)
	amount0.Mul(amount0, new(big.Int).Sub(sqrtPriceBX96, sqrtPriceAX96))
	amount0.Div(amount0, sqrtPriceBX96)
	amount0.Div(amount0, sqrtPriceAX96)

	return amount0
}

func AddLiquidity(x, y *big.Int) *big.Int {
	z := new(big.Int)
	if y.Cmp(big.NewInt(0)) == -1 {
		neg_y := new(big.Int)
		neg_y.Neg(y)
		z.Sub(x, neg_y)
	} else {
		z.Add(x, y)
	}
	return z
}

func GetNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if amount.Cmp(Zero) == 0 {
		return sqrtPX96, nil
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	if add {
		product := MultiplyIn256(amount, sqrtPX96)
		if new(big.Int).Div(product, amount).Cmp(sqrtPX96) == 0 {
			denominator := AddIn256(numerator1, product)
			if denominator.Cmp(numerator1) >= 0 {
				return MulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
			}
		}
		return MulDivRoundingUp(numerator1, One, new(big.Int).Add(new(big.Int).Div(numerator1, sqrtPX96), amount)), nil
	} else {
		product := MultiplyIn256(amount, sqrtPX96)
		/*if new(big.Int).Div(product, amount).Cmp(sqrtPX96) != 0 {
			return nil, errors.New(fmt.Sprint("invariant violation: ", product, " sqrtPX96: ", sqrtPX96, " Amount: ", amount))
		}*/
		if numerator1.Cmp(product) <= 0 {
			return nil, errors.New(fmt.Sprint("invariant violation Numerator: ", numerator1, " sqrtPX96: ", sqrtPX96, " Amount: ", amount))
		}
		denominator := new(big.Int).Sub(numerator1, product)
		return MulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
	}
}

func GetNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if add {
		var quotient *big.Int
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = new(big.Int).Div(new(big.Int).Lsh(amount, 96), liquidity)
		} else {
			quotient = new(big.Int).Div(new(big.Int).Mul(amount, Q96), liquidity)
		}
		return new(big.Int).Add(sqrtPX96, quotient), nil
	}

	var quotient *big.Int
	if amount.Cmp(MaxUint160) <= 0 {
		quotient = DivRoundingUp(amount, liquidity)
	 } else {
		quotient = MulDivRoundingUp(amount, Q96, liquidity)
	 }

	if sqrtPX96.Cmp(quotient) <= 0 {
		return nil, errors.New(fmt.Sprint("invariant violation Quotient: ", quotient, " sqrtPX96: ", sqrtPX96, " liquidity ", liquidity))

	}
	return new(big.Int).Sub(sqrtPX96, quotient), nil
}

// Function to calculate the square root price
func CalculateSqrtPrice(tick int32) *big.Int {
	sqrtPrice := new(big.Float).SetFloat64(math.Pow(1.0001, float64(tick)))
	sqrtPrice.Mul(sqrtPrice, big.NewFloat(math.Pow(2, 96)))
	result := new(big.Int)
	sqrtPrice.Int(result)
	return result
}

func GetAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) >= 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	numerator2 := new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)

	if roundUp {
		return MulDivRoundingUp(MulDivRoundingUp(numerator1, numerator2, sqrtRatioBX96), One, sqrtRatioAX96)
	}
	return MulDivRoundingDown(MulDivRoundingDown(numerator1, numerator2, sqrtRatioBX96), One, sqrtRatioAX96)
}

func GetAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	if roundUp {
		return MulDivRoundingUp(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), Q96)
	}
	return MulDivRoundingDown(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), Q96)
}

func GetAmount0Delta_noBool(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int) *big.Int {
	output := new(big.Int)
	if liquidity.Cmp(Zero) == -1 {
		liquidityNeg := new(big.Int)
		output = GetAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidityNeg.Neg(liquidity), false)
		output.Neg(output)
	} else {
		output = GetAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity, true)
	}
	return output
}

func GetAmount1Delta_noBool(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int) *big.Int {
	output := new(big.Int)
	if liquidity.Cmp(Zero) == -1 {
		liquidityNeg := new(big.Int)
		output = GetAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidityNeg.Neg(liquidity), false)
		output.Neg(output)
	} else {
		output = GetAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity, true)
	}
	return output
}

func MaxLiquidityForAmount0Precise(sqrtRatioAX96, sqrtRatioBX96, amount0 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	numerator := new(big.Int).Mul(new(big.Int).Mul(amount0, sqrtRatioAX96), sqrtRatioBX96)
	denominator := new(big.Int).Mul(Q96, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	return new(big.Int).Div(numerator, denominator)
}

func MaxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	return new(big.Int).Div(new(big.Int).Mul(amount1, Q96), new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
}

func MaxLiquidityForAmount0Imprecise(sqrtRatioAX96, sqrtRatioBX96, amount0 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	intermediate := new(big.Int).Div(new(big.Int).Mul(sqrtRatioAX96, sqrtRatioBX96), Q96)
	return new(big.Int).Div(new(big.Int).Mul(amount0, intermediate), new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
}

func MaxLiquidityForAmounts(sqrtRatioCurrentX96 *big.Int, sqrtRatioAX96, sqrtRatioBX96 *big.Int, amount0, amount1 *big.Int, useFullPrecision bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	var maxLiquidityForAmount0 func(*big.Int, *big.Int, *big.Int) *big.Int
	if useFullPrecision {
		maxLiquidityForAmount0 = MaxLiquidityForAmount0Precise
	} else {
		maxLiquidityForAmount0 = MaxLiquidityForAmount0Imprecise
	}
	if sqrtRatioCurrentX96.Cmp(sqrtRatioAX96) <= 0 {
		return maxLiquidityForAmount0(sqrtRatioAX96, sqrtRatioBX96, amount0)
	} else if sqrtRatioCurrentX96.Cmp(sqrtRatioBX96) < 0 {
		liquidity0 := maxLiquidityForAmount0(sqrtRatioCurrentX96, sqrtRatioBX96, amount0)
		liquidity1 := MaxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioCurrentX96, amount1)
		if liquidity0.Cmp(liquidity1) < 0 {
			return liquidity0
		}
		return liquidity1

	} else {
		return MaxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1)
	}
}
