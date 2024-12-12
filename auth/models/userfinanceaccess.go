packge models
type UserFinancialAccess struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"not null"`
	PeriodID uint `gorm:"not null"`
}