package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rekar.ir/v2/auth/internal/domain/models"
)

// LoadEnvVariables بارگذاری متغیرهای محیطی از فایل .env

// ConnectDatabase ایجاد اتصال به پایگاه داده و انجام مهاجرت مدل‌ها
func ConnectDatabase() *gorm.DB {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// مهاجرت مدل‌ها
	err = db.AutoMigrate(&models.User{}, &models.Business{}, &models.FinancialPeriod{}, &models.Role{}, &models.UserBusinessAccess{}, &models.UserFinancialAccess{})
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	return db
}
