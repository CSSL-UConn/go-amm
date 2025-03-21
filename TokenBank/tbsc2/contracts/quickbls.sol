// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.20;

library QuickBLS{

    uint256 constant N = 21888242871839275222246405745257275088696311157297823662689037894645226208583;

     // This is the generator negated, to use for pairing
    uint256 public constant G2_NEG_X_RE  = 0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 public constant G2_NEG_X_IM  = 0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 public constant G2_NEG_Y_RE  = 0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 public constant G2_NEG_Y_IM  = 0x1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d;


    function htp(bytes memory message) internal view returns (uint256[2] memory r) {
        uint[3] memory input;
        bytes32  s = keccak256(message);
        input[0] = 0x1;
        input[1] = 0x2;
        input[2] = uint256(s);
        bool success;
        assembly {
            success := staticcall(sub(gas(), 2000), 7, input, 0x80, r, 0x60)
        // Use "invalid" to make gas estimation work
            switch success case 0 {invalid()}
        }
        require(success);
    }




    function verifySignature (uint256[2] memory signature,uint256[4] memory pubkey, bytes memory message) internal  returns (bool) {
       uint256[2] memory hm = htp(message);
       return verify(signature, pubkey, hm);
    }

    function verify(
    uint256[2] memory signature,
    uint256[4] memory pubkey,
    uint256[2] memory message
    ) internal  returns (bool) {

        uint256[12] memory input = [signature[0], signature[1] ,G2_NEG_X_RE , G2_NEG_X_IM, G2_NEG_Y_RE, G2_NEG_Y_IM, message[0], message[1], pubkey[0], pubkey[1], pubkey[2], pubkey[3]];
        uint256[1] memory out;
        bool success;
        return ecCheckPairing(input);
    }

   function ecCheckPairing(uint256[12] memory input) private returns (bool) {
    uint256[1] memory result;
    bool success;
    assembly {
      // 0x08     id of the bn256CheckPairing precompile    (checking the elliptic curve pairings)
      // 0        number of ether to transfer
      // 0        since we have an array of fixed length, our input starts in 0
      // 384      size of call parameters, i.e. 12*256 bits == 384 bytes
      // 32       size of result (one 32 byte boolean!)
      success := call(sub(gas(), 2000), 0x08, 0, input, 384, result, 32)
    }
    require(success, "elliptic curve pairing failed");
    return result[0] == 1;
  }

      function VerifyMsgHash(uint256[2] memory hash, bytes memory message)internal view returns (bool) {
        uint256[2] memory hm = htp(message);
        return hm[0] == hash[0] && hm[1] == hash[1];
    }
}


