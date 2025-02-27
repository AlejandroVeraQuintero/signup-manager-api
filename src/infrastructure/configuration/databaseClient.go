package configuration

import (
	"log"
	"os"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository/entity"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/constants"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDatabaseInstance() *gorm.DB {
	var err error
	var connectionString string = os.Getenv(constants.KeyConnectionString)
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	dbConnection = database
	if err != nil {
		log.Fatalf(message.MessageFailConnectionDb, err)
	}
	migrateDatabase(dbConnection)
	log.Println(message.MessageConnectionDb)

	return dbConnection
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&entity.ProfileEntity{})
}
