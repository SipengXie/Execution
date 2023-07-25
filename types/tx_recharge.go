package types

import (
	"encoding/json"
	"execution/common"
	"execution/types/gadget"
	"math/big"
)

type TxRecharge struct {
	Preface TxPrefaceRecharge `json:"preface"`
	Inner   TxInnerRecharge   `json:"inner"`
}

func (tx *TxRecharge) TxPreface() TxPreface {
	return &tx.Preface
}

func (tx *TxRecharge) TxInner() TxInner {
	return &tx.Inner
}

func (tx *TxRecharge) TxExtends() TxExtends {
	return nil
}

func (tx *TxRecharge) Serialize() []byte {
	ret, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return ret
}

func (tx *TxRecharge) Cost() *big.Int {
	gasCost := new(big.Int).Mul(tx.Preface.GasPrice(), new(big.Int).SetUint64(tx.Preface.GasLimit()))
	return gasCost.Add(gasCost, tx.Preface.Value())
}

func (tx *TxRecharge) Size() uint64 {
	return uint64(len(tx.Serialize()))
}

type TxPrefaceRecharge struct {
	txHash     common.Hash
	inputCoins []gadget.InputCoin
	witenesses []gadget.Witness
	gasPrice   gadget.GasPrice
}

type TxInnerRecharge struct {
	to common.Address
}

func (txPreface *TxPrefaceRecharge) TxHash() common.Hash {
	return txPreface.txHash
}

func (txPreface *TxPrefaceRecharge) InputCoins() []gadget.InputCoin {
	return txPreface.inputCoins
}

func (txPreface *TxPrefaceRecharge) Witenesses() []gadget.Witness {
	return txPreface.witenesses
}

func (txPreface *TxPrefaceRecharge) GasPrice() *big.Int {
	return txPreface.gasPrice.Price()
}

func (txPreface *TxPrefaceRecharge) GasLimit() uint64 {
	return 0
}

func (txPreface *TxPrefaceRecharge) From() common.Address {
	return common.Address{}
}

func (txPreface *TxPrefaceRecharge) Nonce() uint64 {
	return 0
}

func (txPreface *TxPrefaceRecharge) Value() *big.Int {
	return big.NewInt(0)
}

func (txPreface *TxPrefaceRecharge) Validation() gadget.Validation {
	return nil
}

func (txPreface *TxPrefaceRecharge) OutputCoins() []gadget.OutputCoin {
	return nil
}

func (txInner *TxInnerRecharge) To() common.Address {
	return txInner.to
}

func (txInner *TxInnerRecharge) Data() []byte {
	return nil
}

func (txInner *TxInnerRecharge) AccessList() gadget.AccessList {
	return nil
}
