package application

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/application/repository"
	"github.com/pi9min/gonpe/server/domain"
	"github.com/pi9min/gonpe/server/proto"
)

type AdminApp struct {
	repos *repository.AllRepository
}

func NewAdminApp(repos *repository.AllRepository) *AdminApp {
	return &AdminApp{
		repos: repos,
	}
}

func (a *AdminApp) GetAllUser(ctx context.Context) ([]*domain.User, error) {
	return a.repos.User.GetAll(ctx, a.repos.MySQL)
}

func (a *AdminApp) ChangeRole(ctx context.Context, uID string, role pb.Role, now time.Time) error {
	u, err := a.repos.User.Get(ctx, a.repos.MySQL, uID)
	if err != nil {
		if err == domain.ErrNotFound {
			return ErrUserNotFound
		}
	}

	u.ChangeRole(role, now)
	if err := a.repos.User.Update(ctx, a.repos.MySQL, u); err != nil {
		return err
	}

	cli, err := a.repos.Firebase.Auth(ctx)
	if err != nil {
		return err
	}
	if err := cli.SetCustomUserClaims(ctx, u.AuthProviderUserID, u.FirebaseCustomUserClaims()); err != nil {
		return err
	}

	return nil
}
