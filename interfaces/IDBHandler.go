package interfaces

import "database/sql"

type IDBHandler interface {
	Execute(statement string) (sql.Result, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
