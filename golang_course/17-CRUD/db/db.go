package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Connection driver to sql
)

// Connect open database connection
func Connect() (*sql.DB, error) {
	stringConnection := "root:pass@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
