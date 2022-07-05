package entity

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type Users struct {
	gorm.Model
	Nama       string
	Username   string
	Password   string
	Created_at string
	Updated_at string
	Book       []Buku `gorm:"foreignKey:Sumber_Buku"`
}

type AksesUsers struct {
	DB *gorm.DB
}

func (au *AksesUsers) TambahUser(newUsers Users) Users {
	// fmt.Print("Masukkan nama: ")
	// fmt.Scanln(&newUsers.Nama)
	// fmt.Print("Masukkan username: ")
	// fmt.Scanln(&newUsers.Username)
	// fmt.Print("Masukkan password: ")
	// fmt.Scanln(&newUsers.Password)

	err := au.DB.Create(&newUsers).Error
	if err != nil {
		log.Fatal(err)
		return Users{}
	}
	return newUsers
}

func (au *AksesUsers) LoginUser(bool, error) {
	newUsers := Users{}
	fmt.Println("Silahkan Masukkan Username dan Password Anda")
	fmt.Print("Username: ")
	fmt.Scanln(&newUsers.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&newUsers.Password)
	var daftarUser = []Users{}
	result, err := au.DB.Raw("SELECT username, password FROM users WHERE username = ? AND password = ?").Scan(&daftarUser)
	// err := ab.DB.Find(&daftarBuku)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return
	}
	return result, nil
}
