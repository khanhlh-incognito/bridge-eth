all: beacon bridge burn erc20

beacon: build
	go test -run=TestSimulatedSwapBeacon

bridge: build
	go test -run=TestSimulatedSwapBridge

burn: build
	go test -run=TestSimulatedBurn

erc20: build
	go test -run=TestSimulatedErc20

build: incognito_proxy/incognito_proxy.go vault/vault.go erc20/ERC20.go ecdsa_sig/Ecdsa.go

.PHONY: all beacon bridge burn erc20 build

incognito_proxy/incognito_proxy.go: incognito_proxy/incognito_proxy.vy
	./gengo.sh incognito_proxy/incognito_proxy.vy incognito_proxy incognito_proxy

vault/vault.go: vault/vault.vy
	./gengo.sh vault/vault.vy vault vault

erc20/ERC20.go: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20 erc20

ecdsa_sig/Ecdsa.go: ecdsa_sig/contracts/Ecdsa.sol
	./gengo.sh ecdsa_sig/contracts/Ecdsa.sol ecdsa_sig ecdsa_sig

