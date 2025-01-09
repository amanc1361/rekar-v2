package models

type FinancialPeriod struct {
	ID          uint   `gorm:"primaryKey"`
	BusinessID  uint   `gorm:"not null"`
	Name        string `gorm:"size:50;not null"`
	Start       string `gorm:"type:date;not null"`
	End         string `gorm:"type:date;not null"`
	Description string `gorm:"size:200;not null"`
	IsClosed    bool   `gorm:"not null"`
}
