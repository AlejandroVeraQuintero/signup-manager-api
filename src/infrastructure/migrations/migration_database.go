package migrations

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository/entity"
	"gorm.io/gorm"
)

func ApplyMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.ProfileEntity{})
}
