package getListProfiles

import (
	"context"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
)

type GetListProfilesHandler struct {
	profileRepository ports.IProfileRepository
}

func NewGetListProfilesHandler(profileRepository ports.IProfileRepository) *GetListProfilesHandler {
	return &GetListProfilesHandler{profileRepository: profileRepository}
}

func (handler *GetListProfilesHandler) Handle(context context.Context, query *GetListProfilesQuery) (*[]models.Profile, error) {
	var err error
	var profiles []models.Profile

	profiles, err = handler.profileRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return &profiles, err
}
