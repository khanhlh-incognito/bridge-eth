pragma solidity >=0.4.21 <0.6.0;

contract MulSigP256 {

    // Set parameters for curve.
    uint constant a = 0xFFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC;
    uint constant b = 0x5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B;
    uint constant gx = 0x6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296;
    uint constant gy = 0x4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5;
    uint constant p = 0xFFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF;
    uint constant n = 0xFFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551;

    uint constant lowSmax = 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0;

    /**
     * @dev Inverse of u in the field of modulo modular.
     */
    function inverseMod(uint num, uint modular)
    public pure returns (uint k)
    {
        if (num == 0 || num == modular)
            return 0;
        int x = 0;
        assembly{
            let t
            let q := div(num,modular)
            let m := mod(num,modular)
            let a0 := modular
            let y := 1
            for {} gt(a0,1) {} {
                q := div(a0,m)
                t := m
                m := mod(a0,m)
                a0 := t
                t := y
                y := sub(x, mul(q, y))
                x := t
            }
        }
        if (x<0) {
            k = modular - uint(-x);
        } else {
            k = uint(x);
        }
    }

    /**
     * @dev Transform affine coordinates into projective coordinates.
     */
    function toProjectivePoint(uint x0, uint y0) public pure
        returns (uint[3] memory P)
    {
        P[2] = addmod(0, 1, p);
        P[0] = mulmod(x0, P[2], p);
        P[1] = mulmod(y0, P[2], p);
    }

    /**
     * @dev Add two points in affine coordinates and return projective point.
     */
    function addAndReturnProjectivePoint(uint x1, uint y1, uint x2, uint y2) public pure
        returns (uint[3] memory P)
    {
        uint x;
        uint y;
        (x, y) = add(x1, y1, x2, y2);
        P = toProjectivePoint(x, y);
    }

    /**
     * @dev Transform from projective to affine coordinates.
     */
    function toAffinePoint(uint x0, uint y0, uint z0) public pure
        returns (uint x1, uint y1)
    {
        uint z0Inv;
        z0Inv = inverseMod(z0, p);
        x1 = mulmod(x0, z0Inv, p);
        y1 = mulmod(y0, z0Inv, p);
    }

    /**
     * @dev Return the zero curve in projective coordinates.
     */
    function zeroProj() public pure
        returns (uint x, uint y, uint z)
    {
        return (0, 1, 0);
    }

    /**
     * @dev Return the zero curve in affine coordinates.
     */
    function zeroAffine() public pure
        returns (uint x, uint y)
    {
        return (0, 0);
    }

    /**
     * @dev Check if the curve is the zero curve.
     */
    function isZeroCurve(uint x0, uint y0) public pure
        returns (bool isZero)
    {
        if(x0 == 0 && y0 == 0) {
            return true;
        }
        return false;
    }

    /**
     * @dev Check if a point in affine coordinates is on the curve.
     */
    function isOnCurve(uint x, uint y) public pure
        returns (bool)
    {
        if (0 == x || x == p || 0 == y || y == p) {
            return false;
        }

        uint LHS = mulmod(y, y, p); // y^2
        uint RHS = mulmod(mulmod(x, x, p), x, p); // x^3

        if (a != 0) {
            RHS = addmod(RHS, mulmod(x, a, p), p); // x^3 + a*x
        }
        if (b != 0) {
            RHS = addmod(RHS, b, p); // x^3 + a*x + b
        }

        return LHS == RHS;
    }

    /**
     * @dev Double an elliptic curve point in projective coordinates. See
     * https://www.nayuki.io/page/elliptic-curve-point-addition-in-projective-coordinates
     */
    function twiceProj(uint x, uint y, uint z) public pure
        returns (uint x1, uint y1, uint z1)
    {
        uint x0 = x;
        uint y0 = y;
        uint z0 = z;
        uint t;
        uint u;
        uint v;
        uint w;

        if(isZeroCurve(x0, y0)) {
            return zeroProj();
        }

        u = mulmod(y0, z0, p);
        u = mulmod(u, 2, p);

        v = mulmod(u, x0, p);
        v = mulmod(v, y0, p);
        v = mulmod(v, 2, p);

        x0 = mulmod(x0, x0, p);
        t = mulmod(x0, 3, p);

        z0 = mulmod(z0, z0, p);
        z0 = mulmod(z0, a, p);
        t = addmod(t, z0, p);

        w = mulmod(t, t, p);
        x0 = mulmod(2, v, p);
        w = addmod(w, p-x0, p);

        x0 = addmod(v, p-w, p);
        x0 = mulmod(t, x0, p);
        y0 = mulmod(y0, u, p);
        y0 = mulmod(y0, y0, p);
        y0 = mulmod(2, y0, p);
        y1 = addmod(x0, p-y0, p);

        x1 = mulmod(u, w, p);

        z1 = mulmod(u, u, p);
        z1 = mulmod(z1, u, p);
    }

    /**
     * @dev Add two elliptic curve points in projective coordinates. See
     * https://www.nayuki.io/page/elliptic-curve-point-addition-in-projective-coordinates
     */
    function addProj(uint x0, uint y0, uint z0, uint x1, uint y1, uint z1) public pure
        returns (uint x2, uint y2, uint z2)
    {
        uint t0;
        uint t1;
        uint u0;
        uint u1;

        if (isZeroCurve(x0, y0)) {
            return (x1, y1, z1);
        }
        else if (isZeroCurve(x1, y1)) {
            return (x0, y0, z0);
        }

        t0 = mulmod(y0, z1, p);
        t1 = mulmod(y1, z0, p);

        u0 = mulmod(x0, z1, p);
        u1 = mulmod(x1, z0, p);

        if (u0 == u1) {
            if (t0 == t1) {
                return twiceProj(x0, y0, z0);
            }
            else {
                return zeroProj();
            }
        }

        (x2, y2, z2) = addProj2(mulmod(z0, z1, p), u0, u1, t1, t0);
    }

    /**
     * @dev Helper function that splits addProj to avoid too many local variables.
     */
    function addProj2(uint k, uint uu0, uint uu1, uint tt1, uint tt0) private pure
        returns (uint x2, uint y2, uint z2)
    {
        uint v = k;
        uint u0 = uu0;
        uint u1 = uu1;
        uint t1 = tt1;
        uint t0 = tt0;
        uint u;
        uint u2;
        uint u3;
        uint w;
        uint t;

        t = addmod(t0, p-t1, p);
        u = addmod(u0, p-u1, p);
        u2 = mulmod(u, u, p);

        w = mulmod(t, t, p);
        w = mulmod(w, v, p);
        u1 = addmod(u1, u0, p);
        u1 = mulmod(u1, u2, p);
        w = addmod(w, p-u1, p);

        x2 = mulmod(u, w, p);

        u3 = mulmod(u2, u, p);
        u0 = mulmod(u0, u2, p);
        u0 = addmod(u0, p-w, p);
        t = mulmod(t, u0, p);
        t0 = mulmod(t0, u3, p);

        y2 = addmod(t, p-t0, p);

        z2 = mulmod(u3, v, p);
    }

    /**
     * @dev Add two elliptic curve points in affine coordinates.
     */
    function add(uint x0, uint y0, uint x1, uint y1) public pure
        returns (uint, uint)
    {
        uint z0;

        (x0, y0, z0) = addProj(x0, y0, 1, x1, y1, 1);

        return toAffinePoint(x0, y0, z0);
    }

    /**
     * @dev Double an elliptic curve point in affine coordinates.
     */
    function twice(uint x0, uint y0) public pure
        returns (uint, uint)
    {
        uint z0;

        (x0, y0, z0) = twiceProj(x0, y0, 1);

        return toAffinePoint(x0, y0, z0);
    }

    /**
     * @dev Multiply an elliptic curve point by a 2 power base (i.e., (2^exp)*P)).
     */
    function multiplyPowerBase2(uint x0, uint y0, uint exp) public pure
        returns (uint, uint)
    {
        uint base2X = x0;
        uint base2Y = y0;
        uint base2Z = 1;

        for(uint i = 0; i < exp; i++) {
            (base2X, base2Y, base2Z) = twiceProj(base2X, base2Y, base2Z);
        }

        return toAffinePoint(base2X, base2Y, base2Z);
    }

    /**
     * @dev Multiply an elliptic curve point by a scalar.
     */
    function multiplyScalar(uint x, uint y, uint k) public pure
        returns (uint x1, uint y1)
    {
        uint x0 = x;
        uint y0 = y;
        uint scalar = k;
        if(scalar == 0) {
            return zeroAffine();
        }
        else if (scalar == 1) {
            return (x0, y0);
        }
        else if (scalar == 2) {
            return twice(x0, y0);
        }

        uint base2X = x0;
        uint base2Y = y0;
        uint base2Z = 1;
        uint z1 = 1;
        x1 = x0;
        y1 = y0;

        if(scalar&1 == 0) {
            x1 = y1 = 0;
        }

        scalar = scalar >> 1;

        while(scalar > 0) {
            (base2X, base2Y, base2Z) = twiceProj(base2X, base2Y, base2Z);

            if(scalar&1 == 1) {
                (x1, y1, z1) = addProj(base2X, base2Y, base2Z, x1, y1, z1);
            }

            scalar = scalar >> 1;
        }

        return toAffinePoint(x1, y1, z1);
    }

    /**
     * @dev Multiply the curve's generator point by a scalar.
     */
    function multipleGeneratorByScalar(uint scalar) public pure
        returns (uint, uint)
    {
        return multiplyScalar(gx, gy, scalar);
    }

    function compressPoint(uint x, uint y)
    private pure returns(bytes memory){
        bytes1 format = 0x02;
        if (y&1 == 1) format |= 0x01;
        return abi.encodePacked(format, bytes32(x));
    }

    /**
     * @dev Generate common parameters for check multi-signature
     */
    function genCommonParams(
        uint[8] memory xPks,
        uint[8] memory yPks,
        uint[8] memory idxSigs,
        uint8 sigLen,
        bytes memory bytesR,
        bytes32 mess)
    public pure returns (uint[3] memory res)
    {
        uint[5] memory xArr;// 0: AggPk, 1: X, 2: Xi; 3: XTempi, 4: (ai, zRes)
        uint[4] memory yArr;
        uint[2] memory zArr;
        bytes memory temp;
        xArr[0] = xPks[0];
        yArr[0] = yPks[0];
        zArr[0] = 1;
        for (uint i = 1; i < xPks.length; i++) {
            if ((xPks[i] | yPks[i]) == 0) {
                break;
            }
            (xArr[0], yArr[0], zArr[0]) = addProj(xArr[0], yArr[0], zArr[0], xPks[i], yPks[i], 1);
        }
        (xArr[0], yArr[0]) = toAffinePoint(xArr[0], yArr[0], zArr[0]);
        xArr[1] = 0;
        yArr[1] = 0;
        zArr[1] = 1;
        zArr[0] = 1;
        xArr[2] = 0;
        yArr[2] = 0;
        xArr[4] = 0;
        res[0] = 0;
        res[1] = 0;
        res[2] = 0;
        for (uint i = 0; i < xPks.length; i++) {
            if ((xPks[i] | yPks[i]) == 0) {
                break;
            }
            (xArr[3], yArr[3]) = add(xArr[0], yArr[0], xPks[i], yPks[i]);
            temp = compressPoint(xArr[3], yArr[3]);
            xArr[4] = uint(keccak256(temp))%n;
            (xArr[2], yArr[2]) = multiplyScalar(xPks[i], yPks[i], xArr[4]);
            if ((i==idxSigs[res[0]]) && (res[0]<sigLen)) {
               (res[1], res[2], zArr[0]) = addProj(res[1], res[2], zArr[0], xArr[2], yArr[2], 1);
                res[0]++;
            }
            (xArr[1], yArr[1], zArr[1]) = addProj(xArr[1], yArr[1], zArr[1], xArr[2], yArr[2], 1);
        }
        (xArr[3], yArr[3]) = toAffinePoint(xArr[1], yArr[1], zArr[1]);
        temp = compressPoint(xArr[3], yArr[3]);
        res[0] = uint(keccak256(abi.encodePacked(temp,bytesR,mess)));
        res[0] = res[0]%n;
        (res[1], res[2]) = toAffinePoint(res[1], res[2], zArr[0]);
    }

    function checkMulSig(
        uint[8] memory xPks,
        uint[8] memory yPks,
        uint[8] memory idxSigs,
        int128 sigsLen,
        uint xR,
        uint yR,
        bytes memory bytesR,
        uint sig,
        bytes32 mess)
    public pure returns(bool)
    {
        if (sigsLen>8) return false;
        uint8 sigLen = uint8(sigsLen);
        uint[3] memory Info;//0:C; 1: xX, 2: yX
        Info = genCommonParams(xPks, yPks, idxSigs, sigLen, bytesR, mess);
        uint xLPoint = 0;
        uint yLPoint = 0;
        uint xRPoint = 0;
        uint yRPoint = 0;
        (xLPoint, yLPoint) = multiplyScalar(gx, gy, sig);
        (xRPoint, yRPoint) = multiplyScalar(Info[1], Info[2], Info[0]);
        (xRPoint, yRPoint) = add(xRPoint, yRPoint,xR,yR);
        return ((xLPoint==xRPoint) && (yLPoint==yRPoint));
    }

}
