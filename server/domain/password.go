package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptedPassword struct {
	passwordHash string
}

func NewEmptyPassword() *EncryptedPassword {
	return &EncryptedPassword{passwordHash: ""}
}

func NewEncryptedPassword(passHash string) *EncryptedPassword {
	return &EncryptedPassword{passwordHash: passHash}
}

func (ep *EncryptedPassword) Set(pass string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return err
	}

	ep.passwordHash = string(hashed)
	return nil
}

func (ep *EncryptedPassword) Compare(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(ep.passwordHash), []byte(pass))
}

func (ep *EncryptedPassword) String() string {
	return ep.passwordHash
}
