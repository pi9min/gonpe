package mysql

type EntityBehavior interface {
	TableName() string
	PrimaryKey() PrimaryKey
	Indexes() Indexes
}

type PrimaryKey struct {
	AutoIncrement bool
	Columns       []string
}

type Indexes []Index

type Index struct {
	Name    string
	Columns []string
}
