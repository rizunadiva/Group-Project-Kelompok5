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
	Book       []Bukus `gorm:"foreignKey:Sumber_Buku"`
}

type AksesUsers struct {
	DB *gorm.DB
}

func (au *AksesUsers) TambahUser(newUsers Users) Users {
	err := au.DB.Create(&newUsers).Error
	if err != nil {
		log.Fatal(err)
		return Users{}
	}
	return newUsers
}

func (au *AksesUsers) LoginUser(newUsers Users) (result bool, err error) {
	// newUsers := Users{}
	var Username, Password string
	fmt.Println("Silahkan Masukkan Username dan Password Anda")
	fmt.Print("Username: ")
	fmt.Scanln(&newUsers.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&newUsers.Password)
	// var daftarUser = []Users{}
	result_ := au.DB.Where("Username = ? AND Password = ?", Username, Password).Find(&newUsers)
	// err := ab.DB.Find(&daftarBuku)
	if result_.Error != nil {
		log.Fatal(result_.Statement.SQL.String())
		return false, nil
	}
	return true, nil
}
