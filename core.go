package ammcore

import (
	"ammcore/lib"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strconv"

	"golang.org/x/exp/maps"
)

type liquidityPosition struct {
	ID                       string
	LowerBound               int32 // 24-bit integer for lower bound
	UpperBound               int32 // 24-bit integer for upper bound
	feeGrowthInside0LastX128 *big.Int
	feeGrowthInside1LastX128 *big.Int
	Token0Owed               *big.Int // number of tokens from fees owed to position
	Token1Owed               *big.Int // number of tokens from fees owed to position
	Token0Fees               *big.Int
	Token1Fees               *big.Int
	LiquidAmount             *big.Int // measure of the total liqudity owned by the position
	Initialized              bool
}

type Tick_info struct {
	Initialized           bool
	LiquidityGross        *big.Int
	LiquidityNet          *big.Int
	feeGrowthOutside0X128 *big.Int
	feeGrowthOutside1X128 *big.Int
}

// liquidityPool represents a liquidity pool.
type LiquidityPool struct {
	Positions            map[string]liquidityPosition
	PositionSummaries    map[string]position_Summary // mapping of position summary struct
	SwapSummaries        map[string]swap_Summary
	ticks                map[int32]Tick_info // Slice to store tick instances
	Token0TotalLiquidity *big.Int            // Summation of Token0Amount across all positions
	Token1TotalLiquidity *big.Int            // Summation of Token1Amount across all positions
	MAX_TICK             int32               // Maximum tick value
	MIN_TICK             int32               // Minimum tick value
	tick_bitmap          map[int16]*big.Int
	tick_spacing         int32
	fee                  uint64
	Current_tick         int32
	Price                *big.Int
	feeGrowthGlobal0X128 *big.Int
	feeGrowthGlobal1X128 *big.Int
	ModifiedPositions    []string
	Liquidity            *big.Int
}

// SwapState represents the state of a swap.
type SwapState struct {
	AmountSpecifiedRemaining *big.Int // uint256 amountSpecifiedRemaining
	AmountCalculated         *big.Int // uint256 amountCalculated
	SqrtPriceX96             *big.Int // uint160 sqrtPriceX96
	Tick                     int32    // int24 tick
	feeGrowthGlobalX128      *big.Int
	liquidity                *big.Int
}

// StepState represents the state of a step.
type StepState struct {
	SqrtPriceStartX96 *big.Int // uint160 sqrtPriceStartX96
	NextTick          int32    // int24 nextTick
	Initialized       bool
	SqrtPriceNextX96  *big.Int // uint160 sqrtPriceNextX96
	AmountIn          *big.Int // uint256 amountIn
	AmountOut         *big.Int // uint256 amountOut
	feeAmount         *big.Int
}

type position_Summary struct {
	ID          string
	LowerBound  int32
	UpperBound  int32
	FeeChangeT0 *big.Int
	FeeChangeT1 *big.Int
	DeltaT0     *big.Int
	DeltaT1     *big.Int
	Initialized bool
}

type swap_Summary struct {
	Address     string
	DeltaT0     *big.Int
	DeltaT1     *big.Int
	Initialized bool
}

func PoolBuilder(price *big.Int) *LiquidityPool {
	ticks_made := make(map[int32]Tick_info)
	positions_made := make(map[string]liquidityPosition)
	positions_summary := make(map[string]position_Summary)
	swaps_summary := make(map[string]swap_Summary)
	lp := new(LiquidityPool)
	lp.Positions = positions_made
	lp.PositionSummaries = positions_summary
	lp.SwapSummaries = swaps_summary
	lp.ticks = ticks_made
	lp.tick_bitmap = make(map[int16]*big.Int)
	lp.tick_spacing = int32(lib.TickSpacings[lib.FeeMedium])
	lp.fee = uint64(lib.FeeMedium)
	lp.MAX_TICK = lib.MaxTick
	lp.MIN_TICK = lib.MinTick
	//---------------------------------------------------//
	lp.Price = price
	lp.Token0TotalLiquidity = big.NewInt(0)
	lp.Token1TotalLiquidity = big.NewInt(0)
	int_tick, _ := lib.GetTickAtSqrtRatio(lp.Price)
	lp.Current_tick = int32(int_tick)
	lp.feeGrowthGlobal0X128 = big.NewInt(0)
	lp.feeGrowthGlobal1X128 = big.NewInt(0)
	lp.Liquidity = big.NewInt(0)
	modifiedPositions := make([]string, 0)
	lp.ModifiedPositions = modifiedPositions
	return lp
}

