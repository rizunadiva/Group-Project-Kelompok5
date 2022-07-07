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
	var viewProfile _entity.Users
	var EditProfile _entity.Users
	var idUserGlobal uint
	var editBuku _entity.Books
	// var UpdateBuku _entity.Books
	// var confirmDelete _entity.Users
	var DeleteUser _entity.Users
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
			result, idLogin, err := aksesUser.LoginUser(newUsers)
			if err != nil {
				fmt.Println("Login Gagal, username tidak terdaftar")
			}
			idUserGlobal = idLogin
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
					result1 := aksesUser.LihatProfile(idUserGlobal, viewProfile)
					fmt.Println("Profil Anda")
					fmt.Println("ID: ", result1.ID_User)
					fmt.Println("Nama: ", result1.Nama)
					fmt.Println("Username: ", result1.Username)
					fmt.Println("Password: ", result1.Password)
				case 2:
					fmt.Println("Masukkan Nama Baru: ")
					fmt.Scanln(&EditProfile.Nama)
					fmt.Println("Masukkan Username Baru: ")
					fmt.Scanln(&EditProfile.Username)
					fmt.Println("Masukkan Password: ")
					fmt.Scanln(&EditProfile.Password)

					result2 := aksesUser.EditProfile(idUserGlobal, EditProfile)
					fmt.Println("Nama Baru", result2.Nama)
					fmt.Println("Username Baru", result2.Username)
					fmt.Println("Password Baru", result2.Password)

				case 3:
					var confirmDel int
					fmt.Println("Anda yakin akan menghapus akun?")
					fmt.Println("1. Ya")
					fmt.Println("2. Tidak")
					fmt.Scanln(&confirmDel)
					switch confirmDel {
					case 1:
						// fmt.Scanln(&confirmDelete.ID_User)
						result3 := aksesUser.HapusProfile(idUserGlobal, DeleteUser)
						fmt.Print(result3)
					}

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
					var id_book int
					fmt.Println("ok")
					fmt.Print("Masukkan ID Buku yang akan diedit: ")
					fmt.Scanln(&id_book)
					fmt.Println("Masukkan Nama Judul Buku Baru: ")
					fmt.Scanln(&editBuku.Judul_Buku)
					fmt.Println("Masukkan Nama Penulis Baru: ")
					fmt.Scanln(&editBuku.Penulis)
					fmt.Println("Masukkan Nama Penerbit Baru: ")
					fmt.Scanln(&editBuku.Penerbit)
					fmt.Println("Masukkan Tahun Terbit Baru: ")
					fmt.Scanln(&editBuku.Tahun_terbit)
					fmt.Println("Masukkan ID Anda: ")
					fmt.Scanln(&editBuku.Sumber_Buku)

					result6 := aksesBuku.EditBuku(id_book, editBuku)
					fmt.Println("Judul Buku Baru", result6.Judul_Buku)
					fmt.Println("Penulis Buku Baru", result6.Penulis)
					fmt.Println("Penulis Buku Baru", result6.Penulis)
					fmt.Println("Penerbit Buku Baru", result6.Penerbit)
					fmt.Println("Tahun Terbit Baru", result6.Tahun_terbit)

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
