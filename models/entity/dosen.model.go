package entity

import (
	"gorm.io/gorm"
	"time"
)

type Dosen struct {
	ID             uint           `json:"id,omitempty" gorm:"primary_key"`
	BidangKeahlian string         `json:"bidang_keahlian" gorm:"not null"`
	NamaDosen      string         `json:"nama_dosen" gorm:"not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}
