package entity

import (
	"time"
)

type ProfileEntity struct {
	Id        string    `gorm:"primaryKey"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Age       int64     `gorm:"not null"`
	Job       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ProfileEntity) TableName() string {
	return "profile"
}
