package entity

import (
	"gorm.io/gorm"
	"time"
)

type Mahasiswa struct {
	ID            uint           `json:"id,omitempty" gorm:"primary_key"`
	Nama          string         `json:"nama" gorm:"not null"`
	NIM           int            `json:"nim" gorm:"not null"`
	Program_Studi string         `json:"program_sudi" gorm:"not null"`
	Semester      int            `json:"semester" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}
