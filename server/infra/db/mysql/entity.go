package mysql

import (
	"github.com/pi9min/gonpe/server/infra/db/mysql/user"
	"github.com/pi9min/gonpe/server/util/mysql"
)

var AllEntity = []mysql.EntityBehavior{
	&user.Entity{},
}
