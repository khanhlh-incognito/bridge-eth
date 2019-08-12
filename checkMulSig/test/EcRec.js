const ec = artifacts.require('EcRec.sol');
let c;

contract("MulSigP256", (accounts) => {
    before(async() => {
        c = await ec.deployed()
    })
    describe('Test ECREC', () => {
        it.only('Test', async() => {
            var address = accounts[0]
            var msg = 'OpenZeppelin'

            var h = web3.utils.sha3(msg)
            var sig = await web3.eth.sign(h, address)
            sig = sig.slice(2)
            var r = `0x${sig.slice(0, 64)}`
            var s = `0x${sig.slice(64, 128)}`
            var v = web3.utils.toDecimal(sig.slice(128, 130)) + 27
            const res = await c.go(v);
            console.log(res);
        });
    });
});
