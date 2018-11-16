package repository

import (
	"context"

	"github.com/ponpe/server/domain"
)

type User interface {
	Get(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
	Update(ctx context.Context, u *domain.User) error
}
