package handler

import "tidy/service"

type userHandle struct {
	userService service.UserService
}

func NewUserHandle(userService service.UserService) userHandle {
	return userHandle{userService: userService}
}
