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

	/*
	* ! Update
	 */
	stmt, _ := db.Prepare("update users set name = ? where id = ? ")

	stmt.Exec("Lucas", 1)
	stmt.Exec("Ted", 2)
	stmt.Exec("Bob", 3)

	/*
	* ! Delete
	 */
	stmt2, _ := db.Prepare("delete from users where id= ?")

	stmt2.Exec(3)

}
