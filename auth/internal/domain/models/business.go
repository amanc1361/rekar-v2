packge models
type Business struct {
	ID          uint   `gorm:"primaryKey"`
	OwnerID     uint   `gorm:"not null"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255;not null"`
}