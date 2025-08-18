package models

type Transactions struct {
	ID            int `db:"id"`
	FromAccountId int `db:"from_account_id"`
	ToAccountId   int `db:"to_account_id"`
	Amount        int `db:"amount"`
}

func (t Transactions) TableName() string {

	return "transactions"

}
