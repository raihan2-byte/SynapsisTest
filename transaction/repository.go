package transaction

import (
	"synapsisid/product"
	"synapsisid/user"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	FindAll() ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	FindById(ID int) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(transaction Transaction) (Transaction, error)
	FindByUserId(productID int, userID int) ([]Transaction, error)
}

type repositoryTransaction struct {
	db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) *repositoryTransaction {
	return &repositoryTransaction{db}
}

func (r *repositoryTransaction) FindAll() ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("Product").Preload("User").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repositoryTransaction) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) FindByUserId(productID int, userID int) ([]Transaction, error) {
	var transaction []Transaction


	err := r.db.Joins("Product", r.db.Where(&product.Products{ID: productID})).Joins("User", r.db.Where(&user.User{ID: userID})).Find(&transaction).Error


	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) FindById(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (r *repositoryTransaction) Delete(transaction Transaction) (Transaction, error) {
	err := r.db.Delete(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}


