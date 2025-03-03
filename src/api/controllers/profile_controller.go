package controllers

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/addProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/deleteProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/queries/getByIdProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/queries/getListProfiles"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

// GetAllProfiles godoc
// @Summary      Retrieve all profiles
// @Description  Returns a complete list of all profiles with their information
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Success      200  {array}  dtos.ProfileDto
// @Failure      400  {object} utils.ErrorResponse "Bad request error"
// @Failure      500  {object} utils.ErrorResponse "Internal server error"
// @Security     BearerAuth
// @Router       /profiles [get]
func GetAllProfiles(context *gin.Context) {
	var query *getListProfiles.GetListProfilesQuery = &getListProfiles.GetListProfilesQuery{}
	response, err := mediatr.Send[*getListProfiles.GetListProfilesQuery, []dtos.ProfileDto](context, query)
	utils.HandleResponse(context, response, err)
}

// GetByIdProfile godoc
// @Summary      Retrieve a profile by ID
// @Description  Returns a single profile that matches the specified ID
// @Param        id path string true "Profile ID" format(uuid)
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Success      200  {object}  dtos.ProfileDto
// @Failure      400  {object}  utils.ErrorResponse "Bad request error"
// @Failure      404  {object}  utils.ErrorResponse "Profile not found"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Security     BearerAuth
// @Router       /profile/{id} [get]
func GetByIdProfile(context *gin.Context) {
	var query *getByIdProfile.GetByIdProfileQuery = &getByIdProfile.GetByIdProfileQuery{Id: context.Param("id")}
	response, err := mediatr.Send[*getByIdProfile.GetByIdProfileQuery, dtos.ProfileDto](context, query)
	utils.HandleResponse(context, response, err)
}

// AddProfile godoc
// @Summary      Create a new profile
// @Description  Creates and saves a new profile to the database
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        profile body models.ProfileRequest true "Profile information"
// @Success      201  {object}  models.Profile
// @Failure      400  {object}  utils.ErrorResponse "Bad request error"
// @Failure      409  {object}  utils.ErrorResponse "Profile already exists"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Security     BearerAuth
// @Router       /profile [post]
func AddProfile(context *gin.Context) {
	var command *addProfile.AddProfileCommand
	context.BindJSON(&command)
	response, err := mediatr.Send[*addProfile.AddProfileCommand, dtos.AddProfileDto](context, command)
	utils.HandleResponse(context, response, err)
}

// DeleteProfile godoc
// @Summary      Delete a profile
// @Description  Permanently removes a profile from the database by its ID
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        id path string true "Profile ID" format(uuid)
// @Success      204  {object}  nil "Profile successfully deleted"
// @Failure      400  {object}  utils.ErrorResponse "Bad request error"
// @Failure      404  {object}  utils.ErrorResponse "Profile not found"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Security     BearerAuth
// @Router       /profile/{id} [delete]
func DeleteProfile(context *gin.Context) {
	var command *deleteProfile.DeleteProfileCommand = &deleteProfile.DeleteProfileCommand{Id: context.Param("id")}
	response, err := mediatr.Send[*deleteProfile.DeleteProfileCommand, dtos.DeleteProfileDto](context, command)
	utils.HandleResponse(context, response, err)
}
