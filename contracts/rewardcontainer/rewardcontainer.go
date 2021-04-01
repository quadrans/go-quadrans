// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rewardcontainer

//go:generate abigen --sol contract/ENS.sol --exc contract/AbstractENS.sol:AbstractENS --pkg contract --out contract/ens.go
//go:generate abigen --sol contract/FIFSRegistrar.sol --exc contract/AbstractENS.sol:AbstractENS --pkg contract --out contract/fifsregistrar.go
//go:generate abigen --sol contract/PublicResolver.sol --exc contract/AbstractENS.sol:AbstractENS --pkg contract --out contract/publicresolver.go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/rewardcontainer/contract"
)

var (
	MainNetAddress = common.HexToAddress("0x3D44ba608CD67f175650596E6ae688f396D867ff")
	TestNetAddress = common.HexToAddress("0x3D44ba608CD67f175650596E6ae688f396D867ff")
)

// RewardContainer is the container for the reward given by QuadransFoundation
type RewardContainer struct {
	*contract.RewardContainerSession
	contractBackend bind.ContractBackend
}

// NewRewardContainer creates a struct exposing convenient high-level operations for interacting with
// the RewardContainer smart contract.
func NewRewardContainer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*RewardContainer, error) {
	rc, err := contract.NewRewardContainer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}
	return &RewardContainer{
		&contract.RewardContainerSession{
			Contract:     rc,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

// DeployRewardContainer deploys an instance of the RewardContainer smart contract.
func DeployRewardContainer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *RewardContainer, error) {

	rcAddr, _, _, err := contract.DeployRewardContainer(transactOpts, contractBackend)
	if err != nil {
		return rcAddr, nil, err
	}
	rc, err := NewRewardContainer(transactOpts, rcAddr, contractBackend)
	if err != nil {
		return rcAddr, nil, err
	}
	
	return rcAddr, rc, nil
}

// SetContentHash sets the content hash associated with a name. Only works if the caller
// owns the name, and the associated resolver implements a `setContent` function.
func (RewardContainer *RewardContainer) GetActualReward() (*big.Int, error) {
	opts := RewardContainer.CallOpts
	return RewardContainer.Contract.GetActualReward(&opts)
}

