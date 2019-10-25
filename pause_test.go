package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/incognitochain/bridge-eth/bridge"
	"github.com/pkg/errors"
)

func TestFixedPauseExpired(t *testing.T) {
	// testCases := []struct {
	// 	desc string
	// 	err  bool
	// }{
	// 	{

	// 		desc: "Expired, fail to pause",
	// 	},
	// }
}

func TestFixedPauseTwice(t *testing.T) {
	// testCases := []struct {
	// 	desc string
	// 	err  bool
	// }{
	// 	{

	// 		desc: "Fail to pause twice",
	// 	},
	// }
}

func TestFixedPauseOnce(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc  string
		admin common.Address
		err   bool
	}{
		{
			desc:  "Admin pauses",
			admin: genesisAcc.Address,
		},
		{

			desc:  "Not admin, fail to pause",
			admin: acc.Address,
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _ := setupPauseContract(tc.admin)
			_, err := p.c.Pause(auth)

			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()
		})
	}
}

func setupPauseContract(admin common.Address) (*PausePlatform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("100000000000000000000", 10) // 100 eth
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	sim := backends.NewSimulatedBackend(alloc, 8000000)

	addr, _, c, err := bridge.DeployAdminPausable(auth, sim, admin)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy AdminPausable contract: %v", err)
	}
	_ = addr
	// fmt.Printf("AdminPausable addr: %s\n", addr.Hex())
	sim.Commit()
	return &PausePlatform{sim: sim, c: c}, nil
}

type PausePlatform struct {
	c   *bridge.AdminPausable
	sim *backends.SimulatedBackend
}
