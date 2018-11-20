package user

import (
	"context"

	"github.com/ponpe/server/domain"
)

type User struct {
	ID                string
	Role              domain.Role
	Email             string
	EncryptedPassword string
}

type repo struct{}

func NewRepository() *repo {
	return &repo{}
}

func (r *repo) Get(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (r *repo) Create(ctx context.Context, u *domain.User) error {
	return nil
}

func (r *repo) Update(ctx context.Context, u *domain.User) error {
	return nil
}
