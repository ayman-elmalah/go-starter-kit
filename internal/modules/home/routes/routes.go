package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "go-starter-kit/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
}
