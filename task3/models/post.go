package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID            int       `db:"id primary key auto_increment"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	Author        string    `db:"author"`
	UserID        int       `db:"user_id"`
	CommentsCount int       `db:"comments_count"`      // Foreign key to User
	Comments      []Comment `gorm:"foreignKey:PostID"` // Comments associated with the post
}

func (p Post) TableName() string {
	return "posts"
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("Post AfterCreate:", p.ID)
	userid := p.UserID
	if userid <= 0 {
		return fmt.Errorf("invalid user ID: %d", userid)
	}
	//查询当前用户的文章数量
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", userid).Count(&count)

	// 更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", userid).UpdateColumn("post_count", count).Error; err != nil {
		return fmt.Errorf("failed to update post count for user %d: %w", userid, err)
	}
	return nil
}

func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("Post AfterDelete:", p.ID)
	userid := p.UserID
	if userid <= 0 {
		return fmt.Errorf("invalid user ID: %d", userid)
	}
	//查询当前用户的文章数量
	var count int64
	tx.Model(&Post{}).Where("UserID = ?", userid).Count(&count)

	// 更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", userid).UpdateColumn("post_count", count+1).Error; err != nil {
		return fmt.Errorf("failed to update post count for user %d: %w", userid, err)
	}
	return nil
}
