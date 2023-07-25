package types

import (
	"encoding/json"
	"execution/common"
	"execution/types/gadget"
	"math/big"
)

type TxNormal struct {
	Preface TxNormalPreface `json:"preface"` // preface of transaction

	Inner TxNormalInner `json:"inner"` // inner transaction, which could be encrypted

	Extends TxNormalExtends `json:"extends"` // extends of transaction
}

func (tx *TxNormal) TxPreface() TxPreface {
	return &tx.Preface
}

func (tx *TxNormal) TxInner() TxInner {
	return &tx.Inner
}

func (tx *TxNormal) TxExtends() TxExtends {
	return &tx.Extends
}

func (tx *TxNormal) Serialize() []byte {
	ret, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return ret
}

func (tx *TxNormal) Cost() *big.Int {
	gasCost := new(big.Int).Mul(tx.Preface.GasPrice(), new(big.Int).SetUint64(tx.Preface.GasLimit()))
	return gasCost.Add(gasCost, tx.Preface.Value())
}

func (tx *TxNormal) Size() uint64 {
	return uint64(len(tx.Serialize()))
}

type TxNormalPreface struct {
	txHash     common.Hash
	from       common.Address
	nonce      uint64
	gasLimit   uint64
	gasPrice   gadget.GasPrice
	value      *big.Int
	validation gadget.Validation
}

func (txPreface *TxNormalPreface) TxHash() common.Hash {
	return txPreface.txHash
}

func (txPreface *TxNormalPreface) From() common.Address {
	return txPreface.from
}

func (txPreface *TxNormalPreface) Nonce() uint64 {
	return txPreface.nonce
}

func (txPreface *TxNormalPreface) GasLimit() uint64 {
	return txPreface.gasLimit
}

func (txPreface *TxNormalPreface) GasPrice() *big.Int {
	return txPreface.gasPrice.Price()
}

func (txPreface *TxNormalPreface) Value() *big.Int {
	return txPreface.value
}

func (txPreface *TxNormalPreface) Validation() gadget.Validation {
	return txPreface.validation
}

func (TxPreface *TxNormalPreface) InputCoins() []gadget.InputCoin {
	return nil
}

func (TxPreface *TxNormalPreface) OutputCoins() []gadget.OutputCoin {
	return nil
}

func (TxPreface *TxNormalPreface) Witenesses() []gadget.Witness {
	return nil
}

type TxNormalInner struct {
	to         common.Address
	data       []byte
	accessList gadget.AccessList
}

func (txInner *TxNormalInner) To() common.Address {
	return txInner.to
}

func (txInner *TxNormalInner) Data() []byte {
	return txInner.data
}

func (txInner *TxNormalInner) AccessList() gadget.AccessList {
	return txInner.accessList
}

type TxNormalExtends struct {
	refund           gadget.Refund
	extend           []byte
	strictAccessList gadget.AccessList
}

func (txExtends *TxNormalExtends) Refund() gadget.Refund {
	return txExtends.refund
}

func (txExtends *TxNormalExtends) Extend() []byte {
	return txExtends.extend
}

func (txExtends *TxNormalExtends) StrictAccessList() gadget.AccessList {
	return txExtends.strictAccessList
}

// from cannot be determined here
// in fact it should be determined by private key
func NewNormalTransaction(from common.Address, to common.Address,
	gasLimit uint64, gasPrice gadget.GasPrice, value *big.Int, data, extend []byte,
	accessList gadget.AccessList, refund gadget.Refund,
	prv []byte) *TxNormal {
	preface := TxNormalPreface{
		from:     from,
		nonce:    0, // nonce should be fecthed from stateDB
		gasLimit: gasLimit,
		gasPrice: gasPrice,
		value:    value,
	}
	inner := TxNormalInner{
		to:         to,
		data:       data,
		accessList: accessList,
	}

	extends := TxNormalExtends{
		refund:           refund,
		extend:           extend,
		strictAccessList: nil,
	}

	tx := TxNormal{
		Preface: preface,
		Inner:   inner,
		Extends: extends,
	}
	txBytes := tx.Serialize()
	hash := common.GenerateHash(txBytes)
	tx.Preface.txHash = hash

	var validate gadget.SignatureEcdsa
	validate.Sign(hash, prv)
	tx.Preface.validation = &validate

	return &tx
}
