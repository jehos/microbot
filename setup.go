package microbot

import (
	"database/sql"
	"errors"

	"github.com/elvinchan/microbot/db"
)

var dialects []db.Dialect

func RegisterDB(d *sql.DB, dbType db.DBType) error {
	if d == nil {
		return errors.New("nil sql.DB")
	}
	dialect := db.QueryDialect(dbType)
	if dialect == nil {
		return errors.New("DBType not support")
	}
	dialect.Init(d, dbType)
	dialects = append(dialects, dialect)
	return nil
}
