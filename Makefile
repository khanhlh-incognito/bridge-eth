all: beacon bridge burn erc20

beacon: build
	go test -run=TestSimulatedSwapBeacon

bridge: build
	go test -run=TestSimulatedSwapBridge

burn: build
	go test -run=TestSimulatedBurn

erc20: build
	go test -run=TestSimulatedErc20

build: bridge/incognito_proxy.go

.PHONY: all beacon bridge burn erc20 build

incognito_proxy/incognito_proxy.go: incognito_proxy/incognito_proxy.vy
	./gengo.sh incognito_proxy/incognito_proxy.vy incognito_proxy

vault/vault.go: vault/vault.vy
	./gengo.sh vault/vault.vy vault

erc20/ERC20.go: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20

ecdsa_sig/Ecdsa.go: ecdsa_sig/contracts/Ecdsa.sol
	./gengo.sh ecdsa_sig/contracts/Ecdsa.sol ecdsa_sig

bridge/incognito_proxy.go: bridge/contracts/incognito_proxy.sol
	./gengo.sh bridge/contracts/incognito_proxy.sol bridge
