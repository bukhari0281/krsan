package user

import (
	"gorm.io/gorm"
	"time"
)

type UserLogin struct {
	ID         uint           `json:"id,omitempty" gorm:"primary_key"`
	Username   string         `json:"username,omitempty" gorm:"not_null"`
	Password   string         `json:"password,omitempty" gorm:"not_null" validate:"required,min=8"`
	Created_At time.Time      `json:"created_at"`
	Updated_At time.Time      `json:"updated_at"`
	Deleted_At gorm.DeletedAt `json:"-"`
}
