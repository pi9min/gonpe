package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MysqlDsn string `required:"true" split_words:"true"`
}

var Ponpe Config

func Setup() {
	if err := envconfig.Process("ponpe", &Ponpe); err != nil {
		panic(err)
	}
}
