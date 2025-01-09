package models

import "gorm.io/gorm"

type UserFinancialAccess struct {
	gorm.Model
	UserID   uint `gorm:"not null"`
	PeriodID uint `gorm:"not null"`
}
