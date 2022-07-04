package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/Kelompok5")
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi ke Server Berhasil!.")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