func get_tick_pos(tick int32) (int16, uint8) {
	return int16(tick >> 8), uint8(tick % 256)
}

func (lp *LiquidityPool) flip_tick(tick, tick_spacing int32) error {
	// Check if tick % tick_spacing is zero
	if tick%tick_spacing != 0 {
		fmt.Printf("tick is not divisible by tick_spacing")
		return nil
	}
	// Get tick position
	word_pos, bit_pos := get_tick_pos(tick / tick_spacing)
	// Create a new big.Int named mask containing 2^bit_pos
	mask := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bit_pos)), nil)
	// XOR the value stored at tick_bitmap[word_pos] with mask
	if lp.tick_bitmap[word_pos] == nil {
		lp.tick_bitmap[word_pos] = big.NewInt(0)
	}
	copy := lp.tick_bitmap[word_pos]
	copy.Xor(copy, mask)
	lp.tick_bitmap[word_pos] = copy
	return nil
}

func (lp *LiquidityPool) nextInitializedTickWithinOneWord(tick int32, lte bool) (next int32, initialzied bool) {
	keyList := maps.Keys(lp.ticks)
	if _, ok := lp.ticks[tick]; !ok {
		keyList = append(keyList, tick)
	}
	sort.Slice(keyList, func(i, j int) bool { return keyList[i] < keyList[j] })
	indexCurrentTick := slices.Index(keyList, tick)
	if lte {
		if _, ok := lp.ticks[tick]; ok {
			return tick, true
		} else if indexCurrentTick == 0 {
			return lib.MinTick, false
		} else {
			next = keyList[indexCurrentTick-1]
		}

	} else {
		if indexCurrentTick == len(keyList)-1 {
			return lib.MaxTick, false
		}

		next = keyList[indexCurrentTick+1]
	}
	return next, true
}

func (lp *LiquidityPool) update_tick(tick, currentTick int32, liquidityDelta, feeGrowthGlobal0X128, feeGrowthGlobal1X128 *big.Int, upper bool) bool {
	tick_info := lp.ticks[tick]
	liquidityBefore := new(big.Int)
	liquidityNet := new(big.Int)

	if !tick_info.Initialized {
		tick_info.LiquidityGross = big.NewInt(0)
		tick_info.LiquidityNet = big.NewInt(0)
		tick_info.feeGrowthOutside0X128 = big.NewInt(0)
		tick_info.feeGrowthOutside1X128 = big.NewInt(0)
	}

	liquidityBefore.Set(tick_info.LiquidityGross)
	liquidityAfter := lib.AddLiquidity(liquidityBefore, liquidityDelta)

	flipped := (liquidityAfter.Cmp(big.NewInt(0)) == 0) != (liquidityBefore.Cmp(big.NewInt(0)) == 0)

	if liquidityBefore.Cmp(big.NewInt(0)) == 0 {
		if tick <= currentTick {
			tick_info.feeGrowthOutside0X128.Set(feeGrowthGlobal0X128)
			tick_info.feeGrowthOutside1X128.Set(feeGrowthGlobal1X128)
		}
		tick_info.Initialized = true
	}
	tick_info.LiquidityGross = liquidityAfter
	tmpliq := new(big.Int)
	if upper {
		tmpliq.Sub(liquidityNet, liquidityDelta)
	} else {
		tmpliq.Add(liquidityNet, liquidityDelta)
	}

	tick_info.LiquidityNet = tmpliq
	lp.ticks[tick] = tick_info
	//fmt.Println("FGO: ", tick_info.feeGrowthOutside0X128)
	return flipped
}

