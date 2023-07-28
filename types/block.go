package types

import (
	"execution/common"
	"math/big"
)

type Header interface {
	Hash() common.Hash
	ParentHash() common.Hash
	Number() *big.Int
	GasLimit() uint64
}

type Body interface{}

type Block interface {
	Hash() common.Hash
	ParentHash() common.Hash
	NumberU64() uint64

	Transactions() Transactions
}
