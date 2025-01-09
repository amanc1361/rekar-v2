package repository

import (
	"gorm.io/gorm"
)

// Repository مدیریت پایگاه‌داده و ایجاد مخازن
type Repository struct {
	DB *gorm.DB
}

// NewRepository ایجاد یک نمونه جدید از Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
