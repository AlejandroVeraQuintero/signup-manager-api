package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers 		godoc
// @Summary			Get All profiles.
// @Description		Return list of profiles.
// @Tags			profiles
// @Router			/profiles [get]
func GetAllProfiles(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "GetAllProfiles"})
}

// GetByIdUser 			godoc
// @Summary				Get Single profile by id.
// @Param				id path string true "Profile ID"
// @Description			Return the tahs whoes id valu mathes id.
// @Produce				application/json
// @Tags				profiles
// @Router				/profile/{id} [get]
func GetByIdProfile(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "GetByIdProfile"})
}

// AddUser			godoc
// @Summary			Create Profile
// @Description		Save profile db
// @Produce			application/json
// @Tags			profiles
// @Router			/profile [post]
func AddProfile(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "AddProfile"})
}

// DeleteUser		godoc
// @Summary			Delete Profile
// @Description		Remove profile data by id.
// @Produce			application/json
// @Tags			profiles
// @Router			/profile/{id} [delete]
func DeleteProfile(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "DeleteProfile"})
}
