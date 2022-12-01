package usecase

import (
	"mymodule/entity"
	"mymodule/entity/response"
	"mymodule/repository"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
}

type UserUseCase struct {
	UserRepository repository.IUserRepository
}

func NewUserService(UserRepository repository.IUserRepository) *UserUseCase{
	return &UserUseCase{UserRepository:UserRepository}
}

func (u UserUseCase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	err := u.UserRepository.Create(user)
	if err!=nil{
		return err
	}
	return nil
}

func (u UserUseCase) GetList() ([]response.GetUserResponse, error){
	users, err := u.UserRepository.GetAll()

	if err != nil {
		return nil,nil
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy( &userResponse,&users)
	return userResponse,nil

}