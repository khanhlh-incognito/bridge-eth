all: beacon bridge burn

beacon: incognito_proxy/incognito_proxy.go vault/vault.go erc20/ERC20.go checkMulSig/MulSigP256.go
	go test -run=TestSimulatedSwapBeacon

bridge: incognito_proxy/incognito_proxy.go vault/vault.go erc20/ERC20.go checkMulSig/MulSigP256.go
	go test -run=TestSimulatedSwapBridge

burn: incognito_proxy/incognito_proxy.go vault/vault.go erc20/ERC20.go checkMulSig/MulSigP256.go
	go test -run=TestSimulatedBurn

erc20: incognito_proxy/incognito_proxy.go vault/vault.go erc20/ERC20.go checkMulSig/MulSigP256.go
	go test -run=TestSimulatedErc20

.PHONY: all swap burn

incognito_proxy/incognito_proxy.go: incognito_proxy/incognito_proxy.vy
	./gengo.sh incognito_proxy/incognito_proxy.vy incognito_proxy incognito_proxy

vault/vault.go: vault/vault.vy
	./gengo.sh vault/vault.vy vault vault

erc20/ERC20.go: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20 erc20

checkMulSig/MulSigP256.go: checkMulSig/contracts/MulSigP256.sol
	./gengo.sh checkMulSig/contracts/MulSigP256.sol checkMulSig checkMulSig

