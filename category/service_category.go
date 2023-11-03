package category

type ServiceCategory interface {
	CreateCategory(input CategoryInput) (Categorys, error)
	// Login(input LoginInput) (User, error)
	GetCategory(ID int) (Categorys, error)
	GetCategorys() ([]Categorys, error)
	DeleteCategory(ID int) (Categorys, error)
	UpdatedCategory(getUpdatedInput GetinputCategoryID, inputUser UpdatedCategory) (Categorys, error)
}

type serviceCategory struct {
	repositoryCategory RepositoryCategory
}

func NewServiceCategory(repositoryCategory RepositoryCategory) *serviceCategory {
	return &serviceCategory{repositoryCategory}
}

func (s *serviceCategory) CreateCategory(input CategoryInput) (Categorys, error) {
	category := Categorys{}

	category.Type = input.Type

	newCategory, err := s.repositoryCategory.Save(category)
	if err != nil {
		return newCategory, err
	}
	return newCategory, nil
}

func (s *serviceCategory) GetCategorys() ([]Categorys, error) {

	category, err := s.repositoryCategory.FindAll()
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *serviceCategory) GetCategory(ID int) (Categorys, error) {

	category, err := s.repositoryCategory.FindById(ID)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *serviceCategory) UpdatedCategory(getUpdatedInput GetinputCategoryID, inputCategory UpdatedCategory) (Categorys, error) {

	category, err := s.repositoryCategory.FindById(getUpdatedInput.ID)
	if err != nil {
		return category, err
	}

	// if category.ID != inputCategory.User.ID {
	// 	return category, errors.New("not an owner the account")
	// }

	category.Type = inputCategory.Type

	categoryUpdated, err := s.repositoryCategory.Update(category)
	if err != nil {
		return categoryUpdated, err
	}

	return categoryUpdated, nil

}

func (s *serviceCategory) DeleteCategory(categoryID int) (Categorys, error) {
	category, err := s.repositoryCategory.FindById(categoryID)
	if err != nil {
		return category, err
	}
	categoryDel, err := s.repositoryCategory.Delete(category)

	if err != nil {
		return categoryDel, err
	}
	return categoryDel, nil
}

// func (s *serviceCategory) GetUserByid(ID int) (User, error) {
// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	if user.ID == 0 {
// 		return user, errors.New("User Not Found With That ID")
// 	}

// 	return user, nil

// }
