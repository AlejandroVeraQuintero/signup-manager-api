package getByIdProfile

import (
	"context"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/dtos"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
)

type GetByIdProfileHandler struct {
	profileRepository ports.IProfileRepository
}

func NewGetByIdProfileHandler(profileRepository ports.IProfileRepository) *GetByIdProfileHandler {
	return &GetByIdProfileHandler{profileRepository: profileRepository}
}

func (handler *GetByIdProfileHandler) Handle(context context.Context, query *GetByIdProfileQuery) (dtos.ProfileDto, error) {
	var err error
	var profile models.Profile
	profile, err = handler.profileRepository.FindById(query.Id)

	if err != nil {
		return dtos.ProfileDto{}, err
	}

	return dtos.ProfileModelToProfileDto(profile), err

}
