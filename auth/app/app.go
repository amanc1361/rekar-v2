package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rekar.ir/v2/auth/config"
	"rekar.ir/v2/auth/internal/delivery/http"
	"rekar.ir/v2/auth/internal/infastructure/repository"
	"rekar.ir/v2/auth/internal/usecase"
)

type App struct {
	DB          *gorm.DB
	UserRepo    repository.UserRepository
	UserUseCase *usecase.UserUseCase
	Router      *gin.Engine
}

// Initialize تنظیمات اولیه برنامه شامل اتصال به پایگاه‌داده و ایجاد مخازن
func (a *App) Initialize() {
	// اتصال به پایگاه‌داده
	a.DB = config.ConnectDatabase()

	// ایجاد مخزن کاربران (Repository)
	a.UserRepo = repository.NewUserRepository(a.DB)

	// ایجاد UseCase برای کاربران
	a.UserUseCase = usecase.NewUserUseCase(a.UserRepo)

	// ایجاد Router جدید
	a.Router = gin.Default()

	// ثبت مسیرهای HTTP
	http.RegisterRoutes(a.Router, a.UserUseCase)
}

// Run اجرای سرور HTTP
func (a *App) Run(addr string) {
	log.Printf("Server is running on port %s", addr)
	if err := a.Router.Run(addr); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
