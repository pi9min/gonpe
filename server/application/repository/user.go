package repository

import (
	"context"

	"github.com/pi9min/gonpe/server/domain"
	"gopkg.in/gorp.v2"
)

type User interface {
	ExistByAuthProviderUserID(ctx context.Context, sqle gorp.SqlExecutor, uID string) (bool, error)
	Get(ctx context.Context, sqle gorp.SqlExecutor, id string) (*domain.User, error)
	GetByAuthProviderUserID(ctx context.Context, sqle gorp.SqlExecutor, uID string) (*domain.User, error)
	GetAll(ctx context.Context, sqle gorp.SqlExecutor) ([]*domain.User, error)
	Create(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
	Update(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error
}
