packge models
type UserBusinessAccess struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"not null"`
	BusinessID uint `gorm:"not null"`
	PeriodID   uint `gorm:"not null"`
}