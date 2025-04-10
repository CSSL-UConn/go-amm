//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./quickbls.sol";

contract TokenBank {
    IERC20 public tokenA;
    IERC20 public tokenB;
    address private owner;

    uint256[4] pubkey;
    bool first_run = true;

    uint256 _liquidity;
    uint256 _sqrtPrice;

    constructor(
        address _tokenA,
        address _tokenB
    ) {
        tokenA = IERC20(_tokenA);
        tokenB = IERC20(_tokenB);
        owner = msg.sender;
    }
    // define deposit mapping 
    struct token_struct{
        int256 amount_tokenA;
        int256 amount_tokenB;
    }
    
    struct deposit_struct {
        address sepoliaAddr;
        bool initialized;
        mapping(uint256 => token_struct) deposits;
    }
    

    mapping(uint256 => deposit_struct) public deposit_map;
    // define position mapping 
    struct position_stuct {
        int256 amount_tokenA;
        int256 amount_tokenB;
        int32 lower_bound;
        int32 upper_bound;
        int256 fees_earned_A;
        int256 fees_earned_B;
    }
    mapping(address => mapping(string => position_stuct)) public position_map;
    // define sync_structure and mapping
    struct sync_struct {
        bool tx_type_id;
        uint256 sidechainAddr;
        string position_id; 
        int256 amount_tokenA;
        int256 amount_tokenB;
        int32 lower_bound;
        int32 upper_bound;
        int256 fees_earned_A;
        int256 fees_earned_B;
    }
    function deposit (int256 _amount, bool _type, uint256 epoch, uint256 sidechainAddr) public payable {
        // set minimum amounts
        uint _minAmount = 1*(10**18);
        //check if address pairing is claimed, if not claim it 
        if(deposit_map[sidechainAddr].initialized == false){
            deposit_map[sidechainAddr].sepoliaAddr = msg.sender;
            deposit_map[sidechainAddr].initialized = true;
        }
        require(uint256(_amount) >= _minAmount, "Amount less than minimum amount");
        // transfer tokens from sender to contract, conditional check for the type of token, adjust mapping
        if (_type == true) {
            tokenA.transferFrom(msg.sender, address(this), uint256(_amount));
            deposit_map[sidechainAddr].deposits[epoch].amount_tokenA += _amount;
        } else {
            tokenB.transferFrom(msg.sender, address(this), uint256(_amount));
            deposit_map[sidechainAddr].deposits[epoch].amount_tokenB += _amount;
        }
    }

    /*** 
    * 
    * Handle Sync[public]: Synchronizes the state of the sidechain at the end of the epoch.
    *
    * @input: sync_struct[] syncTx: The sync transaction entries.
    * @input: uint[4] nextPubkey: The public key used by the next committee to sign the transaction.
    * @input: uint[2] signature: the BLS signature authenticating the sync transaction
    *
    * @outputs: None 
    */
    function handle_sync (sync_struct[] calldata syncTx, uint256 epoch, uint[4] calldata nextPubkey, uint[2] memory signature, uint256 liquidity, uint256 sqrtPrice) public {
        if (first_run){
            first_run = false;
        } else {
            bytes memory encodedTx = abi.encode(syncTx);
            bool verified = QuickBLS.verifySignature(signature, pubkey, encodedTx);
            require(verified, "Invalid signature on the sync transaction");
        }
        pubkey = nextPubkey;
        liquidity = _liquidity;
        _sqrtPrice = sqrtPrice;
        for (uint i=0; i < syncTx.length; i++) {
            // for each element within the list, check the tx_type_id
            
            if (syncTx[i].tx_type_id == true) {
                uint256 _addr = syncTx[i].sidechainAddr;
                tokenA.transfer(deposit_map[_addr].sepoliaAddr, uint(syncTx[i].amount_tokenA));
                tokenB.transfer(deposit_map[_addr].sepoliaAddr, uint(syncTx[i].amount_tokenB));
                deposit_map[_addr].deposits[epoch].amount_tokenA = 0;
                deposit_map[_addr].deposits[epoch].amount_tokenB = 0;
            } else {
                // update position mapping with respective values 
                uint256 _addr = syncTx[i].sidechainAddr;
                string memory _uid = syncTx[i].position_id;
                address mcAddr = deposit_map[_addr].sepoliaAddr;
                position_map[mcAddr][_uid].amount_tokenA = syncTx[i].amount_tokenA;
                position_map[mcAddr][_uid].amount_tokenB = syncTx[i].amount_tokenB;
                position_map[mcAddr][_uid].lower_bound = syncTx[i].lower_bound;
                position_map[mcAddr][_uid].upper_bound = syncTx[i].upper_bound;
                position_map[mcAddr][_uid].fees_earned_A = syncTx[i].fees_earned_A;
                position_map[mcAddr][_uid].fees_earned_B = syncTx[i].fees_earned_B;
            }
        }
    }
    function getContractBalance() public view returns(uint tokenAval, uint tokenBval){
        return (IERC20(tokenA).balanceOf(address(this)), IERC20(tokenB).balanceOf(address(this)));
    }
    // optimization option: recieves a list of addresses which have deposited and returns all of their mappings, not just one at a time, reduce TX count overall. 
    function getDepositBalance(uint256 _addr, uint256 epoch) public view returns(int256 tokenAval, int256 tokenBval) {
        return (deposit_map[_addr].deposits[epoch].amount_tokenA, deposit_map[_addr].deposits[epoch].amount_tokenB);
    }
    // not tested yet
    function getPosition(address _addr, string memory _uid) public view returns(position_stuct memory) {
        return position_map[_addr][_uid];
    }
    function withdrawAll() public {
        require(owner == msg.sender);
        tokenA.transferFrom(msg.sender, address(this), uint256(IERC20(tokenA).balanceOf(address(this))));
        tokenB.transferFrom(msg.sender, address(this), uint256(IERC20(tokenB).balanceOf(address(this))));
    }

}