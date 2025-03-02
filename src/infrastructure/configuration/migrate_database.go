package configuration

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository/entity"
	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&entity.ProfileEntity{})
}
