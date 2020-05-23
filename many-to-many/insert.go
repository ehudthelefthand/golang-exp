package main

import "github.com/jinzhu/gorm"

func Insert(db *gorm.DB) {
	doctor1 := Doctor{Name: "A", Specialist: "หมอตา"}
	db.Create(&doctor1)
	db.Create(&Doctor{Name: "B", Specialist: "หมอเด็ก"})
	db.Create(&Doctor{Name: "C", Specialist: "หมอฟัน"})

	db.Create(&Patient{Name: "แดง"})
	db.Create(&Patient{Name: "ดำ"})
	db.Create(&Patient{Name: "ชาย"})

	db.Create(&Cases{Doctor_ID: 1, Patient_ID: 1})
	db.Create(&Cases{Doctor_ID: 1, Patient_ID: 2})
	db.Create(&Cases{Doctor_ID: 3, Patient_ID: 3})
}