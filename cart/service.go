package cart

import (
	"errors"
	"fmt"
	"synapsisid/category"
	"synapsisid/product"
	"synapsisid/transaction"
	transactiondetail "synapsisid/transactionDetail"
	"synapsisid/user"
)

type ServiceCarts interface {
	AddToCart(getIdProduct GetIdProduct, userID int, input InputCart) (Carts, error)
	GetAllCartByUserId(userID int) ([]Carts, error)
	DeleteCart(userID int, ID GetIDCart) (Carts, error)
	PayByUserID(userID int) (transactiondetail.TransactionDetails, error)
	UpdatedCart(getIdCart int, inputProduct InputCart, userID int) (Carts, error)
	PayFromCartID(getIDCart GetIDCart, userID int) (transactiondetail.TransactionDetails, error)
}

type serviceCarts struct {
	repositoryCarts  RepositoryCart
	categoryRepository category.RepositoryCategory
	repositoryProduct product.RepositoryProduct
	repositoryUser user.RepositoryUser
	repositoryTransactionDetails transactiondetail.RepositoryTransactionDetails
	repositoryTransaction transaction.RepositoryTransaction
}

func NewServiceCart(repositoryCarts  RepositoryCart, categoryRepository category.RepositoryCategory, repositoryProduct product.RepositoryProduct, repositoryUser user.RepositoryUser, repositoryTransactionDetails transactiondetail.RepositoryTransactionDetails, repositoryTransaction transaction.RepositoryTransaction) *serviceCarts {
	return &serviceCarts{repositoryCarts, categoryRepository, repositoryProduct, repositoryUser, repositoryTransactionDetails, repositoryTransaction}
} 


func (s *serviceCarts) DeleteCart(userID int, ID GetIDCart) (Carts, error) {

	cart, err := s.repositoryCarts.FindById(ID.ID)
	if err != nil {
		return cart, err
	}

	findID, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return Carts{}, errors.New("user id not found")
	}
	if cart.UserID != findID.ID {
		return Carts{}, errors.New("not an owner userid")
	}

	newCart, err := s.repositoryCarts.Delete(cart)
	if err != nil {
		return newCart, err
	}

	return newCart, nil
}

func (s *serviceCarts) PayFromCartID(getIDCart GetIDCart, userID int) (transactiondetail.TransactionDetails, error) {
	var transactionDetails transactiondetail.TransactionDetails

	cek, err := s.repositoryCarts.FindById(getIDCart.ID)
	if err != nil {
		return transactionDetails, errors.New("item not found")
	}

	productItem, err := s.repositoryProduct.FindById(cek.ProductID)
	if err != nil {
		return transactionDetails, errors.New("product not found")
	}
	// Validasi stok produk

	if productItem.Stock < cek.Quantity {
		return transactionDetails, errors.New("stock is not enough")
	}

	cekSaldo, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return transactionDetails, errors.New("user not found")
	}
	if cekSaldo.Balance < (cek.Quantity * productItem.Price) {
		return transactionDetails, errors.New("balance not enough")
	}

	totalBelanja, err := s.repositoryProduct.FindById(cek.ProductID)
	if err != nil {
		return transactionDetails, errors.New("product not found")
	}

	cek.TotalPrice = cek.Quantity * totalBelanja.Price

	cekSaldo.Balance = cekSaldo.Balance - (cek.Quantity * productItem.Price)

	_, err = s.repositoryUser.Update(cekSaldo)
	if err != nil {
		return transactionDetails, err
	}

	productItem.Stock = productItem.Stock - cek.Quantity


	transactionDetails.Quantity = cek.Quantity
	transactionDetails.TotalPrice = cek.TotalPrice
	transactionDetails.ProductID = productItem.ID
	transactionDetails.UserID = cekSaldo.ID

	newProduct, err := s.repositoryTransactionDetails.Save(transactionDetails)
	if err != nil {
		print(err)
		return newProduct, err
	}

	return newProduct, nil

}

