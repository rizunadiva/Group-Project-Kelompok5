package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	ID_Rent      int `gorm:"primaryKey"`
	ID_Penyewa   int
	ID_Buku      int
	Tanggal_Sewa time.Time `gorm:"autoCreateTime"`
}

type AksesRent struct {
	DB *gorm.DB
}

func (ar *AksesRent) SewaBuku(Rental Rent) Rent {
	err := ar.DB.Create(&Rental).Error
	if err != nil {
		log.Fatal(err)
		return Rent{}
	}
	return Rental
}

func (ar *AksesRent) ReturnBuku(ID_Rent int) bool {
	postExc := ar.DB.Where("ID_Rent = ?", ID_Rent).Delete(&Rent{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak bisa return")
		return false
	}
	log.Println("Return Berhasil")
	return true
}
