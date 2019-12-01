package models

import "github.com/jinzhu/gorm"

// setting-model Models
type DocType struct {
	gorm.Model
	Name string `json:"name" example:"สำนักเทคโนโลยีสารสนเทศ" ` // ชื่อศาล

}

func (DocType) TableName() string {
	return "tb_doctype"
}

func InitDoctypeData(db *gorm.DB) {
	data := []DocType{
		{Name: "แบบ ขงป 3-1 : สำหรับครุภัณฑ์แต่ละรายการ"},
		{Name: "แบบ ขงป 3-2 : สำหรับครุภัณฑ์แต่ละรายการ"},
		{Name: "แบบ 28 : สำหรับทุกหน่วยงาน"},
		{Name: "แบบ 29 : สำหรับทุกหน่วยงาน"},
	}

	for i := 0; i < len(data); i++ {
		db.Create(&data[i])
	}

}
