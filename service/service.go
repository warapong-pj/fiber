package service

import "tidy/domain"

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserService interface {
	Reads()
	Read()
	Create()
	Update()
	Delete()
}

type userService struct {
	userDomain domain.UserDomain
}

func NewUserService(userDomain domain.UserDomain) userService {
	return userService{userDomain: userDomain}
}
