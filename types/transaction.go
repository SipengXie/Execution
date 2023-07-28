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

	IntrinsicGas() (uint64, error)
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

type Transactions []Transaction

func (txs Transactions) Len() int { return len(txs) }

type TxByNonce Transactions

func (s TxByNonce) Len() int { return len(s) }
func (s TxByNonce) Less(i, j int) bool {
	return (s[i]).TxPreface().Nonce() < (s[j]).TxPreface().Nonce()
}
func (s TxByNonce) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// TxDifference returns a new set which is the difference between a and b.
func TxDifference(a, b Transactions) Transactions {
	keep := make(Transactions, 0, len(a))

	remove := make(map[common.Hash]struct{})
	for _, tx := range b {
		remove[tx.TxPreface().TxHash()] = struct{}{}
	}

	for _, tx := range a {
		if _, ok := remove[tx.TxPreface().TxHash()]; !ok {
			keep = append(keep, tx)
		}
	}

	return keep
}