func (lp *LiquidityPool) GetFeeGrowthInside(lowerTick, upperTick, currentTick int32, feeGrowthGlobal0X128, feeGrowthGlobal1X128 *big.Int) (*big.Int, *big.Int) {
	tick_info_low := lp.ticks[lowerTick]
	tick_info_upper := lp.ticks[upperTick]

	if !tick_info_low.Initialized {
		tick_info_low.feeGrowthOutside0X128 = big.NewInt(0)
		tick_info_low.feeGrowthOutside1X128 = big.NewInt(0)
	}
	if !tick_info_upper.Initialized {
		tick_info_upper.feeGrowthOutside0X128 = big.NewInt(0)
		tick_info_upper.feeGrowthOutside1X128 = big.NewInt(0)
	}

	feeGrowthBelow0X128 := new(big.Int)
	feeGrowthBelow1X128 := new(big.Int)
	if currentTick >= lowerTick {
		feeGrowthBelow0X128.Set(tick_info_low.feeGrowthOutside0X128)
		feeGrowthBelow1X128.Set(tick_info_low.feeGrowthOutside1X128)
	} else {
		feeGrowthBelow0X128.Sub(feeGrowthGlobal0X128, tick_info_low.feeGrowthOutside0X128)
		feeGrowthBelow1X128.Sub(feeGrowthGlobal1X128, tick_info_low.feeGrowthOutside1X128)
	}
	feeGrowthAbove0X128 := new(big.Int)
	feeGrowthAbove1X128 := new(big.Int)
	if currentTick < upperTick {
		feeGrowthAbove0X128.Set(tick_info_upper.feeGrowthOutside0X128)
		feeGrowthAbove1X128.Set(tick_info_upper.feeGrowthOutside1X128)
	} else {
		feeGrowthAbove0X128.Sub(feeGrowthGlobal0X128, tick_info_upper.feeGrowthOutside0X128)
		feeGrowthAbove1X128.Sub(feeGrowthGlobal1X128, tick_info_upper.feeGrowthOutside1X128)
	}
	feeGrowthInside0X128 := new(big.Int)
	feeGrowthInside0X128.Sub(feeGrowthGlobal0X128, feeGrowthBelow0X128)
	feeGrowthInside0X128.Sub(feeGrowthInside0X128, feeGrowthAbove0X128)
	feeGrowthInside1X128 := new(big.Int)
	feeGrowthInside1X128.Sub(feeGrowthGlobal1X128, feeGrowthBelow1X128)
	feeGrowthInside1X128.Sub(feeGrowthInside1X128, feeGrowthAbove1X128)

	return feeGrowthInside0X128, feeGrowthInside1X128
}

func (lp *LiquidityPool) GetPosition(address string, lowerTick, upperTick int32) (liquidityPosition, string) {
	lowerTick_str := strconv.FormatInt(int64(lowerTick), 10)
	upperTick_str := strconv.FormatInt(int64(upperTick), 10)
	index := lowerTick_str + address + upperTick_str
	pos := lp.Positions[index]

	if !pos.Initialized {
		pos.ID = index
		pos.LowerBound = lowerTick
		pos.UpperBound = upperTick
		pos.feeGrowthInside0LastX128 = big.NewInt(0)
		pos.feeGrowthInside1LastX128 = big.NewInt(0)
		pos.Token0Owed = big.NewInt(0)
		pos.Token1Owed = big.NewInt(0)
		pos.Token0Fees = big.NewInt(0)
		pos.Token1Fees = big.NewInt(0)
		pos.LiquidAmount = big.NewInt(0)
		pos.Initialized = true
	}
	return pos, index
}

func (lp *LiquidityPool) cross(tick int32, feeGrowthGlobal0X128, feeGrowthGlobal1X128 *big.Int) (liquidityNet *big.Int) {
	tick_info := lp.ticks[tick]
	tick_info.feeGrowthOutside0X128.Sub(feeGrowthGlobal0X128, tick_info.feeGrowthOutside0X128)
	tick_info.feeGrowthOutside1X128.Sub(feeGrowthGlobal1X128, tick_info.feeGrowthOutside1X128)
	liquidityNet = tick_info.LiquidityNet
	lp.ticks[tick] = tick_info
	return
}

func (pos *liquidityPosition) UpdatePosition(liquidityDelta, feeGrowthInside0X128, feeGrowthInside1X128 *big.Int) (liquidityPosition, *big.Int, *big.Int) {
	a := new(big.Int)
	a.Sub(feeGrowthInside0X128, pos.feeGrowthInside0LastX128)

	fmt.Println("feeGrowthInside0X128", feeGrowthInside0X128)
	fmt.Println("pos.feeGrowthInside0LastX128", pos.feeGrowthInside0LastX128)
	fmt.Println("feeGrowthInside1X128", feeGrowthInside1X128)
	fmt.Println("pos.feeGrowthInside1LastX128", pos.feeGrowthInside1LastX128)

	token0Owed := lib.MulDivRoundingDown(a, pos.LiquidAmount, lib.MaxQ128)
	b := new(big.Int)
	b.Sub(feeGrowthInside1X128, pos.feeGrowthInside1LastX128)
	token1Owed := lib.MulDivRoundingDown(b, pos.LiquidAmount, lib.MaxQ128)

	pos.LiquidAmount = lib.AddLiquidity(pos.LiquidAmount, liquidityDelta)
	pos.feeGrowthInside0LastX128 = feeGrowthInside0X128
	pos.feeGrowthInside1LastX128 = feeGrowthInside1X128
	if token0Owed.Cmp(big.NewInt(0)) == 1 || token1Owed.Cmp(big.NewInt(0)) == 1 {
		pos.Token0Owed.Add(pos.Token0Owed, token0Owed)
		pos.Token1Owed.Add(pos.Token1Owed, token1Owed)
		// needs reset at end of epoch
		pos.Token0Fees.Add(pos.Token0Fees, token0Owed)
		pos.Token1Fees.Add(pos.Token1Fees, token1Owed)
	}
	return *pos, token0Owed, token1Owed
}

