package main

import (
	"fmt"
	"log"
	"database/sql"
      _ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Printf("Hello, world.\n")

	db, err := sql.Open("mysql", "chris:passwd@tcp(127.0.0.1:3306)/test")
	if err != nil {
	log.Fatal(err)
	}
	defer db.Close()
	var str string
	q := "select uname from tbltest1 where id = 1"
	err = db.QueryRow(q).Scan(&str)
	if err != nil {
		log.Fatal(err)
		}
	log.Println(str)

}