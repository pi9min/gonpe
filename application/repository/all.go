package repository

import (
	"github.com/ponpe/server/config"
	"github.com/ponpe/server/infra/db/mysql"
	"github.com/ponpe/server/infra/db/mysql/user"
	"gopkg.in/gorp.v2"
)

type AllRepository struct {
	User
	MySQL gorp.SqlExecutor
}

var repos *AllRepository

func NewAllRepository() *AllRepository {
	if repos != nil {
		return repos
	}

	mysqlCli, err := mysql.NewClient(config.Ponpe.MysqlDsn)
	if err != nil {
		panic(err)
	}

	return &AllRepository{
		User:  user.NewRepository(),
		MySQL: mysqlCli,
	}
}
