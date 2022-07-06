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
	// Tanggal_Kembali string `gorm:"autoCreateTime"`
	Users []Users `gorm:"foreignKey:ID_User;references:ID_Penyewa"`
	Books []Books `gorm:"foreignKey:ID_Buku;references:ID_Buku"`
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
