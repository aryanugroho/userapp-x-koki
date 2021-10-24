package crypto

import "golang.org/x/crypto/bcrypt"

const (
	cost = 4 // keep the minimum for demo purpose, better move to config
)

type BcryptProvider struct {
}

func (p *BcryptProvider) Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	return string(bytes), err
}

func (p *BcryptProvider) Compare(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
