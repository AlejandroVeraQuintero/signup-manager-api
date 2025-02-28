package routes

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/controllers"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/constants"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	router.GET(constants.ApiVersionOnePath+"/profiles", controllers.GetAllProfiles)
	router.POST(constants.ApiVersionOnePath+"/profile", controllers.AddProfile)
	router.GET(constants.ApiVersionOnePath+"/profile/:id", controllers.GetByIdProfile)
	router.DELETE(constants.ApiVersionOnePath+"/profile/:id", controllers.DeleteProfile)
}