func (s *serviceCarts) PayByUserID(userID int) (transactiondetail.TransactionDetails, error) {
	var transactionDetails transactiondetail.TransactionDetails

	userData, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return transactionDetails, errors.New("user not found")
	}

	// Temukan semua cart milik pengguna berdasarkan userID
	carts, err := s.repositoryCarts.FindAllByUserID(userID)
	if err != nil {
		return transactionDetails, err
	}

	if len(carts) == 0 {
		return transactionDetails, errors.New("products in the cart not found")
	}

	var totalAmount int
	var productsInTransaction []transaction.Transaction

	for _, cart := range carts {
		if cart.Quantity <= 0 {
			return transactionDetails, errors.New("invalid quantity")
		}

		productData, err := s.repositoryProduct.FindById(cart.ProductID)
		if err != nil {
			return transactionDetails, err
		}

		if productData.Stock < cart.Quantity {
			return transactionDetails, errors.New("stock not enough in the cart")
		}

		cart.TotalPrice = cart.Quantity * productData.Price

		productData.Stock = productData.Stock -  cart.Quantity
		_, err = s.repositoryProduct.Update(productData)
		if err != nil {
			return transactionDetails, fmt.Errorf("update stock product failed : %v", err)
		}

		totalAmount = totalAmount + cart.TotalPrice

		// Tambahkan produk ke dalam transaksi
		transaction := transaction.Transaction{
			ProductID:  cart.ProductID,
			UserID:     userID,
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		}
		productsInTransaction = append(productsInTransaction, transaction)
	}

	if userData.Balance < totalAmount {
		return transactionDetails, errors.New("balance not enough")
	}

	userData.Balance = userData.Balance - totalAmount
	_, err = s.repositoryUser.Update(userData)
	if err != nil {
		return transactionDetails, err
	}

	// Simpan semua produk dalam satu transaksi
	for _, productTransaction := range productsInTransaction {
		_, err = s.repositoryTransaction.Save(productTransaction)
		if err != nil {
			return transactionDetails, err
		}
	}

	transactionDetails.Quantity = len(carts)
	transactionDetails.UserID = userID
	transactionDetails.TotalPrice = totalAmount
	transactionDetails.ProductID = carts[0].ProductID 

	saveTransaction, err := s.repositoryTransactionDetails.Save(transactionDetails)
	if err != nil {
		return saveTransaction, err
	}

	return saveTransaction, nil
}

func (s *serviceCarts) GetAllCartByUserId(userID int) ([]Carts, error) {

	cart, err := s.repositoryCarts.FindAllByUserID(userID)
	if err != nil {
		return cart, errors.New("user not found")
	}
	return cart, nil
}

func (s *serviceCarts) UpdatedCart(cartID int, input InputCart, userID int) (Carts, error) {
    cart := Carts{}

    existingCart, err := s.repositoryCarts.FindById(cartID)
    if err != nil {
        return existingCart, err
    }

    IdUser, err := s.repositoryUser.FindById(userID)
    if err != nil || IdUser.ID != userID {
        return Carts{}, err
    }
	if IdUser.ID != existingCart.UserID{
		return Carts{}, errors.New("not the owner")
	}

    // Ambil data produk dari basis data menggunakan existingCart.ProductID
    productData, err := s.repositoryProduct.FindById(existingCart.ProductID)
    if err != nil {
        return Carts{}, err
    }

    existingCart.Quantity = input.Quantity
    existingCart.TotalPrice = productData.Price * input.Quantity

    cart.ProductID = existingCart.ProductID
    cart.CategoryID = existingCart.CategoryID
    cart.TotalPrice = existingCart.TotalPrice
    cart.UserID = IdUser.ID

    // Lakukan pembaruan ke basis data
    updatedCart, err := s.repositoryCarts.Update(existingCart)
    if err != nil {
        return Carts{}, err
    }

    return updatedCart, nil
}



func (s *serviceCarts) 	AddToCart(getIdProduct GetIdProduct, userID int, input InputCart) (Carts, error){
	cart := Carts{}

	productId, err := s.repositoryProduct.FindById(getIdProduct.ID)
	if err != nil {
		return Carts{}, err
	}

	totalPrice, err := s.repositoryProduct.FindById(getIdProduct.ID)
	if err != nil {
		return Carts{}, err
	}

	productId.Price = input.Quantity * totalPrice.Price

	

	IdUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return Carts{}, err
	}


	cart.Quantity = input.Quantity
	cart.TotalPrice = productId.Price
	cart.ProductID = productId.ID
	cart.CategoryID = productId.Category.ID
	cart.UserID = IdUser.ID 

	
	existingCart, err := s.repositoryCarts.FindByProductAndUser(getIdProduct.ID, userID)
		if err != nil {
			return Carts{}, err
		}

		if existingCart.ID != 0 {
			existingCart.Quantity += input.Quantity
			tes, _ := s.repositoryProduct.FindById(getIdProduct.ID)

			existingCart.TotalPrice = tes.Price * existingCart.Quantity

			updatedCart, err := s.repositoryCarts.Update(existingCart)
			if err != nil {
				return Carts{}, err
			}
			return updatedCart, nil
		}

	newCart, err := s.repositoryCarts.Save(cart)
	if err != nil {
		return newCart, err
	}

	
	return newCart, nil
}