func (lp *LiquidityPool) MintInterface(addr_owner string, lowerBound, upperBound int32, token0Amount, token1Amount *big.Int) (*big.Int, *big.Int) {

	sqrtPriceLowerX96, err := lib.GetSqrtRatioAtTick(int(lowerBound))
	if err != nil {
		fmt.Println(err)
	}
	sqrtPriceUpperX96, err := lib.GetSqrtRatioAtTick(int(upperBound))
	if err != nil {
		fmt.Println(err)
	}
	sqrtPriceCurrentX96, err := lib.GetSqrtRatioAtTick(int(lp.Current_tick))
	if err != nil {
		fmt.Println(err)
	}

	liquidity := lib.MaxLiquidityForAmounts(sqrtPriceCurrentX96, sqrtPriceLowerX96, sqrtPriceUpperX96, token0Amount, token1Amount, true)
	//fmt.Println("\nIn mint_interface, passing main price value", sqrtPriceCurrentX96)
	//fmt.Println("\nIn mint_interface, passing sqrtPriceLow value", sqrtPriceLowerX96)
	//fmt.Println("\nIn mint_interface, passing sqrtPriceHi value", sqrtPriceUpperX96)
	//fmt.Println("\nIn mint_interface, passing token0 value", token0Amount)
	//fmt.Println("\nIn mint_interface, passing token1 value", token1Amount)
	//fmt.Println("\nIn mint_interface, passing liquidity value", liquidity)
	amount0, amount1 := lp.Mint(addr_owner, lowerBound, upperBound, liquidity)
	return amount0, amount1
}

func (lp *LiquidityPool) ModifyPosition(addr_owner string, lowerTick, upperTick int32, liquidityDelta *big.Int) (*big.Int, *big.Int) {

	position, index := lp.GetPosition(addr_owner, lowerTick, upperTick)
	flipped_lower := lp.update_tick(lowerTick, lp.Current_tick, liquidityDelta, lp.feeGrowthGlobal0X128, lp.feeGrowthGlobal1X128, false)
	flipped_upper := lp.update_tick(upperTick, lp.Current_tick, liquidityDelta, lp.feeGrowthGlobal0X128, lp.feeGrowthGlobal1X128, true)

	if flipped_lower {
		lp.flip_tick(lowerTick, lp.tick_spacing)
	}
	if flipped_upper {
		lp.flip_tick(upperTick, lp.tick_spacing)
	}

	feeGrowthInside0X128, feeGrowthInside1X128 := lp.GetFeeGrowthInside(lowerTick, upperTick, lp.Current_tick, lp.feeGrowthGlobal0X128, lp.feeGrowthGlobal1X128)

	pos, token0FeeGrowth, token1FeeGrowth := position.UpdatePosition(liquidityDelta, feeGrowthInside0X128, feeGrowthInside1X128)

	lp.Positions[index] = pos
	// create a list of modified positions so that we can reset earned fee counter at the end of the epoch, only add UID if it is not already in list
	if !lp.member(index) {
		lp.ModifiedPositions = append(lp.ModifiedPositions, index)
	}
	// update the summary for this position with the fee growth of each token
	lp.ModifyPosSummary(addr_owner, lowerTick, upperTick, token0FeeGrowth, token1FeeGrowth, lib.Zero, lib.Zero)

	amount0 := new(big.Int)
	amount1 := new(big.Int)
	sqrt_low, _ := lib.GetSqrtRatioAtTick(int(lowerTick))
	sqrt_hi, _ := lib.GetSqrtRatioAtTick(int(upperTick))

	if lp.Current_tick < lowerTick {
		amount0 = lib.GetAmount0Delta_noBool(sqrt_low, sqrt_hi, liquidityDelta)
	} else if lp.Current_tick < upperTick {
		amount0 = lib.GetAmount0Delta_noBool(lp.Price, sqrt_hi, liquidityDelta)
		amount1 = lib.GetAmount1Delta_noBool(sqrt_low, lp.Price, liquidityDelta)
		lp.Liquidity = lib.AddLiquidity(lp.Liquidity, liquidityDelta)
	} else {
		amount1 = lib.GetAmount1Delta_noBool(sqrt_low, sqrt_hi, liquidityDelta)
	}
	return amount0, amount1
}

