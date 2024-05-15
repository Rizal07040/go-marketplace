package service

import "go-marketplace/response/v1"

type UserService interface {
	GetAllUsers() (users []response.UserResponse, err error)
}
