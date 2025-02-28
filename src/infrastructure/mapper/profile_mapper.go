package mapper

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository/entity"
	"github.com/google/uuid"
)

func ProfileToProfileEntity(profile models.Profile) entity.ProfileEntity {
	var profileEntity entity.ProfileEntity
	profileEntity.Id = uuid.New().String()
	profileEntity.FirstName = profile.FirstName
	profileEntity.LastName = profile.LastName
	profileEntity.Email = profile.Email
	profileEntity.Age = profile.Age
	profileEntity.Job = profile.Job

	return profileEntity
}

func ProfileEntityToProfile(profileEntity entity.ProfileEntity) models.Profile {
	var profile models.Profile
	profile.Id = profileEntity.Id
	profile.FirstName = profileEntity.FirstName
	profile.LastName = profileEntity.LastName
	profile.Email = profileEntity.Email
	profile.Age = profileEntity.Age
	profile.Job = profileEntity.Job
	return profile
}

func ProfilesEntitiesToProfiles(profilesEntities []entity.ProfileEntity) []models.Profile {
	var profiles []models.Profile
	for _, profileEntity := range profilesEntities {
		profile := ProfileEntityToProfile(profileEntity)
		profiles = append(profiles, profile)
	}
	return profiles
}
