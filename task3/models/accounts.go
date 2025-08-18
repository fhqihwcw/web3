package models

type Accounts struct {
	ID      int     `db:"id"`
	Balance float64 `db:"balance"`
}

func (a Accounts) TableName() string {
	return "accounts"
}
