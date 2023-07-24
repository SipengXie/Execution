package types

import (
	"encoding/json"
	"execution/common"
	"execution/types/gadget"
	"math/big"
)

type TxWithdraw struct {
	Preface TxWithdrawPreface `json:"preface"`
}

func (tx *TxWithdraw) TxPreface() TxPreface {
	return &tx.Preface
}

func (tx *TxWithdraw) TxInner() TxInner {
	return nil
}

func (tx *TxWithdraw) TxExtends() TxExtends {
	return nil
}

func (tx *TxWithdraw) Serialize() []byte {
	ret, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return ret
}

func (tx *TxWithdraw) Cost() *big.Int {
	gasCost := new(big.Int).Mul(tx.Preface.GasPrice(), new(big.Int).SetUint64(tx.Preface.GasLimit()))
	return gasCost.Add(gasCost, tx.Preface.Value())
}

type TxWithdrawPreface struct {
	txHash      common.Hash
	from        common.Address
	nonce       uint64
	validation  gadget.Validation
	gasPrice    gadget.GasPrice
	outputCoins []gadget.OutputCoin
}

func (txPreface *TxWithdrawPreface) TxHash() common.Hash {
	return txPreface.txHash
}

func (txPreface *TxWithdrawPreface) GasLimit() uint64 {
	return 0
}

func (txPreface *TxWithdrawPreface) GasPrice() *big.Int {
	return txPreface.gasPrice.Price()
}

func (txPreface *TxWithdrawPreface) From() common.Address {
	return txPreface.from
}

func (txPreface *TxWithdrawPreface) Nonce() uint64 {
	return txPreface.nonce
}

func (txPreface *TxWithdrawPreface) Value() *big.Int {
	outBalance := big.NewInt(0)
	for _, out := range txPreface.outputCoins {
		outBalance.Add(outBalance, out.Amount)
	}
	return outBalance
}

func (txPreface *TxWithdrawPreface) Validation() gadget.Validation {
	return txPreface.validation
}

func (txPreface *TxWithdrawPreface) InputCoins() []gadget.InputCoin {
	return nil
}

func (txPreface *TxWithdrawPreface) OutputCoins() []gadget.OutputCoin {
	return txPreface.outputCoins
}

func (TxPreface *TxWithdrawPreface) Witenesses() []gadget.Witness {
	return nil
}
