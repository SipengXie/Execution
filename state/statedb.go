package state

import (
	"execution/common"
	"math/big"
)

type StateDB interface {
	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	GetBalance(common.Address) *big.Int
	SetBalance(common.Address, *big.Int)
	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)
	Copy() StateDB
}

type EasyStateDB struct {
	balances map[common.Address]*big.Int
	nonces   map[common.Address]uint64
}

func NewEasyStateDB() *EasyStateDB {
	return &EasyStateDB{
		balances: make(map[common.Address]*big.Int),
		nonces:   make(map[common.Address]uint64),
	}
}

func (stateDB *EasyStateDB) Copy() StateDB {
	newStateDB := NewEasyStateDB()
	for addr, balance := range stateDB.balances {
		newStateDB.balances[addr] = balance
	}
	for addr, nonce := range stateDB.nonces {
		newStateDB.nonces[addr] = nonce
	}
	return newStateDB
}

func (stateDB *EasyStateDB) GetNonce(addr common.Address) uint64 {
	return stateDB.nonces[addr]
}

func (stateDB *EasyStateDB) SetNonce(addr common.Address, nonce uint64) {
	stateDB.nonces[addr] = nonce
}

func (stateDB *EasyStateDB) GetBalance(addr common.Address) *big.Int {
	_, ok := stateDB.balances[addr]
	if !ok {
		return big.NewInt(0)
	}
	return stateDB.balances[addr]
}

func (stateDB *EasyStateDB) AddBalance(addr common.Address, amount *big.Int) {
	balance := stateDB.GetBalance(addr)
	balance.Add(balance, amount)
	stateDB.SetBalance(addr, balance)
}

func (stateDB *EasyStateDB) SubBalance(addr common.Address, amount *big.Int) {
	balance := stateDB.GetBalance(addr)
	balance.Sub(balance, amount)
	stateDB.SetBalance(addr, balance)
}

func (stateDB *EasyStateDB) SetBalance(addr common.Address, amount *big.Int) {
	stateDB.balances[addr] = amount
}
