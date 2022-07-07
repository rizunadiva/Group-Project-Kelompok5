package config

import (
	"group-project-kel5/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

// okk
// ...
func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:Mihrimah220918@tcp(localhost:3306)/Kelompok5?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		// panic(err)
		log.Fatal(err)
		return nil
	}
	return db
	// fmt.Println("Koneksi ke Server Berhasil!.")

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entity.Users{})
	conn.AutoMigrate(entity.Books{})
	conn.AutoMigrate(entity.Rent{})
}
