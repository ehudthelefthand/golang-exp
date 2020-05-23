package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Doctor		Patient
// A			แดง
// A			ดำ
// C			ชาย

// ReportStupid DON'T DO IT THIS WAY!!!
// อย่าทำแบบนี้ พี่ขอร้อง พิมพ์เองยังเหนื่อยเลย
func ReportStupid(db *gorm.DB) {
	doctors := []Doctor{}
	if err := db.Find(&doctors).Error; err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(doctors); i++ {
		cases := []Cases{}
		err := db.
			Where("doctor_id = ?", doctors[i].ID).
			Find(&cases).Error
		if err != nil {
			fmt.Println(err)
		}

		for _, c := range cases {
			patient := Patient{}
			err := db.
				Where("id = ?", c.Patient_ID).
				First(&patient).Error
			if err != nil {
				fmt.Println(err)
			}
			doctors[i].Patients = append(doctors[i].Patients, patient)
		}
	}

	for _, doctor := range doctors {
		for _, patient := range doctor.Patients {
			fmt.Println(doctor.Name, patient.Name)
		}
	}
}
