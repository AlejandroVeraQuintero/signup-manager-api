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
	if err := repository.Connection.First(&profileEntity, id).Error; err != nil {
		return models.Profile{}, fmt.Errorf("error finding profile: %w", err)
	}
	var profile models.Profile = mapper.ProfileEntityToProfile(profileEntity)
	return profile, nil
}

func (repository *ProfileRepository) Add(profile models.Profile) (string, error) {
	var profileEntity entity.ProfileEntity = mapper.ProfileToProfileEntity(profile)
	if err := repository.Connection.Create(&profileEntity).Error; err != nil {
		return "", fmt.Errorf("error creating profile: %w", err)
	}
	return profileEntity.Id, nil
}

func (repository *ProfileRepository) Delete(id string) error {
	var err error
	_, err = repository.FindById(id)

	if err != nil {
		return err
	}

	if err := repository.Connection.Delete(&entity.ProfileEntity{}, id).Error; err != nil {
		return fmt.Errorf("error deleting profile: %w", err)
	}

	return nil
}
