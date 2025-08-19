package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

func (c Comments) TableName() string {
	return "comments"
}
