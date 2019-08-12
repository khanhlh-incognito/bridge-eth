package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/vault"
)

func withdraw(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.Withdraw(
		auth,
		proof.Instruction,

		proof.BeaconHeight,
		proof.BeaconInstPath,
		proof.BeaconInstPathIsLeft,
		proof.BeaconInstPathLen,
		proof.BeaconInstRoot,
		proof.BeaconBlkData,
		proof.BeaconBlkHash,
		proof.BeaconSignerSig,
		proof.BeaconNumR,
		proof.BeaconXs,
		proof.BeaconYs,
		proof.BeaconRIdxs,
		proof.BeaconNumSig,
		proof.BeaconSigIdxs,
		proof.BeaconRp,
		proof.BeaconRpx,
		proof.BeaconRpy,
		proof.BeaconR,

		proof.BridgeHeight,
		proof.BridgeInstPath,
		proof.BridgeInstPathIsLeft,
		proof.BridgeInstPathLen,
		proof.BridgeInstRoot,
		proof.BridgeBlkData,
		proof.BridgeBlkHash,
		proof.BridgeSignerSig,
		proof.BridgeNumR,
		proof.BridgeXs,
		proof.BridgeYs,
		proof.BridgeRIdxs,
		proof.BridgeNumSig,
		proof.BridgeSigIdxs,
		proof.BridgeRp,
		proof.BridgeRpx,
		proof.BridgeRpy,
		proof.BridgeR,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
