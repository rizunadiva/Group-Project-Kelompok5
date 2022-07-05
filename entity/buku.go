package entity

import (
	"log"

	"gorm.io/gorm"
)

type Bukus struct {
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

func (ab *AksesBuku) GetAllData() []Bukus {
	var daftarBuku = []Bukus{}
	// err := as.DB.Raw("Select * from buku").Scan(&daftarBuku)
	err := ab.DB.Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBuku
}

func (ab *AksesBuku) AddBuku(newBuku Bukus) Bukus {
	err := ab.DB.Create(&newBuku).Error
	if err != nil {
		log.Fatal(err)
		return Bukus{}
	}
	return newBuku
}

// func (au *AksesUsers) TambahUser(newUsers Users) Users {
// 	err := au.DB.Create(&newUsers).Error
// 	if err != nil {
// 		log.Fatal(err)
// 		return Users{}
// 	}
// 	return newUsers
// }

func (ab *AksesBuku) HapusBuku(IDBuku int) bool {
	// tmp := Buku{ID: IDBuku}
	postExc := ab.DB.Where("ID = ?", IDBuku).Delete(&Bukus{})
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
