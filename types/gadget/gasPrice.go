package gadget

import "math/big"

type GasPrice interface {
	Price() *big.Int
}
