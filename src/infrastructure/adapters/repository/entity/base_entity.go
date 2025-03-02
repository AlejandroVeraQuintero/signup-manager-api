package entity

import "time"

type BaseEntity struct {
	Id          string    `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	StateEntity string    `gorm:"default:'active'"`
}
