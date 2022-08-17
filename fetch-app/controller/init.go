package controller

import (
	"MyAPI/adapters"
	"MyAPI/usecase"
)

// Service instance controller singleton
var Service controller

// controller is wrapper for interface controller
type controller interface {
	FishController
	UserController
}

// Controller is wrapper for interface usecase
type Controller struct {
	usecase usecase.UseCase
}

// InitController initialize controller
func InitController(adapters adapters.Adapters) {
	usecase := usecase.InitUseCase(adapters)
	Service = Controller{
		usecase: usecase,
	}
}
