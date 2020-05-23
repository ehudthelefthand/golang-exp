package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Doctor struct {
	gorm.Model
	Name       string
	Specialist string
	Patients   []Patient
}

type Patient struct {
	gorm.Model
	Name string
}

type Cases struct {
	ID         uint `gorm:"primary_key"`
	Doctor_ID  uint
	Patient_ID uint
}

func (Cases) TableName() string {
	return "cases"
}

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3307)/eiei3?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true) // Dev Only!

	err = db.AutoMigrate(
		&Doctor{},
		&Patient{},
		&Cases{},
	).Error
	if err != nil {
		fmt.Println(err)
	}

	// Insert(db)
	// ReportBasic(db)
	ReportNormal(db)
}
