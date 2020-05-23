package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	// ID       uint   `gorm:"primary_key"`
	gorm.Model
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
	defer db.Close()
	db.LogMode(true) // Dev Only!
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("สวัสดี")

	err = db.AutoMigrate(
		&User{},
	).Error
	if err != nil {
		log.Fatal(err)
	}

	// Insert
	// user := User{
	// 	Username: "pongneng4",
	// 	Password: "password",
	// }
	// if err := createUser(&user); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("user is created with id = %v", user.ID)

	// Read
	// user1, err := GetUserByID(4)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("before update: %#v\n", user1)

	// Update
	// user1.Username = "pongneng4"
	// user1.Password = "password2"
	// if err := updateUser(user1); err != nil {
	// 	log.Println(err)
	// }

	// Delete
	// user2, err := GetUserByID(4)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("after update: %#v\n", user2)

	// if err := deleteUser(user2.ID); err != nil {
	// 	fmt.Println(err)
	// }
}
