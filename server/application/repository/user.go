package repository

import (
	"context"

	"github.com/pi9min/gonpe/server/domain"
	"github.com/pi9min/gorp"
)

type User interface {
	ExistByEmail(ctx context.Context, sqle gorp.SqlExecutor, email string) (bool, error)
	Get(ctx context.Context, sqle gorp.SqlExecutor, id string) (*domain.User, error)
	GetAll(ctx context.Context, sqle gorp.SqlExecutor) ([]*domain.User, error)
	Create(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
	Update(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
}
