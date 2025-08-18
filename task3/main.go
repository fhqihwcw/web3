package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fhqihwcw/web3/task3/models"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var dbx *sqlx.DB

func init() {
	// 连接数据库
	DB, err = gorm.Open(mysql.Open("root:zhaoxj123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败")
	}

	//sqlx 连接数据库
	dbx, err = sqlx.Connect("mysql", "root:zhaoxj123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("sqlx连接数据库失败")
	}
}

func main() {

	// basicCRUD()

	// accounts1 := models.Accounts{ID: 1, Balance: 0}
	// accounts2 := models.Accounts{ID: 2, Balance: 0}
	// DB.Create(accounts1)
	// DB.Create(accounts2)

	//transactions()

	// for i := 11; i < 20; i++ {
	// 	employee := models.Employees{
	// 		Name:       "员工" + strconv.Itoa(i),
	// 		Department: "技术部",
	// 		Salary:     10000 + float64(i)*200,
	// 	}
	// 	DB.Create(&employee)
	// }

	queries()
	queriesBooks()

	// createModels()

	// queriesWithRelations()

	// hookstest()
}

/*
*

	*SQL语句练习

题目1：基本CRUD操作 (题目1.sql目录)
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*
*/
func basicCRUD() {
	// 插入一条新记录
	// DB.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", "张三", 20, "三年级")
	DB.Create(&models.Students{Name: "李四", Age: 22, Grade: "三年级"})
	// 查询所有年龄大于18岁的学生信息
	students := []models.Students{}
	DB.Where("age > ?", 18).Find(&students)
	data, _ := json.Marshal(students)
	fmt.Println("Students older than 18:", string(data))
	// 更新姓名为"张三"的学生年级为"四年级"
	DB.Model(&models.Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	// 删除年龄小于15岁的学生记录
	DB.Where("age < ?", 15).Delete(&models.Students{})
}

/*
*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*
*/

func transactions() {
	// 事务开始
	tx := DB.Begin()
	// 假设账户 A 和 B 的 ID 分别为 1 和 2，转账金额为 100
	accountsA := []models.Accounts{}
	tx.Where("ID = ?", 1).Find(&accountsA)
	if len(accountsA) == 0 || accountsA[0].Balance < 100 {
		tx.Rollback()
		return
	}
	// 扣除账户 A 的余额
	accountsA[0].Balance -= 100
	err := tx.Save(&accountsA[0]).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 增加账户 B 的余额
	accountsB := []models.Accounts{}
	tx.Where("ID = ?", 2).Find(&accountsB)
	if len(accountsB) == 0 {
		tx.Rollback()
		return
	}
	accountsB[0].Balance += 100
	err = tx.Save(&accountsB[0]).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 记录转账信息
	transaction := models.Transactions{
		FromAccountId: accountsA[0].ID,
		ToAccountId:   accountsB[0].ID,
		Amount:        100,
	}
	err = tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

}

/**
Sqlx入门
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
**/

func queries() {
	// 查询所有技术部员工
	employees := []models.Employees{}
	dbx.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	// 打印查询结果
	for _, employee := range employees {
		println("ID:", employee.ID, "Name:", employee.Name, "Department:", employee.Department, "Salary:", employee.Salary)
	}
	// 查询工资最高的员工
	highestSalaryEmployee := models.Employees{}
	dbx.Get(&highestSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	println("ID:", highestSalaryEmployee.ID, "Name:", highestSalaryEmployee.Name, "Department:", highestSalaryEmployee.Department, "Salary:", highestSalaryEmployee.Salary)
}

/**
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
**/

func queriesBooks() {
	// 查询价格大于50元的书籍
	books := []models.Book{}
	dbx.Select(&books, "SELECT * FROM books WHERE price > ?", 50)

	// 打印查询结果
	for _, book := range books {
		println("ID:", book.ID, "Title:", book.Title, "Author:", book.Author, "Price:", strconv.FormatFloat(book.Price, 'f', 2, 64))
	}
}

/*
*
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*
*/
func createModels() {
	// 创建 User 模型
	DB.AutoMigrate(&models.User{})
	// 创建 Post 模型
	DB.AutoMigrate(&models.Post{})
	// 创建 Comment 模型
	DB.AutoMigrate(&models.Comment{})

	// user := models.User{Username: "zhangsan"}
	// DB.Create(&user)

	// post1 := models.Post{Title: "第一篇文章", Content: "这是我的第一"}
	// post2 := models.Post{Title: "第二篇文章", Content: "这是我的第二篇文章"}
	// post1.UserID = int(user.ID)
	// post2.UserID = int(user.ID)
	// DB.Create(&post1)
	// DB.Create(&post2)

	// comment1 := models.Comment{PostID: int(post1.ID), Content: "评论1", Author: "评论者1"}
	// comment2 := models.Comment{PostID: int(post1.ID), Content: "评论2", Author: "评论者2"}
	// DB.Create(&comment1)
	// DB.Create(&comment2)
}

/**
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
**/

func queriesWithRelations() {
	// 查询某个用户发布的所有文章及其对应的评论信息
	userID := 1 // 假设要查询的用户ID为1
	user := models.User{}
	DB.Preload("Posts.Comments").First(&user, userID)
	data, _ := json.Marshal(user)
	fmt.Println("User:", string(data))

	// 查询评论数量最多的文章信息
	post := models.Post{}
	DB.Select("posts.*, COUNT(comments.id) as comments_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comments_count DESC").
		First(&post)
	data, _ = json.Marshal(post)
	fmt.Println("Post with most comments:", string(data))

}

/*
*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func hookstest() {
	// 钩子函数已经在 Post 模型的 BeforeCreate 方法中实现了自动更新用户的文章数量统计字段
	post1 := models.Post{Title: "第一篇文章3", Content: "这是我的第一篇文章3"}
	post1.UserID = int(1)
	DB.Create(&post1)

	// 钩子函数已经在 Comment 模型的 AfterDelete 方法中实现了检查文章评论数量并更新文章状态
	comment := models.Comment{PostID: 1, ID: 3}
	DB.Delete(&comment)

}
