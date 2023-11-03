package cart

import (
	"synapsisid/category"
	"synapsisid/product"
	"synapsisid/user"
	"time"
)

type Carts struct {
	ID         int
	Quantity   int
	// Price      int
	TotalPrice int
	ProductID  int 
	CategoryID int
	UserID     int
	Product product.Products  `gorm:"foreignKey:ProductID"`
	Category category.Categorys  `gorm:"foreignKey:CategoryID"`
	User user.User  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}