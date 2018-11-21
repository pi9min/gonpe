package mysql

import (
	"github.com/ponpe/server/infra/db/mysql/user"
	"github.com/ponpe/server/util/mysql"
)

var AllEntity = []mysql.EntityBehavior{
	&user.Entity{},
}
