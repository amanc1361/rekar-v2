package usecase

// UserUseCaseInterface تعریف رابط برای UserUseCase

import (
	"errors"

	"rekar.ir/v2/auth/internal/domain/models"
	"rekar.ir/v2/auth/internal/infastructure/repository"
)

type UserUseCaseInterface interface {
	Authenticate(username, password string) (*models.User, error)
	Register(user *models.User) error
	Delete(id int) error
	Update(user *models.User) error
	Read(id int) (*models.User, error)
}

// UserUseCaseInterface تعریف رابط برای UserUseCase

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

// خواندن کاربر را انجام می‌دهد  Read
func (uc *UserUseCase) Read(id uint) (*models.User, error) {
	user, err := uc.userRepo.Read(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc *UserUseCase) Update(user *models.User) error {
	return uc.userRepo.Update(user)
}
func (uc *UserUseCase) Delete(id uint) error {
	return uc.userRepo.Delete(id)
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