func (lp *LiquidityPool) Mint(addr_owner string, lowerTick, upperTick int32, amount *big.Int) (*big.Int, *big.Int) {
	if (lowerTick >= upperTick) || (lowerTick < lp.MIN_TICK) || (upperTick > lp.MAX_TICK) {
		fmt.Printf("invalid ticks @mint func \n")
		return big.NewInt(0), big.NewInt(0)
	}
	if amount.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("Cannot mint 0 amount \n")
		return big.NewInt(0), big.NewInt(0)
	}

	amount0, amount1 := lp.ModifyPosition(addr_owner, lowerTick, upperTick, amount)
	balance0Before := new(big.Int)
	balance1Before := new(big.Int)

	if amount0.Cmp(big.NewInt(0)) > 0 {
		balance0Before = lp.Token0TotalLiquidity
	}
	if amount1.Cmp(big.NewInt(0)) > 0 {
		balance1Before = lp.Token1TotalLiquidity
	}

	sum := new(big.Int)
	sum.Add(balance0Before, amount0)

	if (amount0.Cmp(big.NewInt(0)) > 0) && (sum.Cmp(lp.Token0TotalLiquidity) < 1) {
		fmt.Printf("Insufficient Input Funds amount 0")
	}
	sum2 := new(big.Int)
	sum2.Add(balance1Before, amount1)
	if (amount1.Cmp(big.NewInt(0)) > 0) && (sum2.Cmp(lp.Token1TotalLiquidity) < 1) {
		fmt.Printf("Insufficient Input Funds amount 1")
	}
	lp.Token0TotalLiquidity.Add(lp.Token0TotalLiquidity, amount0)
	lp.Token1TotalLiquidity.Add(lp.Token1TotalLiquidity, amount1)
	lp.ModifyPosSummary(addr_owner, lowerTick, upperTick, lib.Zero, lib.Zero, amount0, amount1)
	return amount0, amount1
}

func (lp *LiquidityPool) BurnInterface(addr_owner string, lowerBound, upperBound int32, token0Amount, token1Amount *big.Int) (*big.Int, *big.Int) {

	sqrtPriceLowerX96, err := lib.GetSqrtRatioAtTick(int(lowerBound))
	if err != nil {
		fmt.Println(err)
	}
	sqrtPriceUpperX96, err := lib.GetSqrtRatioAtTick(int(upperBound))
	if err != nil {
		fmt.Println(err)
	}
	sqrtPriceCurrentX96, err := lib.GetSqrtRatioAtTick(int(lp.Current_tick))
	if err != nil {
		fmt.Println(err)
	}

	liquidity := lib.MaxLiquidityForAmounts(sqrtPriceCurrentX96, sqrtPriceLowerX96, sqrtPriceUpperX96, token0Amount, token1Amount, true)
	amount0, amount1 := lp.Burn(addr_owner, lowerBound, upperBound, liquidity)
	return amount0, amount1
}

func (lp *LiquidityPool) Burn(addr_owner string, lowerTick, upperTick int32, amount *big.Int) (*big.Int, *big.Int) {
	position, index := lp.GetPosition(addr_owner, lowerTick, upperTick)

	if (position.LiquidAmount.Cmp(big.NewInt(0)) == 0) && (position.Token0Owed.Cmp(big.NewInt(0)) == 0) && (position.Token1Owed.Cmp(big.NewInt(0)) == 0) {
		delete(lp.Positions, index)
		return big.NewInt(-1), big.NewInt(-1)
	}

	if !(position.LiquidAmount.Cmp(amount) >= 0) {
		amount = position.LiquidAmount
	}
	neg_amount := new(big.Int).Neg(amount)
	amount0, amount1 := lp.ModifyPosition(addr_owner, lowerTick, upperTick, neg_amount)
	position, index = lp.GetPosition(addr_owner, lowerTick, upperTick)
	amount0_neg := new(big.Int).Neg(amount0)
	amount1_neg := new(big.Int).Neg(amount1)
	if amount0_neg.Cmp(big.NewInt(0)) > 0 || amount1_neg.Cmp(big.NewInt(0)) > 0 {
		sum0 := new(big.Int).Add(position.Token0Owed, amount0_neg)
		sum1 := new(big.Int).Add(position.Token1Owed, amount1_neg)
		position.Token0Owed.Set(sum0)
		position.Token1Owed.Set(sum1)
		lp.Positions[index] = position
	}
	lp.Token0TotalLiquidity.Add(lp.Token0TotalLiquidity, amount0)
	lp.Token1TotalLiquidity.Add(lp.Token1TotalLiquidity, amount1)
	lp.ModifyPosSummary(addr_owner, lowerTick, upperTick, lib.Zero, lib.Zero, amount0, amount1)
	return amount0, amount1
}

