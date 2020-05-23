package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// ReportSmart This is the way. Please do it.
// ทำแบบนี้นะ จะได้ไม่เป็นภาระของลูกหลาน
func ReportSmart(db *gorm.DB) {
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
