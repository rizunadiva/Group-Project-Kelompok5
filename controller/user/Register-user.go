package user

import (
	"fmt"
	_entity "group-project-kel5/entity"
	"log"

	"gorm.io/gorm"
)

type AksesUsers struct {
	DB *gorm.DB
}

// var Users _entity.Users{}
func (au *AksesUsers) TambahUser(newUsers _entity.Users) _entity.Users {
	// var newUsers _entity.Users
	var Users _entity.Users
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&newUsers.Nama)
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&newUsers.Username)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&newUsers.Password)

	err := au.DB.Create(&newUsers).Error
	if err != nil {
		log.Fatal(err)
		return Users
	}
	return newUsers
}
