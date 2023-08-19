package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome go-starter-kit!",
	})
}
