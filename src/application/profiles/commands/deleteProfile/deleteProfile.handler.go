package deleteProfile

import (
	"context"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
)

type DeleteProfileHandler struct {
	profileRepository ports.IProfileRepository
}

func NewDeleteProfileHandler(profileRepository ports.IProfileRepository) *DeleteProfileHandler {
	return &DeleteProfileHandler{profileRepository: profileRepository}
}

func (handler *DeleteProfileHandler) Handle(context context.Context, command *DeleteProfileCommand) (*dtos.DeleteProfileDto, error) {
	var err error = handler.profileRepository.Delete(command.Id)
	if err != nil {
		return nil, err
	}

	response := &dtos.DeleteProfileDto{Id: command.Id}
	return response, err
}
