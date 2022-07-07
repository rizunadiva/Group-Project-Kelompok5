package entity

import (
	"fmt"
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
	Rent         []Rent    `gorm:"foreignKey:ID_Buku"`
}

type AksesBuku struct {
	DB *gorm.DB
}

func (ab *AksesBuku) GetAllData() []Books {
	var daftarBuku = []Books{}
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

func (ab *AksesBuku) HapusBuku(ID_Buku int) bool {
	postExc := ab.DB.Where("ID_Buku = ?", ID_Buku).Delete(&Books{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	fmt.Println("====================")
	log.Println("Berhasil hapus buku")
	return true
}
func (ab *AksesBuku) EditBuku(id_book int, UpdateBuku Books) Books {
	err := ab.DB.Where("id_buku = ?", id_book).Updates(&UpdateBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Books{}
	}
	return UpdateBuku
}

func (ab *AksesBuku) GetMYData(id_user uint) []Books {
	var daftarBuku = []Books{}
	err := ab.DB.Where("Sumber_Buku =?", id_user).Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBuku
}
