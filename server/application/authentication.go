package application

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/application/repository"
	"github.com/pi9min/gonpe/server/domain"
)

type AuthenticationApp struct {
	repos *repository.AllRepository
}

func NewAuthenticationApp(repos *repository.AllRepository) *AuthenticationApp {
	return &AuthenticationApp{
		repos: repos,
	}
}

func (a *AuthenticationApp) RegisterGuestUser(ctx context.Context, authProviderUserID string, now time.Time) error {
	u := domain.NewGuest(authProviderUserID, now)
	if err := a.repos.User.Create(ctx, a.repos.MySQL, u); err != nil {
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
