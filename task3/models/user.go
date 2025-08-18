package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int    `db:"id PRIMARY KEY  AUTO_INCREMENT"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	PostCount int    `db:"post_count"` // 文章数量
	Posts     []Post `gorm:"foreignKey:UserID"`
}

func (u User) TableName() string {
	return "users"
}
