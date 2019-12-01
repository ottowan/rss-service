package controllers

import (
	"rss-service/include"
	"rss-service/models"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type Setting = models.Setting
type DataSetting = models.DataSetting

// Get Settings godoc
// @Summary เรียกดูรายการตั้งค่าหนังสือ
// @Description เรียกดูรายการตั้งค่าหนังสือของศาล
// @Tags ตั้งค่าหนังสือ
// @Accept  json
// @Produce  json
// @Param court_code query string false "รหัสศาล"
// @Success 200 {object} models.DataSetting
// @Router /setting [get]
func GetAllSettings(c *gin.Context) {
	db := include.GetDB()
	var data DataSetting
	var setting []Setting
	var count int

	//Get court from query
	courtID := c.DefaultQuery("court_code", "")

	query := db

	if courtID != "" {
		query = query.Where("court_code = ?", courtID)
	}

	if err := query.Find(&setting).Count(&count).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": err.Error(),
		})
		return
	} else {

		data.Total = count
		data.Data = setting
		data.Status = 200

		c.JSON(200, data)
	}

}

// Get Settings godoc
// @Summary ดูข้อมูลการตั้งค่าหนังสือ
// @Description ดูข้อมูลการตั้งค่าหนังสือของศาล
// @Tags ตั้งค่าหนังสือ
// @Accept  json
// @Produce  json
// @Param id path  int true  "ID"
// @Success 200 {object} models.Setting
// @Router /setting/{id} [get]
func GetSetting(c *gin.Context) {
	db := include.GetDB()
	id := c.Params.ByName("id")
	var setting Setting

	if err := db.Where("id = ? ", id).First(&setting).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": err.Error(),
		})

	} else {

		c.JSON(200, gin.H{
			"status": 200,
			"data":   setting,
		})
	}
}

// CreateSetting godoc
// @Summary เพิ่มข้อมูลการตั้งค่าหนังสือ
// @Description เพิ่มข้อมูลการตั้งค่าหนังสือ <br>*** Required <br>"court_code", "setting_name", "created_uid", "created_uname"
// @Tags ตั้งค่าหนังสือ
// @Accept  json
// @Produce  json
// @Param setting body models.Setting true "Create Setting"
// @Success 201 {object} models.Setting
// @Router /setting [post]
func CreateSetting(c *gin.Context) {
	db := include.GetDB()

	var setting Setting
	var defaultValidate = []models.DefaultValidate{}
	var count int

	c.ShouldBind(&setting)

	// validate
	valid := validation.Validation{}
	valid.Required(&setting.CourtCode, "court_code").Message("court_code is required")
	valid.Required(&setting.SettingName, "setting_name").Message("setting_name is required")
	valid.Required(&setting.CreatedUID, "created_uid").Message("created_uid is number")
	valid.Required(&setting.CreatedUname, "created_uname").Message("created_uname is required")

	if !valid.HasErrors() {

		// count จำนวน rows แต่ละศาล
		if err := db.Model(&Setting{}).Where("court_code = ?", setting.CourtCode).Count(&count).Error; err != nil {
			c.JSON(404, gin.H{
				"status":  404,
				"message": err.Error(),
			})
			return
		}

		// แต่ละศาลห้ามเกิน 10 rows
		if count > 9 {
			c.JSON(409, gin.H{
				"status":  409,
				"message": "data must not exceed 10 rows",
			})
			return
		}

		if err := db.Create(&setting).Error; err != nil {
			c.JSON(404, gin.H{
				"status":  404,
				"message": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"status": 201,
				"data":   setting,
			})
		}
	} else {
		for _, err := range valid.Errors {
			defaultValidate = append(defaultValidate, models.DefaultValidate{Key: err.Key, Message: err.Message})
		}

		c.JSON(400, gin.H{
			"status": 400,
			"errors": defaultValidate,
		})
	}

}

// UpdateSetting godoc
// @Summary แก้ไขข้อมูลการตั้งค่าหนังสือ
// @Description แก้ไขข้อมูลการตั่งค่าหนังสือ <br>*** Required <br>"updated_uid", "updated_uname"
// @Tags ตั้งค่าหนังสือ
// @Accept  json
// @Produce  json
// @Param id path  int true  "ID"
// @Param setting body models.Setting true "Update Setting"
// @Success 200 {object} models.Setting
// @Router /setting/{id} [put]
func UpdateSetting(c *gin.Context) {
	db := include.GetDB()
	var setting Setting
	var defaultValidate = []models.DefaultValidate{}
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&setting).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": err.Error(),
		})
	}

	c.ShouldBind(&setting)

	// validate
	valid := validation.Validation{}
	valid.Required(&setting.UpdatedUID, "updated_uid").Message("updated_uid is number")
	valid.Required(&setting.UpdatedUname, "updated_uname").Message("updated_uname is required")

	if !valid.HasErrors() {
		if err := db.Save(&setting).Error; err != nil {
			c.JSON(404, gin.H{
				"status":  404,
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"status": 200,
				"data":   setting,
			})
		}
	} else {
		for _, err := range valid.Errors {
			defaultValidate = append(defaultValidate, models.DefaultValidate{Key: err.Key, Message: err.Message})
		}

		c.JSON(400, gin.H{
			"status": 400,
			"errors": defaultValidate,
		})
	}

}

// DeleteSetting godoc
// @Summary ลบข้อมูลการตั้งค่าหนังสือ
// @Description ลบข้อมูลการตั่งค่าหนังสือ
// @Tags ตั้งค่าหนังสือ
// @Accept  json
// @Produce  json
// @Param id path  int true  "ID"
// @Success 200 {object} models.DefaultMessage
// @Router /setting/{id} [delete]
func DeleteSetting(c *gin.Context) {
	db := include.GetDB()
	id := c.Params.ByName("id")
	var setting Setting

	if err := db.Where("id = ? ", id).Delete(&setting).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Deleted Success!",
		})
	}
}
