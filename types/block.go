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

type Body interface {
	Transactions() Transactions
}

type Block interface {
	Hash() common.Hash
	ParentHash() common.Hash
	NumberU64() uint64

	Transactions() Transactions
}

type EasyHeader struct {
	hash       common.Hash
	parentHash common.Hash
	number     *big.Int
	gasLimit   uint64
}

func NewEasyHeader(hash common.Hash, parentHash common.Hash, number *big.Int, gasLimit uint64) *EasyHeader {
	return &EasyHeader{
		hash:       hash,
		parentHash: parentHash,
		number:     number,
		gasLimit:   gasLimit,
	}
}

func (header *EasyHeader) Hash() common.Hash {
	return header.hash
}

func (header *EasyHeader) ParentHash() common.Hash {
	return header.parentHash
}

func (header *EasyHeader) Number() *big.Int {
	return header.number
}

func (header *EasyHeader) GasLimit() uint64 {
	return header.gasLimit
}

type EasyBody struct {
	transactions Transactions
}

func NewEasyBody(transactions Transactions) *EasyBody {
	return &EasyBody{
		transactions: transactions,
	}
}

func (body *EasyBody) Transactions() Transactions {
	return body.transactions
}

type EasyBlock struct {
	header Header
	body   Body
}

func NewEasyBlock(header Header, body Body) *EasyBlock {
	return &EasyBlock{
		header: header,
		body:   body,
	}
}

func (block *EasyBlock) Hash() common.Hash {
	return block.header.Hash()
}

func (block *EasyBlock) ParentHash() common.Hash {
	return block.header.ParentHash()
}

func (block *EasyBlock) NumberU64() uint64 {
	return block.header.Number().Uint64()
}

func (block *EasyBlock) Transactions() Transactions {
	return block.body.Transactions()
}
