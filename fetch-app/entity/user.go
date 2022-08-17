package entity

import (
	"time"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id          string    `xorm:"id" json:"id"`
	Name        string    `xorm:"name" json:"name"`
	PhoneNumber string    `xorm:"phone_number" json:"phone_number"`
	Password    string    `xorm:"password" json:"password"`
	Role        string    `xorm:"role" json:"role"`
	CreatedAt   time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time `xorm:"updated_at" json:"updated_at"`
	DeletedAt   time.Time `xorm:"deleted_at" json:"deleted_at"`
}

type UsersRequest struct {
	Id          string `xorm:"id" json:"id"`
	Name        string `xorm:"name" json:"name"`
	PhoneNumber string `xorm:"phone_number" json:"phone_number"`
	Password    string `xorm:"password" json:"password"`
	Role        string `xorm:"role" json:"role"`
}

// NewUser create new user
func NewUser(request UsersRequest) (user Users) {
	user = Users{
		Id:          uuid.NewV4().String(),
		Name:        request.Name,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Role:        request.Role,
	}

	return
}

// HashPassword encrypts user password
func (user *Users) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *Users) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
