package repository

import (
	"context"

	"firebase.google.com/go"
	"github.com/pi9min/gonpe/server/config"
	"github.com/pi9min/gonpe/server/infra/db/mysql"
	"github.com/pi9min/gonpe/server/infra/db/mysql/user"
	fb "github.com/pi9min/gonpe/server/infra/firebase"
	"github.com/pi9min/gorp"
)

type AllRepository struct {
	User
	MySQL    gorp.SqlExecutor
	Firebase *firebase.App
}

var repos *AllRepository

func NewAllRepository() *AllRepository {
	if repos != nil {
		return repos
	}

	ctx := context.Background()

	mysqlCli, err := mysql.NewClient(config.Gonpe.MysqlDsn)
	if err != nil {
		panic(err)
	}

	fbApp, err := fb.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	return &AllRepository{
		User:     user.NewRepository(),
		MySQL:    mysqlCli,
		Firebase: fbApp,
	}
}
