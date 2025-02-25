package configuration

import (
	"log"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func AddConnectionDb(connectionString string) (*gorm.DB, error) {
	var err error
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	dbConnection = database
	if err != nil {
		log.Fatalf(message.MessageFailConnectionDb, err)
	}

	log.Println(message.MessageConnectionDb)

	return dbConnection, err
}
