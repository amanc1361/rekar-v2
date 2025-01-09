package usecase

import (
	"errors"

	"rekar.ir/v2/auth/internal/domain/models"
	"rekar.ir/v2/auth/internal/infastructure/repository"
)

// UserUseCase منطق کسب‌وکار مربوط به کاربران را مدیریت می‌کند
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase ایجاد یک نمونه جدید از UserUseCase
func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

// Authenticate اعتبارسنجی کاربر را انجام می‌دهد
func (uc *UserUseCase) Authenticate(username, password string) (*models.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// در یک پیاده‌سازی واقعی باید از hashing برای رمز عبور استفاده شود
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// Register ثبت‌نام کاربر جدید را انجام می‌دهد
func (uc *UserUseCase) Register(user *models.User) error {
	return uc.userRepo.Create(user)
}
