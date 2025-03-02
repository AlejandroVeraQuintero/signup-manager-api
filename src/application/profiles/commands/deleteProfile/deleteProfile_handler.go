package deleteProfile

import (
	"context"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/services"
)

type DeleteProfileHandler struct {
	deleteProfileService *services.DeleteProfileService
}

func NewDeleteProfileHandler(deleteProfileService *services.DeleteProfileService) *DeleteProfileHandler {
	return &DeleteProfileHandler{deleteProfileService: deleteProfileService}
}

func (handler *DeleteProfileHandler) Handle(context context.Context, command *DeleteProfileCommand) (dtos.DeleteProfileDto, error) {
	var err error = handler.deleteProfileService.Delete(command.Id)
	if err != nil {
		return dtos.DeleteProfileDto{}, err
	}

	response := dtos.DeleteProfileDto{Id: command.Id}
	return response, err
}
