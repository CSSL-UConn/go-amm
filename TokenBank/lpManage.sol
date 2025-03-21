// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity =0.7.6;
pragma abicoder v2;

import '@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol';
import '@uniswap/v3-core/contracts/libraries/TickMath.sol';
import '@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol';
import '@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol';
import '@uniswap/v3-periphery/contracts/interfaces/INonfungiblePositionManager.sol';
import '@uniswap/v3-periphery/contracts/base/LiquidityManagement.sol';
import '@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol';
import '@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol';
import '@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol';


contract lpManage is IERC721Receiver {
    address public constant ABTX = 0x208A653d61f568C19924FE0F404431b62Ba1B82f;
    address public constant ABTY = 0x65d1845521eB8029CB8777f29de5a12DAC34AFbF;
    IERC20 public tokenA;
    IERC20 public tokenB;

    uint24 public constant poolFee = 3000;

    IUniswapV3Pool public immutable pool;
    INonfungiblePositionManager public immutable nonfungiblePositionManager;
    ISwapRouter public immutable swapRouter;

    event FreshMint(uint256 tokenId,address ownerAddr);
    /// Represents the deposit of an NFT
    struct Deposit {
        address owner;
        uint128 liquidity;
        address token0;
        address token1;
    }

    ///  deposits[tokenId] => Deposit
    mapping(uint256 => Deposit) public deposits;

    constructor(
        address _tokenA,
        address _tokenB,
        INonfungiblePositionManager _nonfungiblePositionManager,
        ISwapRouter _swapRouter,
        IUniswapV3Pool _pool
    ) {
        pool = _pool;
        swapRouter = _swapRouter;
        nonfungiblePositionManager = _nonfungiblePositionManager;
        tokenA = IERC20(_tokenA);
        tokenB = IERC20(_tokenB);

    }

    // Implementing `onERC721Received` so this contract can receive custody of erc721 tokens
    function onERC721Received(
        address operator,
        address,
        uint256 tokenId,
        bytes calldata
    ) external override returns (bytes4) {
        // get position information

        _createDeposit(operator, tokenId);

        return this.onERC721Received.selector;
    }

    function _updateDeposit(uint256 tokenId, uint128 liquidity) internal {
        deposits[tokenId].liquidity += liquidity;
    }

    function _createDeposit(address owner, uint256 tokenId) internal {
        (, , address token0, address token1, , , , uint128 liquidity, , , , ) =
            nonfungiblePositionManager.positions(tokenId);

        // set the owner and data for position
        // operator is msg.sender
        deposits[tokenId] = Deposit({owner: owner, liquidity: liquidity, token0: token0, token1: token1});
    }

    /// notice Calls the mint function defined in periphery, mints the same amount of each token.
    /// return tokenId The id of the newly minted ERC721
    /// return liquidity The amount of liquidity for the position
    /// return amount0 The amount of token0
    /// return amount1 The amount of token1
    function mintNewPosition(uint256 amount0desired, uint256 amount1desired)
        external
        returns (
            uint256 tokenId,
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1
        )
    {
        // For this example, we will provide equal amounts of liquidity in both assets.
        // Providing liquidity in both assets means liquidity will be earning fees and is considered in-range.
        uint256 amount0ToMint = amount0desired;
        uint256 amount1ToMint = amount1desired;

        // transfer tokens to contract
        TransferHelper.safeTransferFrom(ABTX, msg.sender, address(this), amount0ToMint);
        TransferHelper.safeTransferFrom(ABTY, msg.sender, address(this), amount1ToMint);

        // Approve the position manager
        TransferHelper.safeApprove(ABTX, address(nonfungiblePositionManager), amount0ToMint);
        TransferHelper.safeApprove(ABTY, address(nonfungiblePositionManager), amount1ToMint);

        INonfungiblePositionManager.MintParams memory params =
            INonfungiblePositionManager.MintParams({
                token0: ABTX,
                token1: ABTY,
                fee: poolFee,
                tickLower: -887220,
                tickUpper: 887220,
                amount0Desired: amount0ToMint,
                amount1Desired: amount1ToMint,
                amount0Min: 0,
                amount1Min: 0,
                recipient: address(this),
                deadline: block.timestamp + 600
            });

        (tokenId, liquidity, amount0, amount1) = nonfungiblePositionManager.mint(params);

        // Create a deposit
        _createDeposit(msg.sender, tokenId);

        // Remove allowance and refund in both assets.
        if (amount0 < amount0ToMint) {
            TransferHelper.safeApprove(ABTX, address(nonfungiblePositionManager), 0);
            uint256 refund0 = amount0ToMint - amount0;
            TransferHelper.safeTransfer(ABTX, msg.sender, refund0);
        }

        if (amount1 < amount1ToMint) {
            TransferHelper.safeApprove(ABTY, address(nonfungiblePositionManager), 0);
            uint256 refund1 = amount1ToMint - amount1;
            TransferHelper.safeTransfer(ABTY, msg.sender, refund1);
        }
        emit FreshMint(tokenId,msg.sender);
    }

    function _sendToOwner(
        uint256 tokenId,
        uint256 amount0,
        uint256 amount1
    ) internal {
        // get owner of token
        address owner = deposits[tokenId].owner;

        // send collected fees to owner
        tokenA.transfer(owner, amount0);
        tokenB.transfer(owner, amount1);
    }

    function retrieveNFT(uint256 tokenId) external {
        // must be the owner of the NFT
        require(msg.sender == deposits[tokenId].owner, 'Not the owner');
        // transfer ownership to original owner
        nonfungiblePositionManager.safeTransferFrom(address(this), msg.sender, tokenId);
        //remove information related to tokenId
        delete deposits[tokenId];
    }

    function increaseLiquidity(uint256 tokenId, uint256 amountAdd0, uint256 amountAdd1) external returns (uint128 liquidity, uint256 amount0, uint256 amount1) {

        TransferHelper.safeTransferFrom(deposits[tokenId].token0, msg.sender, address(this), amountAdd0);
        TransferHelper.safeTransferFrom(deposits[tokenId].token1, msg.sender, address(this), amountAdd1);

        TransferHelper.safeApprove(deposits[tokenId].token0, address(nonfungiblePositionManager), amountAdd0);
        TransferHelper.safeApprove(deposits[tokenId].token1, address(nonfungiblePositionManager), amountAdd1);


        INonfungiblePositionManager.IncreaseLiquidityParams memory params =
            INonfungiblePositionManager.IncreaseLiquidityParams({
                tokenId: tokenId,
                amount0Desired: amountAdd0,
                amount1Desired: amountAdd1,
                amount0Min: 0,
                amount1Min: 0,
                deadline: block.timestamp + 600
            });

        (liquidity, amount0, amount1) = nonfungiblePositionManager.increaseLiquidity(params);
        // UPDATE DEPOSIT WHEN INCREASE
        _updateDeposit(tokenId, liquidity);

    }

    function decreaseLiquidity(uint256 tokenId, uint256 token0Amount, uint256 token1Amount,uint160 lower,uint160 upper) external returns (uint256 amount0, uint256 amount1) {
        // caller must be the owner of the NFT
        require(msg.sender == deposits[tokenId].owner, 'Not the owner');
        // get liquidity data for tokenId
        uint128 currentLiquidity = deposits[tokenId].liquidity;

        (uint160 sqrtPriceX96,,,,,,) = pool.slot0();

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(sqrtPriceX96,lower,upper,token0Amount,token1Amount);

        uint128 liquidityDelta = liquidity - currentLiquidity;

        // amount0Min and amount1Min are price slippage checks
        // if the amount received after burning is not greater than these minimums, transaction will fail
        INonfungiblePositionManager.DecreaseLiquidityParams memory params =
            INonfungiblePositionManager.DecreaseLiquidityParams({
                tokenId: tokenId,
                liquidity: liquidity,
                amount0Min: 0,
                amount1Min: 0,
                deadline: block.timestamp + 600
            });

        (amount0, amount1) = nonfungiblePositionManager.decreaseLiquidity(params);

        _updateDeposit(tokenId, liquidityDelta);

        //send liquidity back to owner
        //_sendToOwner(tokenId, amount0, amount1);
        
    }

    function collect(uint256 tokenId) external returns (uint256 amount0, uint256 amount1) {
        // Caller must own the ERC721 position
        // Call to safeTransfer will trigger `onERC721Received` which must return the selector else transfer will fail
        //nonfungiblePositionManager.safeTransferFrom(msg.sender, address(this), tokenId);

        // set amount0Max and amount1Max to uint256.max to collect all fees
        // alternatively can set recipient to msg.sender and avoid another transaction in `sendToOwner`
        INonfungiblePositionManager.CollectParams memory params =
            INonfungiblePositionManager.CollectParams({
                tokenId: tokenId,
                recipient: address(this),
                amount0Max: type(uint128).max,
                amount1Max: type(uint128).max
            });

        (amount0, amount1) = nonfungiblePositionManager.collect(params);

        // send collected feed back to owner
        _sendToOwner(tokenId, amount0, amount1);
    }

    function swap(uint256 amountIn, bool zeroForOne) external returns (uint256 amountOut) {
        // msg.sender must approve this contract

        if(zeroForOne){
            // Transfer the specified amount of DAI to this contract.
            TransferHelper.safeTransferFrom(ABTX, msg.sender, address(this), amountIn);
            // Approve the router to spend DAI.
            TransferHelper.safeApprove(ABTX, address(swapRouter), amountIn);
            ISwapRouter.ExactInputSingleParams memory params =
                ISwapRouter.ExactInputSingleParams({
                    tokenIn: ABTX,
                    tokenOut: ABTY,
                    fee: poolFee,
                    recipient: msg.sender,
                    deadline: block.timestamp + 600,
                    amountIn: amountIn,
                    amountOutMinimum: 0,
                    sqrtPriceLimitX96: 0
                });
            // The call to `exactInputSingle` executes the swap.
            amountOut = swapRouter.exactInputSingle(params);
        } else {
            // Transfer the specified amount of DAI to this contract.
            TransferHelper.safeTransferFrom(ABTY, msg.sender, address(this), amountIn);
            // Approve the router to spend DAI.
            TransferHelper.safeApprove(ABTY, address(swapRouter), amountIn);
            ISwapRouter.ExactInputSingleParams memory params =
                ISwapRouter.ExactInputSingleParams({
                    tokenIn: ABTY,
                    tokenOut: ABTX,
                    fee: poolFee,
                    recipient: msg.sender,
                    deadline: block.timestamp + 600,
                    amountIn: amountIn,
                    amountOutMinimum: 0,
                    sqrtPriceLimitX96: 0
                });
            // The call to `exactInputSingle` executes the swap.
            amountOut = swapRouter.exactInputSingle(params);
        }
      
    }

}