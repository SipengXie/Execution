package common

import "crypto/sha256"

const (
	HashLength    = 32
	AddressLength = 20
)

type Hash [HashLength]byte
type Address [AddressLength]byte

func GenerateHash(input []byte) Hash {
	hash := sha256.Sum256(input)
	return hash
}
