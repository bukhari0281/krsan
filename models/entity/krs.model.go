package entity

import (
	"gorm.io/gorm"
	"time"
)

type Krs struct {
	ID             uint           `json:"id,omitempty" gorm:"primary_key"`
	MahasiswaID    uint           `json:"mahasiswa_id" gorm:"not null"`
	Mahasiswa      Mahasiswa      `json:"mahasiswa"`
	JadwalKuliahID uint           `json:"jadwal_kuliah_id"`
	JadwalKuliah   JadwalKuliah   `json:"jadwal_kuliah"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}
