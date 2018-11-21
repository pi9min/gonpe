package repository

import (
	"context"

	"github.com/ponpe/server/domain"
	"gopkg.in/gorp.v2"
)

type User interface {
	Get(ctx context.Context, sqle gorp.SqlExecutor, id string) (*domain.User, error)
	GetAll(ctx context.Context, sqle gorp.SqlExecutor) ([]*domain.User, error)
	Create(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
	Update(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
}
