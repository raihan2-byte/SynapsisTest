package transactiondetail

import (
	"synapsisid/product"
	"synapsisid/user"

	"gorm.io/gorm"
)

type RepositoryTransactionDetails interface {
	FindAll() ([]TransactionDetails, error)
	Save(transaction TransactionDetails) (TransactionDetails, error)
	FindById(ID int) (TransactionDetails, error)
	Update(transaction TransactionDetails) (TransactionDetails, error)
	Delete(transaction TransactionDetails) (TransactionDetails, error)
	FindByUserId(productID int, userID int) ([]TransactionDetails, error)
}

type repositoryTransactionDetails struct {
	db *gorm.DB
}

func NewRepositoryTransactionDetails(db *gorm.DB) *repositoryTransactionDetails {
	return &repositoryTransactionDetails{db}
}

func (r *repositoryTransactionDetails) FindAll() ([]TransactionDetails, error) {
	var transaction []TransactionDetails

	err := r.db.Preload("Product").Preload("User").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repositoryTransactionDetails) Save(transaction TransactionDetails) (TransactionDetails, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransactionDetails) FindByUserId(productID int, userID int) ([]TransactionDetails, error) {
	var transaction []TransactionDetails

	// err := r.db.Joins("User", r.db.Where(&user.User{ID: userID})).Joins("Products", r.db.Where(&product.Product{ID: ProductID})).Find(&comment).Error

	err := r.db.Joins("Product", r.db.Where(&product.Products{ID: productID})).Joins("User", r.db.Where(&user.User{ID: userID})).Find(&transaction).Error

	// err := r.db.Preload("User").Preload("Product").Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransactionDetails) FindById(ID int) (TransactionDetails, error) {
	var transaction TransactionDetails

	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransactionDetails) Update(transaction TransactionDetails) (TransactionDetails, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (r *repositoryTransactionDetails) Delete(transaction TransactionDetails) (TransactionDetails, error) {
	err := r.db.Delete(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
