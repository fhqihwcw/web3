package models

type Students struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Grade string `db:"grade"`
}

func (s Students) TableName() string {
	return "students"
}
