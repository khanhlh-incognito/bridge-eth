all: swap burn

swap: incognito_proxy/incognito_proxy.abi vault/vault.abi erc20/ERC20.abi
	go test -run=TestSimulatedSwapBeacon

burn: incognito_proxy/incognito_proxy.abi vault/vault.abi erc20/ERC20.abi
	go test -run=TestSimulatedBurn

erc20: incognito_proxy/incognito_proxy.abi vault/vault.abi erc20/ERC20.abi
	go test -run=TestSimulatedErc20

.PHONY: all swap burn

incognito_proxy/incognito_proxy.abi: incognito_proxy/incognito_proxy.vy
	./gengo.sh incognito_proxy/incognito_proxy.vy incognito_proxy

vault/vault.abi: vault/vault.vy
	./gengo.sh vault/vault.vy vault

erc20/ERC20.abi: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20

