package models

import (
	"time"
)

// data-model Models
type DataSetting struct {
	DefaultData
	Data []Setting `json:"data" ` // ข้อมูล
}

// setting-model Models
type Setting struct {
	ID                   uint      `json:"id" gorm:"primary_key" example:"1"`                                                    // ID
	CourtCode            string    `json:"court_code" example:"0000011"`                                                         // รหัสศาล
	CourtName            string    `json:"court_name" example:"สำนักเทคโนโลยีสารสนเทศ" `                                         // ชื่อศาล
	SettingName          string    `json:"setting_name" example:"ตั้งค่า 1"`                                                     // ชื่อการตั้งค่าหนังสือ
	CourtAddr            string    `json:"court_addr"  example:"อาคารศาลอาญา ถนนรัชดาภิเษก แขวงจอมพล เขตจตุจักร กรุงเทพฯ 10900"` // ที่อยู่ของศาล
	LetterCourtNum       string    `json:"letter_court_num" example:"ศย102/62"`                                                  // เลขที่หนังสือ ศย
	LetterSignedName     string    `json:"letter_signed_name" example:"นายลายเซ็น หนังสือ"`                                      // ชื่อผู้เซ็นในหนังสือ
	LetterSignedPosition string    `json:"letter_signed_position" example:"นักวิจัดการงานทั่วไปชำนาญการ"`                        // ตำแหน่งผู้เซ็นในหนังสือ
	LetterOfficeName     string    `json:"letter_office_name" example:"สำนักเทคโนโลยีสารสนเทศ"`                                  // ชื่อหน่วยงานในหนังสือ
	LetterGroupName      string    `json:"letter_group_name" example:"ระบบงาน"`                                                  // ชื่อกลุ่มงานในหนังสือ
	LetterTel            string    `json:"letter_tel" example:"0255555"`                                                         // เบอร์โทรในหนังสือ
	LetterFax            string    `json:"letter_fax" example:"0266666"`                                                         // แฟกซ์ในหนังสือ
	LetterEmail          string    `json:"letter_email" example:"mail@coj.go.th"`                                                // อีเมล์ในหนังสือ
	CreatedUID           int       `json:"created_uid" example:"1"`                                                              // ID ผู้สร้างข้อมูล
	CreatedUname         string    `json:"created_uname" example:"user01"`                                                       // ชื่อผู้สร้างข้อมูล
	CreatedAt            time.Time `json:"created_at" example:"2019-10-15T14:01:58+07:00"`                                       // วันที่สร้างข้อมูล
	UpdatedUID           int       `json:"updated_uid" example:"1"`                                                              // ID ผู้แก้ไขข้อมูลล่าสุด
	UpdatedUname         string    `json:"updated_uname" example:"user01"`                                                       // ชื่อผู้แก้ไขข้อมูลล่าสุด
	UpdatedAt            time.Time `json:"updated_at" example:"2019-10-15T14:01:58+07:00"`                                       // วันที่แก้ไขข้อมูลล่าสุด
}

func (Setting) TableName() string {
	return "tb_letter_setting"
}

// var settingList = []Setting{
// 	{
// 		1, "court_code 1", "setting_name 1", "court_name 1", "court_addr 1", "letter_court_num 1", "letter_signed_name 1", "letter_signed_position 1", "letter_office_name 1", "letter_group_name 1", "letter_tel 1", "letter_fax 1", "letter_email 1", 1, "created_uname 1", time.Now(), 1, "updated_uname 1", time.Now(),
// 	},
// 	{
// 		1, "court_code 2", "setting_name 2", "court_name 2", "court_addr 2", "letter_court_num 2", "letter_signed_name 2", "letter_signed_position 2", "letter_office_name 2", "letter_group_name 2", "letter_tel 2", "letter_fax 2", "letter_email 2", 2, "created_uname 2", time.Now(), 2, "updated_uname 2", time.Now(),
// 	},
// }
