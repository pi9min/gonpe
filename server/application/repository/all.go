package repository

import (
	"github.com/pi9min/gonpe/server/config"
	"github.com/pi9min/gonpe/server/infra/db/mysql"
	"github.com/pi9min/gonpe/server/infra/db/mysql/user"
	"github.com/pi9min/gorp"
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

	mysqlCli, err := mysql.NewClient(config.Gonpe.MysqlDsn)
	if err != nil {
		panic(err)
	}

	return &AllRepository{
		User:  user.NewRepository(),
		MySQL: mysqlCli,
	}
}
