package controllers

import (
	"task4/models"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	BaseController
}

func (cc CommentController) CreateComment(c *gin.Context) {
	var comments models.Comments
	if err := c.ShouldBindJSON(&comments); err != nil {
		cc.error(c, "Invalid input data")
		return
	}
	if err := models.DB.Create(&comments).Error; err != nil {
		cc.error(c, "Failed to create comment: "+err.Error())
		return
	}
	cc.success(c, gin.H{
		"status":  "success",
		"message": "Comment created successfully",
		"comment": comments,
	})
}

// 读取某篇文章的所有评论
func (cc CommentController) GetComments(c *gin.Context) {
	var comments []models.Comments
	if err := models.DB.Find(&comments, "post_id = ?", c.Param("id")).Error; err != nil {
		cc.error(c, "Failed to retrieve comments: "+err.Error())
		return
	}
	cc.success(c, gin.H{
		"status":   "success",
		"message":  "Comments retrieved successfully",
		"comments": comments,
	})
}
