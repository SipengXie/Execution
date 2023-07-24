package gadget

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"execution/common"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto" // substitude to our crypto
)

type Validation interface {
	GetFrom(input common.Hash) (common.Address, error)
	Sign(input common.Hash, prv []byte)
}

var (
	ErrInvalidSignature = errors.New("invalid signature")
	ErrInvalidPubKey    = errors.New("invalid public key")
)

type SignatureEcdsa struct {
	r, s *big.Int
	v    byte
}

func (sign *SignatureEcdsa) GetFrom(input common.Hash) (common.Address, error) {
	if len(sign.r.Bytes()) != 32 || len(sign.s.Bytes()) != 32 || sign.v != 0 && sign.v != 1 {
		return common.Address{}, ErrInvalidSignature
	}

	sig := make([]byte, 65)
	copy(sig[:32], sign.r.Bytes())
	copy(sig[32:64], sign.s.Bytes())
	sig[64] = sign.v

	pub, err := crypto.Ecrecover(input[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, ErrInvalidPubKey
	}

	pubHash := sha256.Sum256(pub[1:])
	var addr common.Address
	copy(addr[:], pubHash[12:])

	return addr, nil
}

func (sign *SignatureEcdsa) Sign(input common.Hash, prv []byte) {
	var key ecdsa.PrivateKey
	key.D = new(big.Int).SetBytes(prv)
	sig, err := crypto.Sign(input[:], &key)
	if err != nil {
		panic(err)
	}
	sign.r = new(big.Int).SetBytes(sig[:32])
	sign.s = new(big.Int).SetBytes(sig[32:64])
	sign.v = sig[64]
}
