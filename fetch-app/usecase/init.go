package usecase

import (
	"MyAPI/adapters"
	"MyAPI/adapters/database"
	"MyAPI/adapters/resty"
)

// UseCase is wrapper for usecase interface
type UseCase interface {
	UserUseCase
	FishUseCase
}

// DatabaseUseCase data struct of database config
type DatabaseUseCase struct {
	*database.Database
}

// RestyUseCase data struct of resty client
type RestyUseCase struct {
	*resty.Resty
}

// usecase data struct of usecase
type usecase struct {
	DatabaseUseCase
	RestyUseCase
}

// InitUseCase initialize usecase
func InitUseCase(adapters adapters.Adapters) UseCase {
	return usecase{
		DatabaseUseCase{adapters.Database},
		RestyUseCase{adapters.Resty},
	}
}
