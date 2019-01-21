package db

import (
	"strings"
)

func init() {
	providedDialects := map[string]struct {
		dbType     DBType
		getDialect func() Dialect
	}{
		// "mssql":    {"mssql", func() core.Driver { return &odbcDriver{} }, func() core.Dialect { return &mssql{} }},
		// "odbc":     {"mssql", func() core.Driver { return &odbcDriver{} }, func() core.Dialect { return &mssql{} }}, // !nashtsai! TODO change this when supporting MS Access
		"mysql": {"mysql", func() Dialect { return &mysql{} }},
		// "mymysql":  {"mysql", func() core.Driver { return &mymysqlDriver{} }, func() core.Dialect { return &mysql{} }},
		// "postgres": {"postgres", func() core.Driver { return &pqDriver{} }, func() core.Dialect { return &postgres{} }},
		// "pgx":      {"postgres", func() core.Driver { return &pqDriverPgx{} }, func() core.Dialect { return &postgres{} }},
		"sqlite3": {"sqlite3", func() Dialect { return &sqlite3{} }},
		// "oci8":     {"oracle", func() core.Driver { return &oci8Driver{} }, func() core.Dialect { return &oracle{} }},
		// "goracle":  {"oracle", func() core.Driver { return &goracleDriver{} }, func() core.Dialect { return &oracle{} }},
	}

	for _, v := range providedDialects {
		RegisterDialect(v.dbType, v.getDialect)
	}
}

var (
	dialects = map[string]func() Dialect{}
)

// RegisterDialect register database dialect
func RegisterDialect(dbType DBType, dialectFunc func() Dialect) {
	if dialectFunc == nil {
		panic("sensor: Register dialect is nil")
	}
	dialects[strings.ToLower(string(dbType))] = dialectFunc
}

// QueryDialect query if registed database dialect
func QueryDialect(dbType DBType) Dialect {
	if d, ok := dialects[strings.ToLower(string(dbType))]; ok {
		return d()
	}
	return nil
}
