package addProfile

import (
	"context"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/services"
)

type AddProfileHandler struct {
	addProfileService *services.AddProfileService
}

func NewAddProfileHandler(addProfileService *services.AddProfileService) *AddProfileHandler {
	return &AddProfileHandler{addProfileService: addProfileService}
}

func (handler *AddProfileHandler) Handle(context context.Context, command *AddProfileCommand) (dtos.AddProfileDto, error) {
	var newProfile models.Profile = models.Profile{
		Email:     command.Email,
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Age:       command.Age,
		Job:       command.Job,
	}

	profile, err := handler.addProfileService.Add(newProfile)

	if err != nil {
		return dtos.AddProfileDto{}, err
	}

	response := dtos.AddProfileDto{Id: profile.Id, Email: profile.Email}

	return response, err
}
