package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      int    `db:"id"`
	PostID  int    `db:"post_id"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

func (c Comment) TableName() string {
	return "comments"
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("Comment AfterDelete:", c.ID)
	// 检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	err = tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comments_count", count).Error
	if err != nil {
		fmt.Println("Failed to update post comments count:", err)
	}
	return err
}
