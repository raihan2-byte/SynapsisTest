package transactiondetail

type ServiceTransactionDetails interface {
	GetTransactions() ([]TransactionDetails, error)
}

type serviceTransactionDetails struct {
	repositoryTransaction RepositoryTransactionDetails
}

func NewService(repositoryTransaction RepositoryTransactionDetails) *serviceTransactionDetails {
	return &serviceTransactionDetails{repositoryTransaction}
}

func (s *serviceTransactionDetails) GetTransactions() ([]TransactionDetails, error) {
	transaction, err := s.repositoryTransaction.FindAll()
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}