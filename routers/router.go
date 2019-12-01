package routers

import (
	"rss-service/controllers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Healthcheck godoc
// @Summary ตรวจสอบสถานะการ online ของระบบ
// @Description ตรวจสอบสถานะการ online ของระบบ
// @Tags Healthcheck API
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthcheckModels
// @Router /healthcheck [get]
func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "rss service is online",
	})
}

// Healthcheck Models
type HealthcheckModels struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"rss service is online"`
}

// ApplyRoutes function
func ApplyRoutes(router *gin.Engine) {

	// file server
	router.Use(static.Serve("/api/v1/rss/files", static.LocalFile("uploads/v1/", false)))

	apiv1 := router.Group("/api/v1")

	// healthcheck servcie
	apiv1.GET("/healthcheck", Healthcheck)

	// swagger
	apiv1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setting := apiv1.Group("/setting", gin.BasicAuth(gin.Accounts{
		"rss": "wpassword",
	}))
	{
		setting.GET("", controllers.GetAllSettings)
		setting.GET("/:id", controllers.GetSetting)
		setting.POST("", controllers.CreateSetting)
		setting.PUT("/:id", controllers.UpdateSetting)
		setting.DELETE("/:id", controllers.DeleteSetting)
	}

}
