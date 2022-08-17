package entity

import (
	"time"
)

// Users entity of user
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
