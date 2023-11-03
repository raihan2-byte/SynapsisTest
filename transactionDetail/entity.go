package transactiondetail

import (
	"synapsisid/product"
	"synapsisid/user"
	"time"
)

type TransactionDetails struct {
	ID         int
	Quantity   int
	UserID     int
	ProductID  int
	TotalPrice int
	Product    product.Products   `gorm:"foreignKey:ProductID"`
	User       user.User          `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}