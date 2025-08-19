package controllers

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (b *BaseController) success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

func (b *BaseController) error(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": message,
	})
}
