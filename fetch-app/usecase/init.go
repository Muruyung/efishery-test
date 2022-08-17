package usecase

import (
	"MyAPI/adapters"
	"MyAPI/adapters/database"
	"MyAPI/adapters/resty"
)

type UseCase interface {
	UserUseCase
	FishUseCase
}

type DatabaseUseCase struct {
	*database.Database
}

type RestyUseCase struct {
	*resty.Resty
}

type usecase struct {
	DatabaseUseCase
	RestyUseCase
}

func InitUseCase(adapters adapters.Adapters) UseCase {
	return usecase{
		DatabaseUseCase{adapters.Database},
		RestyUseCase{adapters.Resty},
	}
}
