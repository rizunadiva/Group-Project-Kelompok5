package entity

import (
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
}

type AksesUsers struct {
	DB *gorm.DB
}

func (as *AksesUsers) TambahUser(newUsers Users) Users {
	err := as.DB.Create(&newUsers).Error
	if err != nil {
		log.Fatal(err)
		return Users{}
	}
	return newUsers
}
