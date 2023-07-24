package types

import (
	"execution/common"
	"execution/types/gadget"
	"math/big"
)

type Transaction interface {
	TxPreface() TxPreface

	TxInner() TxInner

	TxExtends() TxExtends

	Serialize() []byte

	Cost() *big.Int

	Size() uint64
}

type TxPreface interface {
	TxHash() common.Hash
	GasLimit() uint64
	GasPrice() *big.Int
	From() common.Address
	Nonce() uint64
	Value() *big.Int
	Validation() gadget.Validation

	InputCoins() []gadget.InputCoin
	OutputCoins() []gadget.OutputCoin
	Witenesses() []gadget.Witness
}

type TxInner interface {
	To() common.Address
	Data() []byte
	AccessList() gadget.AccessList
}

type TxExtends interface {
	Refund() gadget.Refund
	Extend() []byte
	StrictAccessList() gadget.AccessList
}

type Transactions []*Transaction

type TxByNonce Transactions

func (s TxByNonce) Len() int { return len(s) }
func (s TxByNonce) Less(i, j int) bool {
	return (*s[i]).TxPreface().Nonce() < (*s[j]).TxPreface().Nonce()
}
func (s TxByNonce) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
