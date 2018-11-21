package mysql

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

type client struct {
	gorp.SqlExecutor
}

func NewClient(dsn string) (*client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}
	dbMap.TraceOn("[SQL]", log.New(os.Stdout, "ponpe:", log.Lmicroseconds))

	// 各種初期化
	// table/indexなどを全て生成
	for i := range AllEntity {
		e := AllEntity[i]
		ei := reflect.Indirect(reflect.ValueOf(e)).Interface()
		tMap := dbMap.AddTableWithName(ei, e.TableName()).SetKeys(e.PrimaryKey().AutoIncrement, e.PrimaryKey().Columns...)
		indexes := e.Indexes()
		for j := range indexes {
			idx := indexes[j]
			tMap.AddIndex(idx.Name, "Btree", idx.Columns)
		}
	}

	if err := dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	if err := dbMap.CreateIndex(); err != nil {
		return nil, err
	}

	return &client{
		SqlExecutor: dbMap,
	}, nil
}
