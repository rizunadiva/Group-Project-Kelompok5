package main

import (
	"fmt"
)

func main() {
	// conn := db.InitDB()
	// db.MigrateDB(conn)
	// aksesBook := entity.aksesBook{DB: conn}
	var input int = 0
	var input_1 int
	for input != 0 {
		fmt.Println("=============================================================")
		fmt.Println("|SELAMAT DATANG DI PERPUSTAKAAN UNIVERSITAS LANGSUNG SARJANA|")
		fmt.Println("=============================================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("0. Kembali")
		fmt.Println("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)

		switch input_1 {
		case 1:
			username, result, err := user.LoginUser(DBConn)
			if err != nil {
				fmt.Println("Login Gagal, username tidak terdaftar")
			}
			if result {
				fmt.Println("\n=Login /berhasil=")
				fmt.Println("\n=================")
				fmt.Println("Selamat Datang\nPilih menu dibawah ini")
				fmt.Println("1. Sewa Buku\n2. Tambah Buku")
			}
			// var
			// fmt.Print("Masukkan username: ")
			// fmt.Scanln(&)
			// fmt.Print("Masukkan password: ")
			// fmt.Scanln(&)
		}
	}
}
