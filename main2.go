package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "root:@/db1")
	defer db.Close()

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from data")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		fmt.Println("")
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
