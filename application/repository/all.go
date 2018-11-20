package repository

import "github.com/ponpe/server/infra/db/mysql/user"

type AllRepository struct {
	User
}

func NewAllRepository() *AllRepository {
	return &AllRepository{
		User: user.NewRepository(),
	}
}
