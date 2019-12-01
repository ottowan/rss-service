package models

type DefaultData struct {
	Status int `json:"status"  example:"200" ` // รหัสสถานะการเรียก service
	Total  int `json:"total"  example:"1"`     // จำนวน Rows ทั้งหมด
}

type DefaultValidate struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type DefaultMessage struct {
	Status  int    `json:"status" example:"200" `
	Message string `json:"message" example:"message example" `
}
