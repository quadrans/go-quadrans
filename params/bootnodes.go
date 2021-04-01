// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Quadrans Foundation Go Bootnodes
	"enode://54aad46c948555cf0b054d099f0186c85e938f06ab1588a6dde55c6f5fab14000a11bea7a01c88e5489c5a7a17ec8bf6d1900650750ed73c121da44d1a0d84a3@40.113.152.189:28657", // Europe
	"enode://f4739708f20e95d78b58a3b830a056d7da7c6ccd489f892e638f9d3fe59d6b4e6242a0965b088a5e76a3bf3dc5c8abb897e8597312ddf54b05d013ceb08a6e83@40.70.129.31:28657",  // Est America
	"enode://d3f5fe83f429fd1349908ecec9edd073a7fa8490e3600ff612b16b4f19cead8b6086f0268cc7f5b5ce9baef81b36c2fc360d99f594bec9b48155a776fa6a909c@23.101.9.111:28657",  // Asia 
	"enode://0fd93f288443798b6ffcd4a5d00ac42eb95088994dce5a80a1d18d611ca05088419f7db80d79df95f803fa9ca34bcc7403f96bd017ac6c5af17b7deb44b27549@23.101.211.229:30301",  // Australia
	"enode://8f5ec5b6a9be9c68c0ce09f6fbf0778788708975152d049b147c6894c8b998a8f96c4822c0f7e4115aa8e7f94f3dca30e78476565d93413e527816daa502b021@23.97.213.110:30301",  // Europe
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
  // Quadrans Foundation Testnet Bootnodes
	"enode://e13221932a06ee3b5ef17dc09d34a73e57c68ae99e85a88ce0ef26b1909f0d0fecd884eececf4eda1978586b14b790c109c00729bdd26ab72547117db5ab29fb@52.169.192.210:28657", // RPC
	"enode://565fe0a0ef0690fe12a2db7e6bf7793e697670862e4ba0fe1c3ff6e067da39c1c8559572872fc85eb80cb16e9b3ddf61178e210eeb05f4e25439caadc3d539de@40.74.72.56:28657",  // Testnet Asia
	"enode://3e05163a55bf163d652fe72545d4c23887cb2a4314665a9cebc9dd3b3d93f1e739439327d202abab559293a42a168c5400b87c56fa7371d8162f1eb25c6d7516@52.167.132.30:28657",  // Testnet America
	"enode://8f645fe27dee31f1551c737fc1b7ef9acf20cf0103d48f89b505de105b1ad16960220d52e9d69cb11e1a98cc576a56a2dda6724125fb69a3df29770e97235de2@104.45.7.40:28657",  // Testnet Europe
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://8641ee7eafd94d87cf48a53eb5494ba81adffaf7cc8c4f97b4ebb8873d3e0d55bf2b7c73e915431566b4d5ca4baf3590d67c5a5e6586b33a9c5d1c6757d1df53@40.114.84.130:30301", // Est America
	"enode://eef6b3db8b5f9e772c1f699772090e62fd44b9844f20cd878db838b7fd3c9dfe297204fd47ca5ee66905bb5ee3303e96be7d1f05272b3743237e9a40f7f6af00@52.183.3.127:30301",  // West America
	"enode://f3b04baf0b297ff821604bfb06651db3eac78c0cdb3683ca8fe4878a8d92aee26ffd842077db7c8d9fc3acf6eb9d7ffadc6f8ac32a456849f60c21454426d63a@168.63.250.64:30301",  // Asia Pacific
	"enode://0fd93f288443798b6ffcd4a5d00ac42eb95088994dce5a80a1d18d611ca05088419f7db80d79df95f803fa9ca34bcc7403f96bd017ac6c5af17b7deb44b27549@23.101.211.229:30301",  // Australia
	"enode://8f5ec5b6a9be9c68c0ce09f6fbf0778788708975152d049b147c6894c8b998a8f96c4822c0f7e4115aa8e7f94f3dca30e78476565d93413e527816daa502b021@23.97.213.110:30301",  // Europe
}
