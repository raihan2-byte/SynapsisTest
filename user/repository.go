package user

import (
	"gorm.io/gorm"
)

type RepositoryUser interface {
	//create User
	Save(user User) (User, error)
	FindById(ID int) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
	UpdateBalance(userID int, updateMoney int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateBalance(userID int, updateMoney int) error {
	var user User

	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindById(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) Delete(user User) (User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
