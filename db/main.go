package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique_index"`
	Password string `gorm:"not null"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3307)/eieidb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true) // Dev Only!
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("สวัสดี")

	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	// Insert
	user := User{
		Username: "pongneng",
		Password: "password",
	}
	if err := createUser(&user); err != nil {
		log.Println(err)
	}

}
