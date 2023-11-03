package cart

import (
	"gorm.io/gorm"
)

type RepositoryCart interface {
	FindAll() ([]Carts, error)
	Save(cart Carts) (Carts, error)
	FindById(ID int) (Carts, error)
	Update(cart Carts) (Carts, error)
	Delete(cart Carts) (Carts, error)
	FindAllByUserID (ID int) ([]Carts, error) 
	FindByProductAndUser(productID int, userID int) (Carts, error)

}

type repositoryCart struct {
	db *gorm.DB
}

func NewRepositoryCart(db *gorm.DB) *repositoryCart {
	return &repositoryCart{db}
}


func (r *repositoryCart) FindAll() ([]Carts, error) {
	var cart []Carts

	err := r.db.Preload("Product").Preload("User").Find(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}



func (r *repositoryCart) FindAllByUserID (ID int) ([]Carts, error)  {
	var cart []Carts

	err := r.db.Preload("Product").Preload("User").Where("user_id = ? ", ID).Find(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *repositoryCart) Save(cart Carts) (Carts, error) {
	err := r.db.Create(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil
}


func (r *repositoryCart) FindByProductAndUser(productID int, userID int) (Carts, error) {
	var cart Carts

	err := r.db.Where("product_id = ? ", productID).Where("user_id = ?", userID).Find(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repositoryCart) FindById(ID int) (Carts, error) {
	var cart Carts

	err := r.db.Where("id = ?", ID).Find(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repositoryCart) Update(cart Carts) (Carts, error) {
	err := r.db.Save(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil

}

func (r *repositoryCart) Delete(cart Carts) (Carts, error) {
	err := r.db.Delete(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}
