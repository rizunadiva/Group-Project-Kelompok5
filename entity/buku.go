package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID_Buku      int `gorm:"primaryKey"`
	Judul_Buku   string
	Penulis      string
	Penerbit     string
	Tahun_terbit string
	Sumber_Buku  int
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoCreateTime"`
	// Rent         []Rent `gorm:"many2many:books_Rent;"`
}

type AksesBuku struct {
	DB *gorm.DB
}

func (ab *AksesBuku) GetAllData() []Books {
	var daftarBuku = []Books{}
	// err := as.DB.Raw("Select * from books").Scan(&daftarBuku)
	err := ab.DB.Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBuku
}

func (ab *AksesBuku) AddBuku(newBuku Books) Books {
	err := ab.DB.Create(&newBuku).Error
	if err != nil {
		log.Fatal(err)
		return Books{}
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
	postExc := ab.DB.Where("ID = ?", IDBuku).Delete(&Books{})
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
