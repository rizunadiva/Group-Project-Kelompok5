package buku

import (
	"fmt"
	_entity "group-project-kel5/entity"
	"log"

	"gorm.io/gorm"
)

type AksesBuku struct {
	DB *gorm.DB
}

func (ab *AksesBuku) AddBuku(newBuku _entity.Buku) _entity.Buku {
	var buku _entity.Buku
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
		return buku
	}
	return newBuku
}
