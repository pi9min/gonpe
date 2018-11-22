package repository

import (
	"context"

	"github.com/pi9min/gonpe/server/domain"
	"github.com/pi9min/gorp"
)

type User interface {
	Get(ctx context.Context, sqle gorp.SqlExecutor, id string) (*domain.User, error)
	GetAll(ctx context.Context, sqle gorp.SqlExecutor) ([]*domain.User, error)
	Create(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
	Update(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
}
