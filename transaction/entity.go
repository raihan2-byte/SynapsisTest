package transaction

import (
	"synapsisid/product"
	"synapsisid/user"
	"time"
)

//langsung beli tanpa ada masukkan cart
type Transaction struct {
	ID         int
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	// StatusPayment string
	Product    product.Products `gorm:"foreignKey:ProductID"`
	User       user.User        `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

