package repository

import (
	"mymodule/entity"

	"gorm.io/gorm"
)


type IUserRepository interface {
	
	Create(user entity.User) error
	// yang didalam kurung itu nilai kembalian
	GetAll() ([]entity.User,error)
}

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err!=nil{
		return err
	}
	return nil
}

func (u UserRepository) GetAll () ([]entity.User, error) {
	// bikin variabe buat nyimpen data yang diambil dari entity. User
	var users []entity.User
	if err := u.db.Find(&users).Error; err!=nil {
		return nil,nil
	}
	return users, nil
}