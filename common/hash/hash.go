package hash

type CryptoProvider interface {
	Hash(plain string) (string, error)
	Compare(plain, hash string) bool
}

func GetProvider() CryptoProvider {
	return &BcryptProvider{}
}
