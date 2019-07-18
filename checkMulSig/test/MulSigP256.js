const Curve = artifacts.require('MulSigP256.sol');
let c;
contract("MulSigP256", () => {
    // const root = accounts[0]
    // var c
    let pks = [2,244,10,50,214,96,112,251,201,201,110,126,232,227,57,201,189,154,6,22,181,143,236,222,127,182,82,236,230,199,207,81,62,3,152,162,3,129,152,121,175,117,94,24,187,247,62,34,63,4,61,1,107,28,4,66,173,176,102,167,204,6,205,27,22,247,2,226,194,204,2,30,129,113,140,197,221,136,234,156,58,27,68,15,156,128,118,38,21,1,253,143,202,87,118,159,168,113,228,3,32,204,211,190,182,245,238,235,179,71,69,216,31,224,16,113,161,125,82,133,152,89,201,114,225,169,23,127,154,109,148,229,2,62,93,133,79,192,235,214,182,146,97,158,153,171,52,180,108,1,195,50,100,117,217,250,213,161,178,131,66,212,21,167,83,2,121,116,243,234,58,232,250,80,224,2,206,23,171,17,53,114,39,76,9,243,65,153,70,99,239,201,207,54,134,172,20,196,3,76,52,144,182,161,75,24,222,6,32,172,18,231,204,209,8,211,201,21,159,102,131,134,20,77,202,118,1,133,214,120,82,3,121,234,213,207,189,122,215,171,88,46,16,6,25,21,32,44,65,231,111,4,248,54,193,135,59,253,250,236,131,36,107,180,3,228,75,238,201,69,202,21,144,94,15,213,216,210,64,76,136,28,44,223,55,239,30,124,89,36,170,53,155,118,23,36,19,3,75,89,232,231,74,6,173,136,59,47,96,120,74,233,52,243,78,114,108,134,47,72,223,173,178,159,6,5,105,138,112,223,2,92,112,90,40,173,187,97,181,72,172,22,44,7,240,168,216,16,4,211,150,98,129,67,34,182,233,69,82,151,162,118,95,3,162,165,146,1,76,157,202,56,127,242,218,12,236,112,142,201,54,122,90,51,183,18,63,174,238,66,205,249,221,63,90,8,2,191,220,170,82,157,96,232,100,79,74,123,4,44,69,180,206,31,203,94,157,26,238,25,122,28,78,152,86,104,209,160,106,3,58,239,140,232,68,142,225,25,85,198,91,171,237,27,67,120,70,153,135,47,152,40,9,127,177,11,148,29,132,90,17,176,3,80,88,197,96,111,60,180,122,101,165,207,73,179,172,108,211,48,87,241,119,117,101,217,44,125,112,86,66,109,12,43,108,2,93,106,245,177,96,144,95,17,42,105,232,161,83,80,50,55,61,160,170,102,77,179,230,30,86,200,145,44,37,202,80,49]
    let xPks = ["0xf40a32d66070fbc9c96e7ee8e339c9bd9a0616b58fecde7fb652ece6c7cf513e", "0x98a203819879af755e18bbf73e223f043d016b1c0442adb066a7cc06cd1b16f7", "0xe2c2cc021e81718cc5dd88ea9c3a1b440f9c8076261501fd8fca57769fa871e4", "0x20ccd3beb6f5eeebb34745d81fe01071a17d52859859c972e1a9177f9a6d94e5","0xf40a32d66070fbc9c96e7ee8e339c9bd9a0616b58fecde7fb652ece6c7cf513e", "0x98a203819879af755e18bbf73e223f043d016b1c0442adb066a7cc06cd1b16f7", "0xf40a32d66070fbc9c96e7ee8e339c9bd9a0616b58fecde7fb652ece6c7cf513e", "0x98a203819879af755e18bbf73e223f043d016b1c0442adb066a7cc06cd1b16f7"]
    let yPks = ["0x6779ac5c5ae05fe6f774279fae540e927e034cb97bbd4791fc054f8d1d196872", "0xdf32097043caae6a2f02f6d6987e3764b8bfd73b7f712d7a702e82aab534d6cd", "0xd2152a96f3eff7e55ec0c236730c04eb53c7c12cc12aa613ab70c63436229c76", "0x953ac439a37085c82e6da60887195f82dd7aa620a238b759c9cb3050db6ed547","0x6779ac5c5ae05fe6f774279fae540e927e034cb97bbd4791fc054f8d1d196872", "0xdf32097043caae6a2f02f6d6987e3764b8bfd73b7f712d7a702e82aab534d6cd", "0x6779ac5c5ae05fe6f774279fae540e927e034cb97bbd4791fc054f8d1d196872", "0xdf32097043caae6a2f02f6d6987e3764b8bfd73b7f712d7a702e82aab534d6cd"]
    let idxs = [0,1,2,3,4,5,6,7]
    let idxr = [0,1,2,3,4,5,6,7]
    let sigLen = 8
    let R = [2,190,53,62,71,90,165,51,83,163,251,184,203,39,129,231,153,168,58,168,137,212,91,232,36,239,115,180,253,214,75,2,69]
    let xR = "0xbe353e475aa53353a3fbb8cb2781e799a83aa889d45be824ef73b4fdd64b0245"
    let yR = "0x938f8035937fcacf1be54d0163edb88ac66b8467196acf354097e9ac2c55e422"
    let sig ="0x938f8035937fcacf1be54d0163edb88ac66b8467196acf354097e9ac2c55e422"
    let mess =[104,31,251,6,255,94,237,227,155,230,178,151,200,0,65,141,22,136,0,62,126,139,233,49,154,28,240,35,59,136,138,7]
    before(async() => {

        // function checkMulSig(bytes memory listCompPk, uint256[] memory xPks, uint256[] memory yPks, uint[] memory idxRs,
  //uint[] memory idxSigs, uint256 xR, uint256 yR, bytes memory bytesR, uint256 sig, bytes memory mess)
        c = await Curve.deployed()
    })
    describe('TestMulSig', () => {
        it('Test inverse mod', async() => {
            // const gasUsed = await c.checkMulSig.estimateGas(pks,xPks,yPks,idxr,idxs,xR,yR,R,sig,mess)
            // txs = await c.scalarMul(xPks[0],xPks[1],xPks[2],xPks[3])
            const gasUsed = await c.inverseMod(yR, "0xFFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF");
            console.log(gasUsed);
            const gasUsed2 = await c.inverseMod.estimateGas(yR, "0xFFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF");
            console.log(gasUsed2);
        });

        it('Test total gas genCommonParams', async() => {
            // const gasUsed = await c.checkMulSig.estimateGas(pks,xPks,yPks,idxr,idxs,xR,yR,R,sig,mess)
            // txs = await c.scalarMul(xPks[0],xPks[1],xPks[2],xPks[3])
            const gasUsed = await c.genCommonParams.estimateGas(xPks,yPks,idxs,R,mess);
            console.log(gasUsed);
        });
        it.only('Test total gas Verify multi sig', async() => {
            // const gasUsed = await c.checkMulSig.estimateGas(pks,xPks,yPks,idxr,idxs,xR,yR,R,sig,mess)
            // txs = await c.scalarMul(xPks[0],xPks[1],xPks[2],xPks[3])
            const gasUsed = await c.checkMulSig.estimateGas(xPks,yPks,idxs,sigLen,xR,yR,R,sig,mess);
            console.log(gasUsed);
        });
    });
});
