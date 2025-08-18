package models

type Employees struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func (e Employees) TableName() string {
	return "employees"
}
