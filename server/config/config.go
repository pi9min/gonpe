package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MysqlDsn                        string `required:"true" split_words:"true"`
	Base64FirebaseServiceAccountKey string `required:"true" split_words:"true"`
}

var Gonpe Config

func Setup() {
	if err := envconfig.Process("gonpe", &Gonpe); err != nil {
		panic(err)
	}
}
