package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Query(db *gorm.DB) {
	user := new(User)
	if err := db.Where("id = ?", 1).First(user).Error; err != nil {
		fmt.Println(err)
	}
	comments := []Comment{}
	err := db.Where("user_id = ?", user.ID).Find(&comments).Error
	if err != nil {
		fmt.Println(err)
	}
}