func (lp *LiquidityPool) Collect(addr_owner string, tickLower, tickUpper int32, amount0Requested, amount1Requested *big.Int) (*big.Int, *big.Int) {
	position, index := lp.GetPosition(addr_owner, tickLower, tickUpper)
	amount0 := new(big.Int)
	amount1 := new(big.Int)
	if amount0Requested.Cmp(position.Token0Owed) > 0 {
		amount0.Set(position.Token0Owed)
	} else {
		amount0.Set(amount0Requested)
	}
	if amount1Requested.Cmp(position.Token1Owed) > 0 {
		amount1.Set(position.Token1Owed)
	} else {
		amount1.Set(amount1Requested)
	}

	feesEarnedT0 := new(big.Int).Set(position.Token0Fees)
	feesEarnedT1 := new(big.Int).Set(position.Token1Fees)
	fee0Collected := new(big.Int)
	fee1Collected := new(big.Int)
	fmt.Println("fee0Collected:", fee0Collected)
	if feesEarnedT0.Cmp(amount0Requested) <= 0 {
		fee0Collected.Set(feesEarnedT0)
	} else {
		fee0Collected.Set(amount0Requested)
	}
	if feesEarnedT1.Cmp(amount1Requested) <= 0 {
		fee1Collected.Set(feesEarnedT1)
	} else {
		fee1Collected.Set(amount1Requested)
	}
	fee0Collected.Neg(fee0Collected)
	fee1Collected.Neg(fee1Collected)
	fmt.Println("fee0Collected.Neg:", fee0Collected)
	position.Token0Fees.Add(position.Token0Fees, fee0Collected)
	position.Token1Fees.Add(position.Token1Fees, fee1Collected)
	lp.Positions[index] = position
	lp.ModifyPosSummary(addr_owner, tickLower, tickUpper, fee0Collected, fee1Collected, lib.Zero, lib.Zero)

	if amount0.Cmp(big.NewInt(0)) > 0 {
		position.Token0Owed.Sub(position.Token0Owed, amount0)
		lp.Token0TotalLiquidity.Sub(lp.Token0TotalLiquidity, amount0)
	}
	if amount1.Cmp(big.NewInt(0)) > 0 {
		position.Token1Owed.Sub(position.Token1Owed, amount1)
		lp.Token1TotalLiquidity.Sub(lp.Token1TotalLiquidity, amount1)
	}
	return amount0, amount1
}

