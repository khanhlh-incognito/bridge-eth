package main

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/incognitochain/bridge-eth/bridge"
	"github.com/pkg/errors"
)

func TestFixedExtendExpired(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// Advance time till expired
	err := p.sim.AdjustTime(366 * 24 * time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// Extend, must fail
	days := int64(10)
	_, err = p.c.Extend(auth, big.NewInt(days))
	if err == nil {
		t.Fatal(errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()
}

func TestFixedExtendOnce(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc  string
		admin common.Address
		err   bool
	}{
		{
			desc:  "Admin extends",
			admin: genesisAcc.Address,
		},
		{

			desc:  "Not admin, fail to extend",
			admin: acc.Address,
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _ := setupPauseContract(tc.admin)
			start, _ := p.c.Expire(nil) // Starting expiration time

			days := int64(10)
			_, err := p.c.Extend(auth, big.NewInt(days))

			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			// Check expiration time after extending
			val, _ := p.c.Expire(nil)
			exp := start.Add(start, big.NewInt(days*24*60*60))
			if val.Cmp(exp) < 0 {
				t.Fatal(errors.Errorf("expect expire >= %d, got %d ", exp, val))
			}
		})
	}
}

func TestFixedUnpauseExpired(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// Pause first
	_, err := p.c.Pause(auth)
	if err != nil {
		t.Fatal(err)
	}

	// Advance time till expired
	err = p.sim.AdjustTime(366 * 24 * time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// Unpause, must success
	_, err = p.c.Unpause(auth)
	if err != nil {
		t.Fatal(errors.Errorf("expect error = nil, got %v", err))
	}
	p.sim.Commit()
}

func TestFixedUnpauseTwice(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// Pause first
	_, err := p.c.Pause(auth)
	if err != nil {
		t.Fatal(err)
	}

	// First unpause, must success
	_, err = p.c.Unpause(auth)
	if err != nil {
		t.Fatal(errors.Errorf("expect error = nil, got %v", err))
	}
	p.sim.Commit()

	// Second unpause, must fail
	_, err = p.c.Unpause(auth)
	if err == nil {
		t.Fatal(errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()
}

func TestFixedUnpauseBeforePause(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// Unpause, must fail
	_, err := p.c.Unpause(auth)
	if err == nil {
		t.Fatal(errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()
}

func TestFixedUnpauseOnce(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc  string
		admin *account
		err   bool
	}{
		{
			desc:  "Admin unpauses",
			admin: genesisAcc,
		},
		{

			desc:  "Not admin, fail to unpause",
			admin: acc,
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _ := setupPauseContract(tc.admin.Address)

			// Pause first
			_, err := p.c.Pause(bind.NewKeyedTransactor(tc.admin.PrivateKey))
			if err != nil {
				t.Fatal(err)
			}

			// Unpause
			_, err = p.c.Unpause(auth)

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

func TestFixedPauseExpired(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// Advance time till expired
	err := p.sim.AdjustTime(366 * 24 * time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// Pause, must fail
	_, err = p.c.Pause(auth)
	if err == nil {
		t.Fatal(errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()
}

func TestFixedPauseTwice(t *testing.T) {
	p, _ := setupPauseContract(genesisAcc.Address)

	// First pause, must success
	_, err := p.c.Pause(auth)
	if err != nil {
		t.Fatal(errors.Errorf("expect error = nil, got %v", err))
	}
	p.sim.Commit()

	// Second pause, must fail
	_, err = p.c.Pause(auth)
	if err == nil {
		t.Fatal(errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()
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
	alloc[admin] = core.GenesisAccount{Balance: balance}
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
