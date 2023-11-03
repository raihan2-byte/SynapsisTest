package category

import (
	"gorm.io/gorm"
)

type RepositoryCategory interface {
	FindAll() ([]Categorys, error)
	Save(category Categorys) (Categorys, error)
	FindById(ID int) (Categorys, error)
	Update(category Categorys) (Categorys, error)
	Delete(category Categorys) (Categorys, error)
}

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db}
}

func (r *repositoryCategory) FindAll() ([]Categorys, error) {
	var category []Categorys

	err := r.db.Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}


func (r *repositoryCategory) Save(category Categorys) (Categorys, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

// func (r *repositoryCategory) FindByEmail(email string) (User, error) {
// 	var user User
// 	err := r.db.Where("email = ?", email).Find(&user).Error

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (r *repositoryCategory) FindById(ID int) (Categorys, error) {
	var category Categorys

	err := r.db.Where("id = ?", ID).First(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repositoryCategory) Update(category Categorys) (Categorys, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil

}

func (r *repositoryCategory) Delete(category Categorys) (Categorys, error) {
	err := r.db.Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
