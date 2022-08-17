package controller

import "MyAPI/entity"

// UserController is wrapper for user controller
type UserController interface {
	GetUserById(id string) (user entity.Users, err error)
}

// GetUserById get user by ID
func (controller Controller) GetUserById(id string) (user entity.Users, err error) {
	user, err = controller.usecase.GetUserById(id)
	return
}
