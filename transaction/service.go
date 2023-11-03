package transaction

import (
	"errors"
	"synapsisid/product"
	transactiondetail "synapsisid/transactionDetail"
	"synapsisid/user"
)

type ServiceTransaction interface {
	CreateTransaction(getIDProduct GetIDProduct, GetUserID int, input TransactionInput) (Transaction, error)
	GetTransaction(productID int, userID int) ([]Transaction, error)
	SaveToTransaction(userID int, transactionID int)(transactiondetail.TransactionDetails, error)

}

type serviceTransaction struct {
	repository        RepositoryTransaction
	repositoryProduct product.RepositoryProduct
	repositoryUser    user.RepositoryUser
	repositoryTransactionDetails transactiondetail.RepositoryTransactionDetails
}

func NewService(repository RepositoryTransaction, repositoryProduct product.RepositoryProduct, repositoryUser user.RepositoryUser, repositoryTransactionDetails transactiondetail.RepositoryTransactionDetails) *serviceTransaction {
	return &serviceTransaction{repository, repositoryProduct, repositoryUser, repositoryTransactionDetails}
}

func (s *serviceTransaction) SaveToTransaction(userID int, transactionID int)(transactiondetail.TransactionDetails, error) {
	transactionDetails := transactiondetail.TransactionDetails{}

	getUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return transactiondetail.TransactionDetails{}, nil
	}

	// transactionID := Transaction{}
	
	getDataTransaction, err := s.repository.FindById(transactionID)
	if err != nil {
		return transactiondetail.TransactionDetails{}, errors.New("item not found")
	}

	transactionDetails.Quantity = getDataTransaction.Quantity 
	transactionDetails.UserID = getUser.ID 
	transactionDetails.ProductID = getDataTransaction.ProductID 
	transactionDetails.TotalPrice = getDataTransaction.TotalPrice

	save, err := s.repositoryTransactionDetails.Save(transactionDetails)
	if err != nil {
		return save, err
	}
	return save, nil

}

func (s *serviceTransaction) CreateTransaction(getIDProduct GetIDProduct, GetUserID int, input TransactionInput) (Transaction, error) {
	transaction := Transaction{}

	cek, err := s.repositoryProduct.FindById(getIDProduct.ID)
	if err != nil {
		print("erorrrrrr")
		return Transaction{}, errors.New("item not found")
	}
	
	if cek.Stock < input.Quantity {
		return Transaction{}, errors.New("quantity not enough")
	}

	cekSaldo, err := s.repositoryUser.FindById(GetUserID)
	if err != nil {
		return Transaction{}, err
	}
	if cekSaldo.Balance < (input.Quantity * cek.Price) {
		return Transaction{}, errors.New("balance not enough")
	}

	totalBelanja, err := s.repositoryProduct.FindById(getIDProduct.ID)
	if err != nil {
		return Transaction{}, err
	}
	cek.Price = input.Quantity * totalBelanja.Price

	cekSaldo.Balance = cekSaldo.Balance - (input.Quantity * cek.Price)

	_, err = s.repositoryUser.Update(cekSaldo)
	if err != nil {
		return Transaction{}, err
	}

	cek.Stock = cek.Stock - input.Quantity


	transaction.Quantity = input.Quantity
	transaction.TotalPrice = cek.Price
	transaction.ProductID = cek.ID
	transaction.UserID = cekSaldo.ID

	newProduct, err := s.repository.Save(transaction)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil

	
}

// func (s *serviceTransaction) GetTransactionUser(ID int) ([]Transaction, error){

// }

func (s *serviceTransaction) GetTransaction(productID int, userID int) ([]Transaction, error) {

	transaction, err := s.repository.FindByUserId(productID, userID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil

	
}
