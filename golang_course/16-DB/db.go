package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlDb := "root:251526@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", urlDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão aberta")
}
