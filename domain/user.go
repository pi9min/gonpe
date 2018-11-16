package domain

type Role uint8

const (
	RoleGuest Role = iota
	RoleAdmin
	RoleViewer
)

type User struct {
	ID                string
	Role              Role
	Email             string
	EncryptedPassword string
}
