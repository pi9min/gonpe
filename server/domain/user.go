package domain

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pi9min/gonpe/server/proto"
	mysqlutil "github.com/pi9min/gonpe/server/util/mysql"
)

type User struct {
	ID                 string
	AuthProviderUserID string
	Role               pb.Role
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func NewGuest(uID string, now time.Time) *User {
	return &User{
		ID:                 mysqlutil.NewID(),
		AuthProviderUserID: uID,
		Role:               pb.Role_ROLE_GUEST,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
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
		Id:                 u.ID,
		AuthProviderUserId: u.AuthProviderUserID,
		Role:               u.Role,
		CreatedAt:          cAt,
		UpdatedAt:          uAt,
	}, nil
}

func (u *User) ChangeRole(r pb.Role, now time.Time) {
	u.Role = r
	u.updateTimestamp(now)
}

func (u *User) FirebaseCustomUserClaims() map[string]interface{} {
	claims := map[string]interface{}{
		"role": u.Role.String(),
	}
	return claims
}

func (u *User) updateTimestamp(now time.Time) {
	u.UpdatedAt = now
}
