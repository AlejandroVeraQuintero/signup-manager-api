package entity

type ProfileEntity struct {
	BaseEntity
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Age       int64  `gorm:"not null"`
	Job       string `gorm:"not null"`
}

func (ProfileEntity) TableName() string {
	return "profile"
}
