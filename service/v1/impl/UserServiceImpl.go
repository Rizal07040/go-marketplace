package service

import (
	"go-marketplace/models/repositories"
	"go-marketplace/response/v1"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}

}

func (t UserServiceImpl) GetAllUsers() (users []response.UserResponse, err error) {
	result, err := t.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		task := response.UserResponse{
			Id:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		}
		users = append(users, task)
	}
	return users, nil
}
