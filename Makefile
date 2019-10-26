all: beacon bridge burn erc20

beacon: build
	go test -run=TestSimulatedSwapBeacon

bridge: build
	go test -run=TestSimulatedSwapBridge

burn: build
	go test -run=TestSimulatedBurn

erc20: build
	go test -run=TestSimulatedErc20

build: bridge/incognito_proxy/incognito_proxy.go bridge/vault/vault.go

.PHONY: all beacon bridge burn erc20 build

erc20/ERC20.go: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20

bridge/incognito_proxy/incognito_proxy.go: bridge/contracts/incognito_proxy.sol
	./gengo.sh bridge/contracts/incognito_proxy.sol bridge/incognito_proxy

bridge/vault/vault.go: bridge/contracts/vault.sol
	./gengo.sh bridge/contracts/vault.sol bridge/vault
