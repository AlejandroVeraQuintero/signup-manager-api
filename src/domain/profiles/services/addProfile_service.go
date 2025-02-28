package services

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
)

type AddProfileService struct {
	ProfileRepository ports.IProfileRepository
}

func (service *AddProfileService) Add(profile models.Profile) (models.Profile, error) {
	id, err := service.ProfileRepository.Add(profile)
	if err != nil {
		return models.Profile{}, err
	}
	profile.Id = id
	return profile, nil
}
