packge models

type User struct {
    gorm.Model
	Name     string `gorm:"size:50;not null"`
	Family   string `gorm:"size:50;not null"`
	Mobile   string `gorm:"size:15;not null"`
	Email    string `gorm:"size:50;not null"`
	Username string `gorm:"size:20;not null"`
	Password string `gorm:"size:20;not null"`
	RoleID   uint   `gorm:"not null"`
}