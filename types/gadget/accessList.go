package gadget

type AccessList interface {
	Len() int
	StorageKeys() int
}
