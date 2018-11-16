package user

import "github.com/ponpe/server/domain"

type User struct {
	ID                string
	Role              domain.Role
	Email             string
	EncryptedPassword string
}
