package lib

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"runtime"

	"github.com/holiman/uint256"
)

var ErrInvalidInput = errors.New("invalid input")

func EncodeSqrtRatioX96(amount1 *big.Int, amount0 *big.Int) *big.Int {
	numerator := new(big.Int).Lsh(amount1, 192)
	denominator := amount0
	ratioX192 := new(big.Int).Div(numerator, denominator)
	return new(big.Int).Sqrt(ratioX192)
}

func FancyHandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, filename, line, _ := runtime.Caller(1)

		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
		b = true
	}
	return
}

/*
func MostSignificantBit(x *big.Int) (int64, error) {
	if x.Cmp(Zero) <= 0 {
		return 0, ErrInvalidInput
	}
	if x.Cmp(MaxUint256) > 0 {
		return 0, ErrInvalidInput
	}
	var msb int64
	for _, power := range []int64{128, 64, 32, 16, 8, 4, 2, 1} {
		min := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(power)), nil)
		if x.Cmp(min) >= 0 {
			x = new(big.Int).Rsh(x, uint(power))
			msb += power
		}
	}
	return msb, nil
}
*/

func MostSignificantBit(xx *big.Int) int64 {
	if !(xx.Cmp(Zero) != 0) {
		fmt.Printf("Input must be greater than zero")
		return 0
	}
	x := uint256.MustFromBig(xx)
	out := 0
	tmp := uint256.MustFromHex("0x100000000000000000000000000000000")

	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 128)
		out += 128
	}
	tmp = uint256.MustFromHex("0x10000000000000000")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 64)
		out += 64
	}
	tmp = uint256.MustFromHex("0x100000000")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 32)
		out += 32
	}
	tmp = uint256.MustFromHex("0x10000")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 16)
		out += 16
	}
	tmp = uint256.MustFromHex("0x100")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 8)
		out += 8
	}
	tmp = uint256.MustFromHex("0x10")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 4)
		out += 4
	}
	tmp = uint256.MustFromHex("0x4")
	if x.Cmp(tmp) >= 0 {
		x.Rsh(x, 2)
		out += 2
	}
	tmp = uint256.MustFromHex("0x2")
	if x.Cmp(tmp) >= 0 {
		out += 1
	}
	return int64(out)
}

func LeastSignificantBit(x *big.Int) int64 {
	if !(x.Cmp(Zero) > 0) {
		fmt.Println("Input must be greater than zero")
		return 0
	}

	r := 255
	tmp := new(big.Int)
	tmp.And(x, MaxUint128)
	if tmp.Cmp(Zero) > 0 {
		r -= 128
	} else {
		x.Rsh(x, 128)
	}
	tmp.And(x, MaxUint64)
	if tmp.Cmp(Zero) > 0 {
		r -= 64
	} else {
		x.Rsh(x, 64)
	}
	tmp.And(x, MaxUint32)
	if tmp.Cmp(Zero) > 0 {
		r -= 32
	} else {
		x.Rsh(x, 32)
	}
	tmp.And(x, MaxUint16)
	if tmp.Cmp(Zero) > 0 {
		r -= 16
	} else {
		x.Rsh(x, 16)
	}
	tmp.And(x, MaxUint8)
	if tmp.Cmp(Zero) > 0 {
		r -= 8
	} else {
		x.Rsh(x, 8)
	}
	four, _ := new(big.Int).SetString("0xf", 0)
	tmp.And(x, four)
	if tmp.Cmp(Zero) > 0 {
		r -= 4
	} else {
		x.Rsh(x, 4)
	}
	two, _ := new(big.Int).SetString("0x3", 0)
	tmp.And(x, two)
	if tmp.Cmp(Zero) > 0 {
		r -= 2
	} else {
		x.Rsh(x, 2)
	}
	one, _ := new(big.Int).SetString("0x1", 0)
	tmp.And(x, one)
	if tmp.Cmp(Zero) > 0 {
		r -= 1
	}
	return int64(r)
}

func MulDivRoundingUp(a, b, denominator *big.Int) *big.Int {
	// Create a new variable result
	result := new(big.Int)

	// Store the following in result: (a * b) / denominator
	result.Mul(a, b)
	result.Div(result, denominator)

	// If (a*b) % denominator is > 0 then increment result by 1
	remainder := new(big.Int).Mod(new(big.Int).Mul(a, b), denominator)
	if remainder.Sign() > 0 {
		result.Add(result, big.NewInt(1))
	}

	return new(big.Int).And(result, MaxUint256)
}

func DivRoundingUp(a, denominator *big.Int) *big.Int {
	// Create a new variable result
	result := new(big.Int)

	// Store the following in result: (a * b) / denominator
	result.Div(a, denominator)

	// If (a*b) % denominator is > 0 then increment result by 1
	remainder := new(big.Int).Mod(a, denominator)
	if remainder.Sign() > 0 {
		result.Add(result, big.NewInt(1))
	}

	return new(big.Int).And(result, MaxUint256)
}

// Function to perform multiplication and division with rounding down.
func MulDivRoundingDown(a, b, denominator *big.Int) *big.Int {
	// Calculate floor(a*b / denominator)
	result := new(big.Int).Mul(a, b)
	result.Div(result, denominator)
	return new(big.Int).And(result, MaxUint256)
}

