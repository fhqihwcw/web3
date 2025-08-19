package controllers

import (
	"task4/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	BaseController
}

func (u UserController) Register(c *gin.Context) {
	// 处理用户注册逻辑
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.error(c, err.Error())
		return
	}
	//密码加密
	psd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.error(c, "Password encryption failed")
		return
	}
	user.Password = string(psd)
	// 保存用户到数据库
	if err := models.DB.Create(&user).Error; err != nil {
		u.error(c, "User registration failed: "+err.Error())
		return
	}
	// 返回成功响应
	u.success(c, "User Register successfully")
}

func (u UserController) Login(c *gin.Context) {
	// 处理用户登录逻辑
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.error(c, err.Error())
		return
	}
	// 查找用户
	var dbUser models.User
	if err := models.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		u.error(c, "User not found")
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		u.error(c, "Invalid password")
		return
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		u.error(c, "Failed to generate token: "+err.Error())
		return
	}
	// 返回成功响应
	u.success(c, gin.H{
		"token": tokenString,
	})
}
