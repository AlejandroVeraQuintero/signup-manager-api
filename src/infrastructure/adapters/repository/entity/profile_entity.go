package entity

type ProfileEntity struct {
	BaseEntity
	FirstName string `gorm:"not null" validate:"required,min=2,max=100"`
	LastName  string `gorm:"not null" validate:"required,min=2,max=100"`
	Email     string `gorm:"not null" validate:"required,email"`
	Age       int64  `gorm:"not null"`
	Job       string `gorm:"not null" validate:"required,min=2,max=200"`
	State     string `gorm:"not null" validate:"required,min=2,max=80"`
}

func (ProfileEntity) TableName() string {
	return "profile"
}
