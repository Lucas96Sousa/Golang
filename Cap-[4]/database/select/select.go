package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "user:pass@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id > ? ", 5)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}

}
