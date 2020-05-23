package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Insert(db *gorm.DB) {
	// ------- Create -------

	user := User{
		Name: "Pongneng",
	}
	comments := []Comment{
		Comment{Text: "Hello"},
		Comment{Text: "World"},
	}

	// insert into users ... => user_id
	// insert into comments with user_id
	// insert into comments with user_id

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	// panic("crashh!!!")

	for _, comment := range comments {
		comment.UserID = user.ID
		if err := tx.Create(&comment).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
		}
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
	}

	// err = db.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.Create(&user).Error; err != nil {
	// 		return err
	// 	}
	// 	for _, comment := range comments {
	// 		comment.UserID = user.ID
	// 		if err := tx.Create(&comment).Error; err != nil {
	// 			return err
	// 		}
	// 	}
	// 	return nil
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

}
