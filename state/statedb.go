package state

import (
	"execution/common"
	"math/big"
)

type StateDB interface {
	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	GetBalance(common.Address) *big.Int
	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)
	Copy() StateDB
}
