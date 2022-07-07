package entity

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
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

	fmt.Println("Silahkan Masukkan Username dan Password Anda")
	fmt.Print("Username: ")
	fmt.Scanln(&newUsers.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&newUsers.Password)
	result_ := au.DB.Where("Username = ? AND Password = ?", newUsers.Username, newUsers.Password).First(&newUsers)
	if result_.Error != nil {
		log.Fatal(result_.Statement.SQL.String())
		return false, 0, nil
	}
	return true, uint(newUsers.ID_User), nil
}

func (au *AksesUsers) LihatProfile(id_user uint) Users {
	var viewProfile Users
	err := au.DB.Select("id_user", "nama", "username", "password").First(&viewProfile, "id_user = ?", id_user)
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

func (au *AksesUsers) HapusProfile(id_user uint, DeleteUser Users) Users {
	err := au.DB.Where("id_user = ?", id_user).Delete(&DeleteUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Users{}
	}
	log.Println("Berhasil hapus akun")
	return DeleteUser
}

// var Books _entity.Books
// func (au *AksesUsers) daftarBukuSaya(id_user uint, rent []Rent, book []Books) []Rent {
// 	err := au.DB.Select("id_penyewa", "id_buku").First(&rent, "id_user = ?", id_user)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		return nil
// 	}
// 	return rent
// }
func (au *AksesUsers) MyBooks(id_user uint) []Rent {
	// var rent []Rent
	// var id_penyewa int
	var daftarBukuSaya []Rent
	err := au.DB.Where("id_penyewa= ?", id_user).Find(&daftarBukuSaya)

	// DB.Find(&daftarBukuSaya).First(&rent, "id_penyewa = ?", id_penyewa)
	// select from
	// err := au.DB.Joins("books").Find(&daftarBukuSaya)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBukuSaya
}

//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
//..
