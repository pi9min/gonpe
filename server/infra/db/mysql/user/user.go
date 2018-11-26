package user

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/pi9min/gonpe/server/domain"
	"github.com/pi9min/gonpe/server/proto"
	mysqlutil "github.com/pi9min/gonpe/server/util/mysql"
	"github.com/pi9min/gorp"
)

var _ mysqlutil.EntityBehavior = (*Entity)(nil)

const tableName = "user"

type Entity struct {
	ID           string       `db:"id"`
	Role         pb.User_Role `db:"role"`
	Email        string       `db:"email"`
	PasswordHash string       `db:"password_hash"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
}

func NewEntity(u *domain.User) *Entity {
	return &Entity{
		ID:           u.ID,
		Role:         u.Role,
		Email:        u.Email,
		PasswordHash: u.EncryptedPassword.String(),
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func (e *Entity) TableName() string {
	return tableName
}

func (e *Entity) PrimaryKey() mysqlutil.PrimaryKey {
	return mysqlutil.PrimaryKey{AutoIncrement: false, Columns: []string{"id"}}
}

func (e *Entity) Indexes() mysqlutil.Indexes {
	return mysqlutil.Indexes{}
}

func (e *Entity) Domain() *domain.User {
	return &domain.User{
		ID:                e.ID,
		Role:              e.Role,
		Email:             e.Email,
		EncryptedPassword: domain.NewEncryptedPassword(e.PasswordHash),
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

type repo struct{}

func NewRepository() *repo {
	return &repo{}
}

func (*repo) ExistByEmail(ctx context.Context, sqle gorp.SqlExecutor, email string) (bool, error) {
	var e Entity
	q := strings.Join([]string{
		"SELECT * FROM",
		tableName,
		"WHERE email=?",
	}, " ")
	if err := sqle.SelectOne(&e, q, email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (*repo) Get(ctx context.Context, sqle gorp.SqlExecutor, id string) (*domain.User, error) {
	var e Entity
	q := strings.Join([]string{
		"SELECT * FROM",
		tableName,
		"WHERE id=?",
	}, " ")
	if err := sqle.SelectOne(&e, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return e.Domain(), nil
}

func (*repo) GetAll(ctx context.Context, sqle gorp.SqlExecutor) ([]*domain.User, error) {
	var es []*Entity
	q := strings.Join([]string{
		"SELECT * FROM",
		tableName,
	}, " ")
	if _, err := sqle.Select(&es, q); err != nil {
		return nil, err
	}

	ds := make([]*domain.User, 0, len(es))
	for i := range es {
		d := es[i].Domain()
		ds = append(ds, d)
	}

	return ds, nil
}

func (*repo) Create(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error {
	e := NewEntity(u)
	if err := sqle.Insert(e); err != nil {
		return err
	}

	return nil
}

func (*repo) Update(ctx context.Context, sqle gorp.SqlExecutor, u *domain.User) error {
	e := NewEntity(u)
	if _, err := sqle.Update(e); err != nil {
		return err
	}

	return nil
}
