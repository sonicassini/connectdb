package main

import "fmt"
import "database/sql"
import "os"
import _ "github.com/go-sql-driver/mysql"

var user, pas, table, nama_table, exit string

func connect() *sql.DB {

	var db, err = sql.Open("mysql", user+":"+pas+"@/"+table)
	err = db.Ping()
	if err != nil {
		fmt.Println("gagal membuka database")
		fmt.Println("database tak bisa dihubungi")
		os.Exit(0)
	}
	return db
}

func outputsql() {
	var db = connect()
	defer db.Close()

	var id, nama string

	fmt.Println(" ")
	fmt.Println("menampilkan data yg diambil :")
	fmt.Println(" ")

	rows, _ := db.Query("select * from " + nama_table)

	for rows.Next() {
		rows.Scan(&id, &nama)
		fmt.Println("id :" + id)
		fmt.Println("nama :" + nama)
		fmt.Println(" ")

	}
	fmt.Println("exit press enter")
	fmt.Scanln(&exit)
}

func main() {
	fmt.Print("masukkan user mysql = ")
	fmt.Scanln(&user)
	fmt.Print("masukkan password mysql = ")
	fmt.Scanln(&pas)
	fmt.Print("masukan nama database = ")
	fmt.Scanln(&table)
	fmt.Print("masukan nama tabel = ")
	fmt.Scanln(&nama_table)
	fmt.Println(" ")
	outputsql()

}
