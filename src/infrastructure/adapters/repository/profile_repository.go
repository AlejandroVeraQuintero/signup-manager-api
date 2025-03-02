package repository

import (
	"errors"
	"fmt"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository/entity"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/mapper"
	"gorm.io/gorm"
)

var _ ports.IProfileRepository = &ProfileRepository{}

type ProfileRepository struct {
	Connection *gorm.DB
}

func (repository *ProfileRepository) FindAll() ([]models.Profile, error) {
	var profiles []models.Profile
	var profilesEntities []entity.ProfileEntity

	if err := repository.Connection.Find(&profilesEntities).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("no profiles found"))
	}

	if len(profilesEntities) == 0 {
		return nil, errors.New(fmt.Sprintf("no profiles found"))
	}

	profiles = mapper.ProfilesEntitiesToProfiles(profilesEntities)
	return profiles, nil
}

func (repository *ProfileRepository) FindById(id string) (models.Profile, error) {

	var profileEntity entity.ProfileEntity
	if err := repository.Connection.Where("id = ?", id).First(&profileEntity).Error; err != nil {
		return models.Profile{}, errors.New(fmt.Sprintf("error finding profile: %s", err))
	}
	var profile models.Profile = mapper.ProfileEntityToProfile(profileEntity)
	return profile, nil
}

func (repository *ProfileRepository) Add(profile models.Profile) (string, error) {
	var profileEntity entity.ProfileEntity = mapper.ProfileToProfileEntity(profile)
	if err := repository.Connection.Create(&profileEntity).Error; err != nil {
		return "", errors.New(fmt.Sprintf("error creating profile: %s", err))
	}
	return profileEntity.Id, nil
}

func (repository *ProfileRepository) Delete(id string) error {
	var err error
	if err = repository.Connection.Where("id = ?", id).Delete(&entity.ProfileEntity{}).Error; err != nil {
		return errors.New(fmt.Sprintf("error deleting profile: %s", err))
	}
	return nil
}
