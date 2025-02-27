package controllers

import (
	"net/http"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/addProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/deleteProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/queries/getByIdProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
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
	var command *getByIdProfile.GetByIdProfileQuery = &getByIdProfile.GetByIdProfileQuery{Id: context.Param("id")}
	response, err := mediatr.Send[*getByIdProfile.GetByIdProfileQuery, *models.Profile](context, command)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	context.IndentedJSON(http.StatusOK, response)
}

// AddUser			godoc
// @Summary			Create Profile
// @Description		Save profile db
// @Produce			application/json
// @Tags			profiles
// @Router			/profile [post]
func AddProfile(context *gin.Context) {
	var command *addProfile.AddProfileCommand
	context.BindJSON(&command)
	response, err := mediatr.Send[*addProfile.AddProfileCommand, *dtos.AddProfileDto](context, command)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	context.IndentedJSON(http.StatusOK, response)
}

// DeleteUser		godoc
// @Summary			Delete Profile
// @Description		Remove profile data by id.
// @Produce			application/json
// @Tags			profiles
// @Router			/profile/{id} [delete]
func DeleteProfile(context *gin.Context) {
	var command *deleteProfile.DeleteProfileCommand = &deleteProfile.DeleteProfileCommand{Id: context.Param("id")}
	response, err := mediatr.Send[*deleteProfile.DeleteProfileCommand, *dtos.DeleteProfileDto](context, command)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	context.IndentedJSON(http.StatusOK, response)
}
