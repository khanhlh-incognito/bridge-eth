pragma solidity >=0.4.21 <0.6.0;

contract CurveP256 {
  uint256 constant a = 0xFFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC;
  uint256 constant b = 41058363725152142129326129780047268409114441015993725554835256314039467401291;
  uint256 constant P = 115792089210356248762697446949407573530086143415290314195533631308867097853951;
  uint256 constant PAdd1Div4 = 28948022302589062190674361737351893382521535853822578548883407827216774463488;
  uint256 constant N = 115792089210356248762697446949407573529996955224135760342422259061068512044369;
  uint256 constant Gx = 0x6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296;
  uint256 constant Gy = 0x4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5;
  function isOnCurve(uint256 x, uint256 y)
  public pure returns(bool){
    uint256 xCube = mulmod(x, y, P);
    // Calculate y^2 = x^3 + ax + b (mod P)
    xCube = mulmod(xCube, x, P);
    xCube = addmod(xCube, mulmod(x, a, P), P);
    xCube = addmod(xCube, b, P);
    uint256 yExp2 = mulmod(y, y, P);
    return xCube == yExp2;
  }
  function compressPoint(uint256 x, uint256 y)
  private pure returns(bytes memory){
    bytes1 format = 0x02;
    if (y&1 == 1) format |= 0x01;
    return abi.encodePacked(format, bytes32(x));
  }
  function chkPointAndCompPoint(uint256 x, uint256 y, bytes memory compPoint)
  public pure returns(bool){
    require(compPoint.length==33, "Invalid len of compressed point");
    require(compPoint[0]==0x02 || compPoint[0]==0x03, "Invalid compressed point");
    bool yBit = (compPoint[0] & 0x01) == 0x01;
    require(((y & 1) == 1) == yBit, "Invalid");
    return isOnCurve(x,y);
  }
  function invModP(uint256 x)
  public pure returns (uint256){
    return powModP(x,P-2);
  }
  function powModP(uint256 x, uint256 k)
  public pure returns (uint256){
    if (k == 0) return 1;
    uint256 p = powModP(x, k/2) % P;
    p = (p * p) % P;
    return (k%2 == 0)? p : (x * p) % P;
  }

  function addPoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2)
  public pure returns (uint256 x, uint256 y) {
    if ((x2==0) && (y2==0)) return (x1,y1);
    if ((x1==0) && (y1==0)) return (x2,y2);
    if (x1==x2) {
      if (y1==y2) return dblPoint(x1,y1);
      return (0,0);
    }
    uint256 m = mulmod(addmod(y2,P-y1,P), invModP(addmod(x2, P-x1, P)), P);
    uint256 xRes = addmod(P-x1, P-x2, P);
    xRes = addmod(xRes, mulmod(m, m, P), P);
    uint256 yRes = mulmod(m, addmod(x1, P-xRes, P), P);
    yRes = addmod(yRes, y1, P);
  }

  function dblPoint(uint256 x, uint256 y)
  public pure returns(uint256, uint256){
    uint256 t1 = mulmod(3, mulmod(x, x, P), P);
    t1 = addmod(t1,b,P);
    uint256 dbly = mulmod(y, 2, P);
    uint256 m = mulmod(t1, invModP(dbly), P);
    uint256 t1Exp2 = mulmod(t1, t1, P);
    uint256 dblyExp2 = mulmod(dbly, dbly, P);
    uint256 xRes = P-mulmod(2, x, P);
    xRes = addmod(xRes, mulmod(t1Exp2, invModP(dblyExp2), P), P);
    uint256 yRes = mulmod(m, addmod(x, P-xRes, P), P);
    yRes = addmod(yRes, P-y, P);
    return (xRes, yRes);
  }
  function scalarMul(uint256 x, uint256 y, uint256 k)
  public pure returns(uint256, uint256){
    uint256 scalar = k;
    if(scalar == 0) return (0,0);
    if (scalar == 1) return (x, y);
    if (scalar == 2) return dblPoint(x, y);
    uint256 base2X = x;
    uint256 base2Y = y;
    uint256 baseX = x;
    uint256 baseY = y;
    if(scalar&1 == 0) baseX = baseY = 0;
    scalar = scalar >> 1;
    while(scalar > 0) {
      (base2X, base2Y) = dblPoint(base2X, base2Y);
      if(scalar&1 == 1) {
        (baseX, baseY) = addPoint(base2X, base2Y, baseX, baseY);
      }
      scalar = scalar >> 1;
    }
    return (baseX, baseY);
  }
  function checkListPk(bytes memory listCompPk, uint256[] memory xPks, uint256[] memory yPks)
  public pure returns(bool){
    if (listCompPk.length%33!=0) return false;
    uint pksLen = listCompPk.length/33;
    if (xPks.length != pksLen) return false;
    if (yPks.length != pksLen) return false;
    //TODO Add more logic
    return true;
  }
  function genCommonParForChkSig(uint256[] memory xPks, uint256[] memory yPks,
  uint[] memory idxSigs, bytes memory bytesR, bytes memory mess)
  private pure returns (uint256, uint256, uint256){
    uint256[4] memory xArr;// 0: AggPk, 1: X, 2: Xi; 3: XTempi
    uint256[4] memory yArr;
    xArr[0] = xPks[0];
    yArr[0] = yPks[0];
    for (uint i = 1; i < xPks.length; i++) {
      (xArr[0], yArr[0]) = addPoint(xArr[0], yArr[0], xPks[i], yPks[i]);
    }
    xArr[1] = 0;
    yArr[1] = 0;
    xArr[2] = 0;
    yArr[2] = 0;
    uint256 ai = 0;
    for (uint i = 0; i < xPks.length; i++) {
      (xArr[3], yArr[3]) = addPoint(xArr[0], yArr[0],xPks[i], yPks[i]);
      ai = uint256(keccak256(compressPoint(xArr[3], yArr[3])))%N;
      (xArr[2], yArr[2]) = scalarMul(xPks[i], yPks[i], ai);
      (xArr[1], yArr[1]) = addPoint(xArr[1], yArr[1], xArr[2], yArr[2]);
    }
    uint256 C = uint256(keccak256(abi.encodePacked(compressPoint(xArr[1], yArr[1]),bytesR,mess)));
    C = C%N;
    if (idxSigs.length!=xPks.length) {
      xArr[1] = 0;
      yArr[1] = 0;
      for (uint i = 0; i < idxSigs.length; i++) {
        (xArr[3], yArr[3]) = addPoint(xArr[0], yArr[0],xPks[idxSigs[i]], yPks[idxSigs[i]]);
        ai = uint256(keccak256(compressPoint(xArr[3], yArr[3])))%N;
        (xArr[2], yArr[2]) = scalarMul(xPks[idxSigs[i]], yPks[idxSigs[i]], ai);
        (xArr[1], yArr[1]) = addPoint(xArr[1], yArr[1], xArr[2], yArr[2]);
      }
    }
    return (C, xArr[1], yArr[1]);
  }

  function checkMulSig(bytes memory listCompPk, uint256[] memory xPks, uint256[] memory yPks, uint[] memory idxRs,
  uint[] memory idxSigs, uint256 xR, uint256 yR, bytes memory bytesR, uint256 sig, bytes memory mess)
  public pure returns(bool){
    // if (!checkListPk(listCompPk, xPks, yPks)) return false;
    // if (idxSigs.length>idxRs.length) return false;
    // if (idxRs.length>xPks.length) return false;
    // if (!chkPointAndCompPoint(xR, yR, bytesR)) return false;
    uint256 C;
    uint256 xX;
    uint256 yX;
    (C, xX, yX) = genCommonParForChkSig(xPks, yPks, idxSigs, bytesR, mess);
    uint256 xLPoint;
    uint256 yLPoint;
    uint256 xRPoint;
    uint256 yRPoint;
    (xLPoint, yLPoint) = scalarMul(Gx, Gy, sig);
    (xRPoint, yRPoint) = scalarMul(xX, yX, C);
    (xRPoint, yRPoint) = addPoint(xRPoint, yRPoint,xR,yR);
    return ((xLPoint==xRPoint) && (yLPoint==yRPoint));
  }
}
