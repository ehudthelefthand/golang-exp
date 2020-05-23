package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Text   string
	UserID uint
}

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3307)/eieidb2?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true) // Dev Only!

	err = db.AutoMigrate(
		&User{},
		&Comment{},
	).Error
	if err != nil {
		fmt.Println(err)
	}

	user := User{
		Name: "Pongneng",
		Comments: []Comment{
			Comment{Text: "Hello"},
			Comment{Text: "World"},
		},
	}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
	}

	if err := db.Delete(&user).Error; err != nil {
		fmt.Println(err)
	}
}
