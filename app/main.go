package main

import (
	"fmt"
	"group-project-kel5/config"
	_entity "group-project-kel5/entity"
)

func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUser := _entity.AksesUsers{DB: conn}
	aksesBuku := _entity.AksesBuku{DB: conn}
	aksesRent := _entity.AksesRent{DB: conn}
	var input int = 0
	var newUsers _entity.Users
	var newBuku _entity.Books
	var Rental _entity.Rent
	// var viewProfile _entity.Users
	for input != 99 {
		fmt.Println("=============================================================")
		fmt.Println("|SELAMAT DATANG DI PERPUSTAKAAN UNIVERSITAS LANGSUNG SARJANA|")
		fmt.Println("=============================================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("99. Kembali")
		fmt.Println("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			// AksesLogin := _entity.AksesLogin{DB: conn}
			result, err := aksesUser.LoginUser(newUsers)
			if err != nil {
				fmt.Println("Login Gagal, username tidak terdaftar")
			}
			if result {
				var input2 int
				fmt.Println("\n=Login /berhasil=")
				fmt.Println("\n=================")
				fmt.Println("Selamat Datang\nPilih menu dibawah ini")
				fmt.Println("1. Lihat Profil\n2. Edit Profil\n3. Hapus Profil\n4. Tambah Buku\n5. Sewa Buku\n6. Edit Buku\n7. Hapus Buku\n8. Kembalikan Buku\n9. Buku Saya")
				fmt.Print("Pilih Menu : ")
				fmt.Scanln(&input2)
				switch input2 {
				case 1:
					fmt.Println("Profil Anda")
					// for _, val := range aksesUser.LihatProfile(viewProfile) {
					// fmt.Println(val)
				// }
				case 2:
					fmt.Println("ok")
				case 3:
					fmt.Println("ok")
				case 4:
					fmt.Print("Masukkan Judul Buku: ")
					fmt.Scanln(&newBuku.Judul_Buku)
					fmt.Println("Masukkan Penulis: ")
					fmt.Scanln(&newBuku.Penulis)
					fmt.Print("Masukkan Penerbit: ")
					fmt.Scanln(&newBuku.Penerbit)
					fmt.Print("Masukkan Tahun Penerbit: ")
					fmt.Scanln(&newBuku.Tahun_terbit)
					fmt.Print("Masukkan ID anda: ")
					fmt.Scanln(&newBuku.Sumber_Buku)

					resultBook := aksesBuku.AddBuku(newBuku)
					if resultBook.ID_Buku == 0 {
						fmt.Println("Tidak bisa tambah buku")
						break
					}
					fmt.Println("Berhasil Input User")
				case 5:
					fmt.Print("Masukkan ID anda: ")
					fmt.Scanln(&Rental.ID_Penyewa)
					fmt.Print("Masukkan ID buku: ")
					fmt.Scanln(&Rental.ID_Buku)

					resultSewa := aksesRent.SewaBuku(Rental)
					if resultSewa.ID_Rent == 0 {
						fmt.Println("Tidak bisa sewa buku")
						break
					}
					fmt.Println("Berhasil Sewa")
				case 6:
					fmt.Println("ok")
				case 7:
					var ID_Buku int
					fmt.Print("Masukkan ID Buku yang akan dihapus: ")
					fmt.Scanln(&ID_Buku)
					fmt.Println(aksesBuku.HapusBuku(ID_Buku))
				case 8:
					var ID_Rent int
					fmt.Print("Masukkan ID Sewa yang akan dikembalikan: ")
					fmt.Scanln(&ID_Rent)
					fmt.Println(aksesRent.ReturnBuku(ID_Rent))
				case 9:
					fmt.Println("ok")
				}
			}

		case 2:

			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newUsers.Nama)
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&newUsers.Username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&newUsers.Password)

			result := aksesUser.TambahUser(newUsers)
			if result.ID_User == 0 {
				fmt.Println("Tidak bisa input user")
				// break
			}
			fmt.Println("Berhasil input user")
		case 3:
			// fmt.Println("ok")
			fmt.Println("Berikut adalah daftar seluruh buku")
			// rescase3 := aksesBuku.GetAllData(daftarBuku)
			for _, val := range aksesBuku.GetAllData() {
				fmt.Println(val)
			}
		}
	}
	fmt.Println("Terima Kasih")
}
