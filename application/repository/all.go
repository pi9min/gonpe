package repository

import (
	"github.com/pi9min/gonpe/config"
	"github.com/pi9min/gonpe/infra/db/mysql"
	"github.com/pi9min/gonpe/infra/db/mysql/user"
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

	mysqlCli, err := mysql.NewClient(config.Gonpe.MysqlDsn)
	if err != nil {
		panic(err)
	}

	return &AllRepository{
		User:  user.NewRepository(),
		MySQL: mysqlCli,
	}
}
