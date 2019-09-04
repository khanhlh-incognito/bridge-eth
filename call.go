package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/bridge"
)

func withdraw(v *bridge.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.Withdraw(
		auth,
		proof.Instruction,
		proof.Heights,

		proof.InstPaths,
		proof.InstPathIsLefts,
		proof.InstRoots,
		proof.BlkData,
		proof.SigIdxs,
		proof.SigVs,
		proof.SigRs,
		proof.SigSs,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SwapBridge(inc *bridge.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	tx, err := inc.SwapBridgeCommittee(
		auth,
		proof.Instruction,

		proof.InstPaths,
		proof.InstPathIsLefts,
		proof.InstRoots,
		proof.BlkData,
		proof.SigIdxs,
		proof.SigVs,
		proof.SigRs,
		proof.SigSs,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SwapBeacon(inc *bridge.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	tx, err := inc.SwapBeaconCommittee(
		auth,
		proof.Instruction,

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
