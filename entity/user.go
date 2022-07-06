package entity

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type Users struct {
	ID_User    int `gorm:"primaryKey"`
	Nama       string
	Username   string
	Password   string
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
	Rent       []Rent    `gorm:"foreignKey:ID_Penyewa"`
	Book       []Books   `gorm:"foreignKey:Sumber_Buku"`
	// Rent       []Rent  `gorm:"many2many:users_Rent;"`
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

func (au *AksesUsers) LoginUser(newUsers Users) (result bool, id uint, err error) {
	// newUsers := Users{}
	// var Username, Password string
	fmt.Println("Silahkan Masukkan Username dan Password Anda")
	fmt.Print("Username: ")
	fmt.Scanln(&newUsers.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&newUsers.Password)
	result_ := au.DB.Where("Username = ? AND Password = ?", newUsers.Username, newUsers.Password).First(&newUsers)
	// err := ab.DB.Find(&daftarBuku)
	// fmt.Println(newUsers.ID_User)
	if result_.Error != nil {
		log.Fatal(result_.Statement.SQL.String())
		return false, 0, nil
	}
	return true, uint(newUsers.ID_User), nil
}

func (au *AksesUsers) LihatProfile(viewProfile Users) Users {
	err := au.DB.Select("ID_User", "Nama", "Username", "Password").Find(&viewProfile)
	// Where("ID_User = ?, Nama = ?, Username = ?, Password = ?").Scan(&viewProfile)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Users{}
	}
	return viewProfile
}
func (au *AksesUsers) EditProfile(id_user uint, EditProfile Users) Users {
	// var ID_User string
	// fmt.Println(ID_User)
	err := au.DB.Where("id_user = ?", id_user).Updates(&EditProfile)
	// Where("ID_User = ?, Nama = ?, Username = ?, Password = ?").Scan(&viewProfile)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Users{}
	}
	return EditProfile
}
