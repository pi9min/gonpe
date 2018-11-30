package repository

import (
	"context"
	"encoding/base64"

	"firebase.google.com/go"
	"github.com/pi9min/gonpe/server/config"
	"github.com/pi9min/gonpe/server/infra/db/mysql"
	"github.com/pi9min/gonpe/server/infra/db/mysql/user"
	fb "github.com/pi9min/gonpe/server/infra/firebase"
	"gopkg.in/gorp.v2"
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

	serviceAccountKey, err := base64.StdEncoding.DecodeString(config.Gonpe.Base64FirebaseServiceAccountKey)
	if err != nil {
		panic(err)
	}
	fbApp, err := fb.NewApp(ctx, serviceAccountKey)
	if err != nil {
		panic(err)
	}

	return &AllRepository{
		User:     user.NewRepository(),
		MySQL:    mysqlCli,
		Firebase: fbApp,
	}
}