// roundUpwards rounds up the given big.Int number.
func RoundUpwards(number *big.Int) *big.Int {
	rounded := new(big.Int).Set(number)
	rounded.Add(rounded, big.NewInt(1))
	return rounded
}

var MaxFee = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

func ComputeSwapStep(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, amountRemaining *big.Int, feePips FeeAmount) (sqrtRatioNextX96, amountIn, amountOut, feeAmount *big.Int, err error) {
	zeroForOne := sqrtRatioCurrentX96.Cmp(sqrtRatioTargetX96) >= 0
	exactIn := amountRemaining.Cmp(Zero) >= 0

	if exactIn {
		amountRemainingLessFee := MulDivRoundingDown(amountRemaining, new(big.Int).Sub(MaxFee, big.NewInt(int64(feePips))), MaxFee)
		//fmt.Println("amountRemainingLessFee: ", amountRemainingLessFee)
		if zeroForOne {
			//fmt.Println("Beforefirst")
			amountIn = GetAmount0Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, true)
			//fmt.Println("amountIn:", amountIn)
		} else {
			//fmt.Println("inside else ")
			amountIn = GetAmount1Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, true)
			//fmt.Println("after get1: ", amountIn)
		}
		if amountRemainingLessFee.Cmp(amountIn) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96, err = GetNextSqrtPriceFromInput(sqrtRatioCurrentX96, liquidity, amountRemainingLessFee, zeroForOne)
			if err != nil {
				FancyHandleError(err)
				panic(err)
				return
			}
		}
	} else {
		if zeroForOne {
			amountOut = GetAmount1Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, false)
		} else {
			amountOut = GetAmount0Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, false)
		}
		if new(big.Int).Mul(amountRemaining, NegativeOne).Cmp(amountOut) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96, err = GetNextSqrtPriceFromOutput(sqrtRatioCurrentX96, liquidity, new(big.Int).Mul(amountRemaining, NegativeOne), zeroForOne)
			if err != nil {
				FancyHandleError(err)
				panic(err)
				return
			}
		}
	}

	max := sqrtRatioTargetX96.Cmp(sqrtRatioNextX96) == 0
	//fmt.Println("exact in: ", exactIn)
	if zeroForOne {
		if !(max && exactIn) {
			amountIn = GetAmount0Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, true)
		}
		if !(max && !exactIn) {
			amountOut = GetAmount1Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, false)
		}
	} else {
		if !(max && exactIn) {
			amountIn = GetAmount1Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, true)
		}
		if !(max && !exactIn) {
			//fmt.Println(sqrtRatioCurrentX96)
			//fmt.Println(sqrtRatioNextX96)
			//fmt.Println(liquidity)
			amountOut = GetAmount0Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, false)
		}
	}

	if !exactIn && amountOut.Cmp(new(big.Int).Mul(amountRemaining, NegativeOne)) > 0 {
		amountOut = new(big.Int).Mul(amountRemaining, NegativeOne)
	}

	if exactIn && sqrtRatioNextX96.Cmp(sqrtRatioTargetX96) != 0 {
		// we didn't reach the target, so take the remainder of the maximum input as fee
		feeAmount = new(big.Int).Sub(amountRemaining, amountIn)
	} else {
		feeAmount = MulDivRoundingUp(amountIn, big.NewInt(int64(feePips)), new(big.Int).Sub(MaxFee, big.NewInt(int64(feePips))))
	}
	return
}

func GetNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(Zero) <= 0 {
		return nil, errors.New("liquidity less than zero")
	}
	if liquidity.Cmp(Zero) <= 0 {
		return nil, errors.New("liquidity less than zero")

	}
	if zeroForOne {
		return GetNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountOut, false)
	}
	return GetNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountOut, false)
}

func GetNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(Zero) <= 0 {
		return nil, errors.New("liquidity less than zero")
	}
	if liquidity.Cmp(Zero) <= 0 {
		return nil, errors.New("liquidity less than zero")
	}
	if zeroForOne {
		return GetNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountIn, true)
	}
	return GetNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountIn, true)
}

func MultiplyIn256(x, y *big.Int) *big.Int {
	product := new(big.Int).Mul(x, y)
	return new(big.Int).And(product, MaxUint256)
}

func AddIn256(x, y *big.Int) *big.Int {
	sum := new(big.Int).Add(x, y)
	return new(big.Int).And(sum, MaxUint256)
}

func Exponentiate(base int64, exponent int64) *big.Int {
	z := big.NewInt(base)
	var i, e = big.NewInt(10), big.NewInt(exponent)
	k := new(big.Int).Exp(i, e, nil)
	return new(big.Int).Mul(z, k)
}

func ZeroFloat() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(14)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return ZeroFloat().Mul(a, b)
}

func Pow(a *big.Float, e uint64) *big.Float {
	result := ZeroFloat().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

func FormatPrice(nbr *big.Int) *big.Float {
	xnbr := new(big.Float).SetInt(nbr)
	power := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil))
	quotient := new(big.Float).Quo(xnbr, power)
	return Pow(quotient, 2)
}
