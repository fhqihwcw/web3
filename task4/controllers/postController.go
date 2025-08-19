package controllers

import (
	"task4/models"
	"time"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	BaseController
}

func (p PostController) CreatePost(c *gin.Context) {

	//创建文章
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		p.error(c, "Invalid input data")
		return
	}
	if err := models.DB.Create(&post).Error; err != nil {
		p.error(c, "Failed to create post: "+err.Error())
		return
	}

	p.success(c, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// 获取文章列表
func (p PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := models.DB.Find(&posts).Error; err != nil {
		p.error(c, "Failed to retrieve posts: "+err.Error())
		return
	}
	p.success(c, gin.H{
		"posts": posts,
	})
}

// 获取单个文章
func (p PostController) GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := models.DB.First(&post, id).Error; err != nil {
		p.error(c, "Post not found: "+err.Error())
		return
	}
	p.success(c, gin.H{
		"post": post,
	})
}

// 文章更新
func (p PostController) UpdatePost(c *gin.Context) {
	var tempPost models.Post

	if err := c.ShouldBindJSON(&tempPost); err != nil {
		p.error(c, "Invalid input data")
		return
	}

	var post models.Post

	if err := models.DB.First(&post, tempPost.ID).Error; err != nil {
		p.error(c, "Post not found: "+err.Error())
		return
	}

	post.Title = tempPost.Title
	post.Content = tempPost.Content
	post.UpdatedAt = time.Now()

	if err := models.DB.Save(&post).Error; err != nil {
		p.error(c, "Failed to update post: "+err.Error())
		return
	}
	p.success(c, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

// 删除文章
func (p PostController) DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := models.DB.First(&post, id).Error; err != nil {
		p.error(c, "Post not found: "+err.Error())
		return
	}
	if err := models.DB.Delete(&post).Error; err != nil {
		p.error(c, "Failed to delete post: "+err.Error())
		return
	}
	p.success(c, gin.H{
		"message": "Post deleted successfully",
	})
}
