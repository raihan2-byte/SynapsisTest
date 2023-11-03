package transaction

import "time"

type TransactionGetFormatter struct {
	ID         int       `json:"id"`
	ProductID  int       `json:"product_id"`
	UserID     int       `json:"user_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatterGet(transaction Transaction) TransactionGetFormatter {
	formatterGet := TransactionGetFormatter{}
	formatterGet.ID = transaction.ID
	formatterGet.ProductID = transaction.ProductID
	formatterGet.UserID = transaction.UserID
	formatterGet.Quantity = transaction.Quantity
	formatterGet.TotalPrice = transaction.TotalPrice
	formatterGet.CreatedAt = transaction.CreatedAt

	return formatterGet
}

func FormatterGetTransaction (transaction []Transaction) []TransactionGetFormatter {
	transactionGetFormatter := []TransactionGetFormatter{}

	for _, transactions := range transaction {
		beritaFormatter := FormatterGet(transactions)
		transactionGetFormatter = append(transactionGetFormatter, beritaFormatter)
	}

	return transactionGetFormatter
}

