package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Doctor		Patient
// A			แดง
// A			ดำ
// C			ชาย
func ReportBasic(db *gorm.DB) {
	doctors := []Doctor{}
	if err := db.Find(&doctors).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", doctors)

	for _, doctor := range doctors {
		cases := []Cases{}
		if err := db.
			Where("doctor_id = ?", doctor.ID).
			Find(&cases).Error; err != nil {
			fmt.Println(err)
		}
		patients := []Patient{}
		for _, c := range cases {
			patient := Patient{}
			if err := db.
				Where("id = ?", c.Patient_ID).
				First(&patient).Error; err != nil {
				fmt.Println(err)
			}
			patients = append(patients, patient)
		}

		doctor.Patients = patients
	}

	for _, doctor := range doctors {
		for _, patient := range doctor.Patients {
			fmt.Printf("%s %s", doctor.Name, patient.Name)
		}
	}
}

type Report struct {
	Doctor  string
	Patient string
}

func ReportNormal(db *gorm.DB) {
	rows, err := db.Table("doctors").
		Select("doctors.name, patients.name").
		Joins("join cases on doctors.id = cases.doctor_id").
		Joins("join patients on patients.id = cases.patient_id").
		Rows()

	if err != nil {
		fmt.Println(err)
	}

	reports := []Report{}
	for rows.Next() {
		report := Report{}
		err := rows.Scan(&report.Doctor, &report.Patient)
		if err != nil {
			fmt.Println(err)
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	for _, r := range reports {
		fmt.Println(r.Doctor, r.Patient)
	}
}
