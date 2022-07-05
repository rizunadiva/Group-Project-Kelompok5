package entity

import (
	"log"

	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	ID           int
	Judul_Buku   string
	Penulis      string
	Penerbit     string
	Tahun_terbit string
	Sumber_Buku  int
}

type AksesBuku struct {
	DB *gorm.DB
}

func (as *AksesBuku) GetAllData() []Buku {
	var daftarBuku = []Buku{}
	err := as.DB.Raw("Select * from buku").Scan(&daftarBuku)
	// err := as.DB.Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBuku
}

func (as *AksesBuku) HapusBuku(IDBuku int) bool {
	// tmp := Buku{ID: IDBuku}
	postExc := as.DB.Where("ID = ?", IDBuku).Delete(&Buku{})
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
