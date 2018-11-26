package application

import (
	"context"

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

func (a *AuthenticationApp) SignIn(ctx context.Context, email, password string) (string, error) {
	u, err := a.repos.User.GetByEmail(ctx, a.repos.MySQL, email)
	if err != nil {
		if err == domain.ErrNotFound {
			return "", ErrUserNotFound
		}
		return "", err
	}

	if err := u.ComparePassword(password); err != nil {
		return "", ErrPasswordNotMatch
	}

	return "token", nil
}
