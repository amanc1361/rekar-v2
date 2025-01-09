package models

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:20;not null"`
	Description string `gorm:"size:200;not null"`
	BusinessID  uint   `gorm:"not null"`
}
