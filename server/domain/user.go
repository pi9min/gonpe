package domain

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pi9min/gonpe/server/proto"
	mysqlutil "github.com/pi9min/gonpe/server/util/mysql"
)

type User struct {
	ID                string
	Role              pb.User_Role
	Email             string
	EncryptedPassword *EncryptedPassword
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewGuest(email string, now time.Time) *User {
	return &User{
		ID:                mysqlutil.NewID(),
		Role:              pb.User_ROLE_GUEST,
		Email:             email,
		EncryptedPassword: NewEmptyPassword(),
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

func (u *User) SetPassword(pass string) error {
	return u.EncryptedPassword.Set(pass)
}

func (u *User) ComparePassword(pass string) error {
	return u.EncryptedPassword.Compare(pass)
}

func (u *User) Proto() (*pb.User, error) {
	cAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	uAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        u.ID,
		Role:      u.Role,
		Email:     u.Email,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}
