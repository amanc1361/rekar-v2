package repository

import (
	"errors"

	"gorm.io/gorm"
	"rekar.ir/v2/auth/internal/domain/models"
)

// UserRepository رابط برای عملیات مربوط به کاربران
type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}

// GormUserRepository پیاده‌سازی GORM از UserRepository
type GormUserRepository struct {
	db *gorm.DB
}

// NewUserRepository ایجاد یک نمونه جدید از GormUserRepository
func NewUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (repo *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (repo *GormUserRepository) Create(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return errors.New("could not create user")
	}
	return nil
}
