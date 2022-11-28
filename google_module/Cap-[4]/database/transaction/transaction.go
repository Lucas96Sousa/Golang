package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	/*
	*Tratando o erro
	 */

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
