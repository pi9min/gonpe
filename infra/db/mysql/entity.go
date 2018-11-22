package mysql

import (
	"github.com/pi9min/gonpe/infra/db/mysql/user"
	"github.com/pi9min/gonpe/util/mysql"
)

var AllEntity = []mysql.EntityBehavior{
	&user.Entity{},
}
