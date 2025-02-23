package main

import (
	"fmt"
	"os"

	constants "github.com/AlejandroVeraQuintero/signup-manager-api/src/api/constants"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/routes"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var router *gin.Engine = gin.Default()

	err := godotenv.Load(constants.PathFileEnv)
	if err != nil {
		fmt.Println(message.ErrorLoadingEnv, err)
	}

	routes.InitRoutes(router)

	if err := router.Run(":" + os.Getenv(constants.KeyPort)); err != nil {
		fmt.Println(message.ErrorRunningServer, err)
	} else {
		fmt.Println(message.StartRunningServer)
	}

}
