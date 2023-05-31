package user

import (
	"gorm.io/gorm"
	"time"
)

type UserRegistration struct {
	ID         uint           `json:"id,omitempty" gorm:"primary_key"`
	Name       string         `json:"name" gorm:"not_null"`
	Username   string         `json:"username,omitempty" gorm:"not_null"`
	Email      string         `json:"email" validate:"required,email"`
	Password   string         `json:"password,omitempty" gorm:"not_null" validate:"required,min=8"`
	Created_At time.Time      `json:"created_at"`
	Updated_At time.Time      `json:"updated_at"`
	Deleted_At gorm.DeletedAt `json:"-"`
}