func (lp *LiquidityPool) Swap(addr_owner string, zeroForOne bool, amountSpecified, sqrtPriceLimitX96 *big.Int) (addr string, amount0, amount1 *big.Int) {
	liquidity_ := lp.Liquidity

	if zeroForOne {
		if (sqrtPriceLimitX96.Cmp(lp.Price) == 1) || (sqrtPriceLimitX96.Cmp(lib.MinSqrtRatio) == -1) {
			// REVERT INVALID PRICE LIMIT
			fmt.Println("Invalid price limit")
		}
	} else {
		if (sqrtPriceLimitX96.Cmp(lp.Price) == -1) || (sqrtPriceLimitX96.Cmp(lib.MaxSqrtRatio) == 1) {
			// REVERT INVALID PRICE LIMIT
			fmt.Println("Invalid price limit")
		}
	}
	// create new swap state to track the progress of execution
	state := new(SwapState)
	state.AmountSpecifiedRemaining = new(big.Int).Set(amountSpecified)
	state.AmountCalculated = big.NewInt(0)
	state.SqrtPriceX96 = new(big.Int).Set(lp.Price)
	state.Tick = lp.Current_tick
	state.liquidity = new(big.Int).Set(liquidity_)

	exactInput := amountSpecified.Cmp(lib.Zero) > 0

	if zeroForOne {
		state.feeGrowthGlobalX128 = new(big.Int).Set(lp.feeGrowthGlobal0X128)
	} else {
		state.feeGrowthGlobalX128 = new(big.Int).Set(lp.feeGrowthGlobal1X128)
	}

	for (state.AmountSpecifiedRemaining.Cmp(big.NewInt(0)) != 0) && (state.SqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0) {

		step := new(StepState)
		step.SqrtPriceStartX96 = new(big.Int).Set(state.SqrtPriceX96)
		step.NextTick, step.Initialized = lp.nextInitializedTickWithinOneWord(state.Tick, zeroForOne)
		step.SqrtPriceNextX96, _ = lib.GetSqrtRatioAtTick(int(step.NextTick))

		sqrtRatioTarget := new(big.Int)
		if zeroForOne {
			if step.SqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0 {
				sqrtRatioTarget = new(big.Int).Set(sqrtPriceLimitX96)
			} else {
				sqrtRatioTarget = new(big.Int).Set(step.SqrtPriceNextX96)
			}
		} else {
			if step.SqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
				sqrtRatioTarget = new(big.Int).Set(sqrtPriceLimitX96)
			} else {
				sqrtRatioTarget = new(big.Int).Set(step.SqrtPriceNextX96)
			}
		}

		//fmt.Println("priceNow: ", lib.FormatPrice(step.SqrtPriceStartX96), lib.FormatPrice(step.SqrtPriceNextX96))
		//fmt.Println("State sqrtPrice: ", lib.FormatPrice(state.SqrtPriceX96))
		//fmt.Println("target: ", sqrtRatioTarget)
		//fmt.Println("State liqudiity: ", state.liquidity)
		//fmt.Println("amount spec remain: ", state.AmountSpecifiedRemaining)
		state.SqrtPriceX96, step.AmountIn, step.AmountOut, step.feeAmount, _ = lib.ComputeSwapStep(state.SqrtPriceX96, sqrtRatioTarget, state.liquidity, state.AmountSpecifiedRemaining, lib.FeeAmount(lp.fee))
		//fmt.Println("priceNow: ", lib.FormatPrice(state.SqrtPriceX96), lib.FormatPrice(step.SqrtPriceNextX96))
		tmp_amount_spec_remain := new(big.Int)
		tmp_amount_spec_remain.Add(step.AmountIn, step.feeAmount)

		if exactInput {
			state.AmountSpecifiedRemaining.Sub(state.AmountSpecifiedRemaining, tmp_amount_spec_remain)
			state.AmountCalculated.Add(state.AmountCalculated, step.AmountOut)
		} else {
			state.AmountSpecifiedRemaining.Add(state.AmountSpecifiedRemaining, step.AmountOut)
			state.AmountCalculated.Add(state.AmountCalculated, step.AmountOut)
		}

		if state.liquidity.Cmp(lib.Zero) > 0 {
			mulDivOut := new(big.Int)
			fmt.Println("Liquidity Is Not Zero: but: ", state.liquidity)
			mulDivOut = lib.MulDivRoundingDown(step.feeAmount, lib.MaxQ128, state.liquidity)
			state.feeGrowthGlobalX128.Add(state.feeGrowthGlobalX128, mulDivOut)
		}

		if state.SqrtPriceX96.Cmp(step.SqrtPriceNextX96) == 0 {
			if step.Initialized {
				arg1 := new(big.Int)
				arg2 := new(big.Int)
				if zeroForOne {
					arg1 = state.feeGrowthGlobalX128
					arg2 = lp.feeGrowthGlobal1X128
				} else {
					arg1 = lp.feeGrowthGlobal0X128
					arg2 = state.feeGrowthGlobalX128
				}
				//fmt.Println("step.NextTick: ", step.NextTick)
				//fmt.Println("arg1: ", arg1)
				//fmt.Println("arg2: ", arg2)
				liquidityDelta := lp.cross(step.NextTick, arg1, arg2)
				//fmt.Println("LiquidityDelta: ", liquidityDelta)
				if zeroForOne {
					liquidityDelta.Neg(liquidityDelta)
				}
				state.liquidity = lib.AddLiquidity(state.liquidity, liquidityDelta)
			}
			if state.liquidity.Cmp(lib.Zero) == 0 {
				// REVERT NOT ENOUGH LIQUIDITY
				fmt.Println("REVERT NOT ENOUGH LIQUIDITY DURING SWAP PROCESS")
			}
			if zeroForOne {
				state.Tick = step.NextTick - 1
				fmt.Println("new state tick", state.Tick)
			} else {
				state.Tick = step.NextTick
			}
		} else if state.SqrtPriceX96.Cmp(step.SqrtPriceStartX96) != 0 {
			out, _ := lib.GetTickAtSqrtRatio(state.SqrtPriceX96)
			state.Tick = int32(out)
		}
	}
	if state.Tick != lp.Current_tick {
		lp.Price = state.SqrtPriceX96
		lp.Current_tick = state.Tick
	} else {
		lp.Price = state.SqrtPriceX96
	}
	if liquidity_.Cmp(state.liquidity) != 0 {
		lp.Liquidity = state.liquidity
	}
	if zeroForOne {
		lp.feeGrowthGlobal0X128 = state.feeGrowthGlobalX128
	} else {
		lp.feeGrowthGlobal1X128 = state.feeGrowthGlobalX128
	}
	amount0 = new(big.Int)
	amount1 = new(big.Int)
	if zeroForOne {
		amount0.Sub(amountSpecified, state.AmountSpecifiedRemaining)
		amount1.Neg(state.AmountCalculated)
	} else {
		amount0.Neg(state.AmountCalculated)
		amount1.Sub(amountSpecified, state.AmountSpecifiedRemaining)
	}
	lp.Token0TotalLiquidity.Add(lp.Token0TotalLiquidity, amount0)
	lp.Token1TotalLiquidity.Add(lp.Token1TotalLiquidity, amount1)
	lp.ModifySwapSummary(addr_owner, amount0, amount1)
	return addr_owner, amount0, amount1
}

