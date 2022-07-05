package entity

import (
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
}

// type AksesUsers struct {
// 	DB *gorm.DB
// }

// func (au *AksesUsers) TambahUser(newUsers Users) Users {
// 	err := au.DB.Create(&newUsers).Error
// 	if err != nil {
// 		log.Fatal(err)
// 		return Users{}
// 	}
// 	return newUsers
// }
