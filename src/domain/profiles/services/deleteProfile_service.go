package services

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
)

type DeleteProfileService struct {
	ProfileRepository ports.IProfileRepository
}

func (service *DeleteProfileService) Delete(id string) error {
	var err error
	_, err = service.ProfileRepository.FindById(id)

	if err != nil {
		return err
	}

	err = service.ProfileRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