func (lp *LiquidityPool) ModifyPosSummary(address string, lowerTick, upperTick int32, feeChangeT0, feeChangeT1, deltaT0, deltaT1 *big.Int) {
	lowerTick_str := strconv.FormatInt(int64(lowerTick), 10)
	upperTick_str := strconv.FormatInt(int64(upperTick), 10)
	index := lowerTick_str + address + upperTick_str
	summary := lp.PositionSummaries[index]
	if !summary.Initialized {
		summary.ID = address
		summary.LowerBound = lowerTick
		summary.UpperBound = upperTick
		summary.FeeChangeT0 = new(big.Int).Set(feeChangeT0)
		summary.FeeChangeT1 = new(big.Int).Set(feeChangeT1)
		summary.DeltaT0 = new(big.Int).Set(deltaT0)
		summary.DeltaT1 = new(big.Int).Set(deltaT1)
		summary.Initialized = true
	} else {
		summary.FeeChangeT0.Add(summary.FeeChangeT0, feeChangeT0)
		summary.FeeChangeT1.Add(summary.FeeChangeT1, feeChangeT1)
		summary.DeltaT0.Add(summary.DeltaT0, deltaT0)
		summary.DeltaT1.Add(summary.DeltaT1, deltaT1)
	}
	lp.PositionSummaries[index] = summary
}

func (lp *LiquidityPool) ModifySwapSummary(address string, deltaT0, deltaT1 *big.Int) {
	summary := lp.SwapSummaries[address]
	if !summary.Initialized {
		summary.Address = address
		summary.DeltaT0 = new(big.Int).Set(deltaT0)
		summary.DeltaT1 = new(big.Int).Set(deltaT1)
		summary.Initialized = true
	} else {
		summary.DeltaT0.Add(summary.DeltaT0, deltaT0)
		summary.DeltaT1.Add(summary.DeltaT1, deltaT1)
	}
	lp.SwapSummaries[address] = summary
}

func (lp *LiquidityPool) GetPosSummary(address string, lowerTick, upperTick int32) position_Summary {
	lowerTick_str := strconv.FormatInt(int64(lowerTick), 10)
	upperTick_str := strconv.FormatInt(int64(upperTick), 10)
	index := lowerTick_str + address + upperTick_str
	return lp.PositionSummaries[index]
}

func (lp *LiquidityPool) CleanSummaries() {
	// funciton to call after the end of an epoch to clean up the maintained summary data so the fresh epoch can record the relevant info
	// first reset the position_summary structure
	positions_summary := make(map[string]position_Summary)
	lp.PositionSummaries = positions_summary
	// reset the swap summary structure
	swap_summary := make(map[string]swap_Summary)
	lp.SwapSummaries = swap_summary
	// reset each position's (the permanant structures) .Token0Fees .Token1Fees
	for _, index := range lp.ModifiedPositions {
		position := lp.Positions[index]
		position.Token0Fees = lib.Zero
		position.Token1Fees = lib.Zero
		lp.Positions[index] = position
	}
	// clean the list of positions which have been modified
	lp.ModifiedPositions = make([]string, 0)
}

func (lp *LiquidityPool) member(id string) (found bool) {
	found = false
	for _, str := range lp.ModifiedPositions {
		if str == id {
			found = true
			break
		}
	}
	return
}
