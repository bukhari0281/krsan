package entity

import (
	"gorm.io/gorm"
	"time"
)

type JadwalKuliah struct {
	ID           uint           `json:"id,omitempty" gorm:"primary_key"`
	MataKuliahID uint           `json:"mata_kuliah_id"`
	MataKuliah   MataKuliah     `json:"mata_kuliah"`
	DosenID      uint           `json:"dosen_id"`
	Dosen        Dosen          `json:"dosen"`
	Hari         string         `json:"hari" gorm:"not null"`
	Jam          time.Time      `json:"jam" gorm:"not null"`
	Ruangan      string         `json:"ruangan" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
