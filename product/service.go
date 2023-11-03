package product

import "synapsisid/category"

type ServiceProduct interface {
	CreateProduct(input ProductInput) (Products, error)
	GetProducts() ([]Products, error)
	GetProduct(ID int) (Products, error)
	DeleteProduct(ID int) (Products, error)
	GetProductByCategory (ID int) ([]Products, error)
	UpdatedProduct(getUpdatedInput GetinputProductID, inputProduct UpdatedProduct) (Products, error)
}

type serviceProduct struct {
	repositoryProduct  RepositoryProduct
	categoryRepository category.RepositoryCategory
}

func NewServiceProduct(repositoryProduct RepositoryProduct, categoryRepository category.RepositoryCategory) *serviceProduct {
	return &serviceProduct{repositoryProduct, categoryRepository}
}

func (s *serviceProduct) GetProductByCategory (ID int) ([]Products, error) {
	product, err := s.repositoryProduct.FindAllProductByCategory(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) GetProducts() ([]Products, error) {

	product, err := s.repositoryProduct.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) CreateProduct(input ProductInput) (Products, error) {
	product := Products{}

	product.Title = input.Title
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	// cek := input.CategoryID
	cekCategory, err := s.categoryRepository.FindById(input.CategoryID)
	if input.CategoryID != cekCategory.ID {
		return Products{}, err
	}

	newProduct, err := s.repositoryProduct.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) GetProduct(ID int) (Products, error) {

	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) UpdatedProduct(getUpdatedInput GetinputProductID, inputProduct UpdatedProduct) (Products, error) {

	// cek := inputProduct.CategoryID
	// cekCategory, err := s.repository.FindById(cek)
	// if cekCategory.ID != cekCategory.Category.ID {
	// 	return cekCategory, err
	// }

	_, err := s.categoryRepository.FindById(inputProduct.CategoryID)
	if err != nil {
		return Products{}, err
	}
	// if newProduct.CategoryID != inputProduct.Category.ID {
	// 	return cekCategory, errors.New("not an owner the account")
	// }

	product, err := s.repositoryProduct.FindById(getUpdatedInput.ID)
	if err != nil {
		return product, err
	}

	// if product.ID != inputProduct.User.ID {
	// 	return product, errors.New("not an owner the account")
	// }

	product.Title = inputProduct.Title
	product.Price = inputProduct.Price
	product.Stock = inputProduct.Stock
	product.CategoryID = inputProduct.CategoryID

	productUpdated, err := s.repositoryProduct.Update(product)
	if err != nil {
		return productUpdated, err
	}

	return productUpdated, nil

}

func (s *serviceProduct) DeleteProduct(ID int) (Products, error) {

	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return product, err
	}
	productDel, err := s.repositoryProduct.Delete(product)

	if err != nil {
		return productDel, err
	}
	return productDel, nil


// func (s *serviceProduct) GetUserByid(ID int) (Products, error) {
// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	if user.ID == 0 {
// 		return user, errors.New("User Not Found With That ID")
// 	}

// 	return user, nil

// }
}