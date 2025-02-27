package main

import (
	"log"
	"os"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/middlewares"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/routes"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/constants"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/containers"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var router *gin.Engine = gin.Default()

	err := godotenv.Load(constants.PathFileEnv)
	if err != nil {
		log.Println(message.ErrorLoadingEnv, err)
	}

	containers.RegisterMediators()
	router.Use(middlewares.LogErrorRequest())
	router.Use(middlewares.LogSuccessRequest())
	routes.InitRoutes(router)

	if err := router.Run(":" + os.Getenv(constants.KeyPort)); err != nil {
		log.Println(message.ErrorRunningServer, err)
	} else {
		log.Println(message.StartRunningServer)
	}

}
