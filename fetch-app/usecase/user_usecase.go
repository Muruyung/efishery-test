package usecase

import (
	"MyAPI/entity"

	_ "github.com/lib/pq"
)

type UserUseCase interface {
	GetUserById(id string) (user entity.Users, err error)
}

// GetUserById query for select data by ID
func (db DatabaseUseCase) GetUserById(id string) (user entity.Users, err error) {
	_, err = db.SQL("SELECT * FROM users WHERE id=?", id).Get(&user)
	// _, err = db.Where(fmt.Sprintf("id=%s", id)).Get(&user)
	return
}
