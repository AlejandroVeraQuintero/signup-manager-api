package routes

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/constants"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/api/controllers"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	router.GET(constants.ApiVersionPath+"/profiles", controllers.GetAllProfiles)
	router.POST(constants.ApiVersionPath+"/profile", controllers.AddProfile)
	router.GET(constants.ApiVersionPath+"/profile/:id", controllers.GetByIdProfile)
	router.DELETE(constants.ApiVersionPath+"/profile/:id", controllers.DeleteProfile)
}
