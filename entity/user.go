package entity

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int
	Nama     string
	Username string
	Password string
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
