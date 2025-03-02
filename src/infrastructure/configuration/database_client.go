package configuration

import (
	"log"
	"os"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/migrations"
	constants "github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/resources"
	message "github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/resources"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDatabaseInstance() *gorm.DB {
	var err error
	var connectionString string = os.Getenv(constants.KeyConnectionString)

	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf(message.MessageFailConnectionDb, err)
	}

	migrations.ApplyMigrations(dbConnection)

	return dbConnection
}
