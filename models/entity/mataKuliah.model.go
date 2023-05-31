package entity

import (
	"gorm.io/gorm"
	"time"
)

type MataKuliah struct {
	ID             uint           `json:"id,omitempty" gorm:"primary_key"`
	KodeMataKuliah string         `json:"kode_mata_kuliah" gorm:"not null"`
	NamaMataKuliah string         `json:"nama_mata_kuliah" gorm:"not null"`
	SKS            int            `json:"sks" gorm:"not null"`
	Nama           string         `json:"nama" gorm:"not null"`
	Semester       int            `json:"semester" gorm:"not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}
