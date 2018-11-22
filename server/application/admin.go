package application

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/application/repository"
	"github.com/pi9min/gonpe/server/domain"
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

func (a *AdminApp) CreateUser(ctx context.Context, email, password string, now time.Time) error {
	// TODO: validation(email, password length, etc...)
	u := domain.NewGuest(email, now)
	if err := u.SetPassword(password); err != nil {
		return err
	}

	return a.repos.User.Create(ctx, a.repos.MySQL, u)
}
