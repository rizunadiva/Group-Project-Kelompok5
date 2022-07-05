package entity

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	Judul_Buku   string
	Penulis      string
	Penerbit     string
	Tahun_terbit string
	Sumber_Buku  int
	Created_at   string
	Updated_at   string
}

type AksesBuku struct {
	DB *gorm.DB
}

func (ab *AksesBuku) GetAllData() []Buku {
	var daftarBuku = []Buku{}
	// err := as.DB.Raw("Select * from buku").Scan(&daftarBuku)
	err := ab.DB.Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBuku
}

func (ab *AksesBuku) AddBuku(newBuku []Buku) []Buku {
	var tambahbuku = []Buku{}
	fmt.Print("Masukkan Judul Buku: ")
	fmt.Scanln(&newBuku.Judul_Buku)
	fmt.Print("Masukkan Penulis: ")
	fmt.Scanln(&newBuku.Penulis)
	fmt.Print("Masukkan Penerbit: ")
	fmt.Scanln(&newBuku.Penerbit)
	fmt.Print("Masukkan Tahun Penerbit: ")
	fmt.Scanln(&newBuku.Tahun_terbit)
	fmt.Print("Masukkan Sumber Buku: ")
	fmt.Scanln(&newBuku.Sumber_Buku)
	err := ab.DB.Create(&newBuku).Error
	if err != nil {
		log.Fatal(err)
		return tambahbuku
	}
	return newBuku
}

func (ab *AksesBuku) HapusBuku(IDBuku int) bool {
	// tmp := Buku{ID: IDBuku}
	postExc := ab.DB.Where("ID = ?", IDBuku).Delete(&Buku{})
	// err := as.DB.Delete(&tmp)
	//cek apakah postexc ada isinya ?
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	// berapa data yang berubah ?
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	return true
}

// func MigrateDB(conn *gorm.DB) {
// 	conn.AutoMigrate([]Buku)
// }
