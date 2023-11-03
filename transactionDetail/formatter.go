package transactiondetail

import (
	"synapsisid/product"
	"synapsisid/user"
	"time"
)

type TransactionGetFormatter struct {
	ID         int
	Quantity   int
	UserID     int
	ProductID  int
	TotalPrice int
	Product    ProductResponse
	User       UserResponse
	CreatedAt  time.Time
}

type TransactionPostFormatter struct {
	ID         int
	Quantity   int
	UserID     int
	ProductID  int
	TotalPrice int
	Product    product.Products `gorm:"foreignKey:ProductID"`
	User       user.User        `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
}


type ProductResponse struct {
	ID         int
    Title      string
    Price      int
}

type UserResponse struct {
	ID        int
    FullName  string
}

func FormatterGet(transaction TransactionDetails) TransactionGetFormatter {
	formatterGet := TransactionGetFormatter{}
	formatterGet.ID = transaction.ID
	formatterGet.ProductID = transaction.ProductID
	formatterGet.UserID = transaction.UserID
	formatterGet.Quantity = transaction.Quantity
	formatterGet.TotalPrice = transaction.TotalPrice
	formatterGet.CreatedAt = transaction.CreatedAt

	formatterProduct := transaction.Product

	productFormatter := ProductResponse{}
	productFormatter.ID = formatterProduct.ID
	productFormatter.Price = formatterProduct.Price

	formatterGet.Product = productFormatter

	formatterUser := transaction.User
	userFormatter := UserResponse{}
	userFormatter.ID = formatterUser.ID
	userFormatter.FullName = formatterUser.FullName

	formatterGet.User = userFormatter

	return formatterGet
}

func FormatterGetTransaction(transaction []TransactionDetails) []TransactionGetFormatter {
	transactionGetFormatter := []TransactionGetFormatter{}

	for _, transactions := range transaction {
		beritaFormatter := FormatterGet(transactions)
		transactionGetFormatter = append(transactionGetFormatter, beritaFormatter)
	}

	return transactionGetFormatter
}
