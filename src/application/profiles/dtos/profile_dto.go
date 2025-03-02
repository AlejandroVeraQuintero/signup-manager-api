package dtos

import "github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"

type ProfileDto struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Age       int64
	Job       string
}

func ProfileModelToProfileDto(profile models.Profile) ProfileDto {
	var profileDto ProfileDto = ProfileDto{
		Id:        profile.Id,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Age:       profile.Age,
		Job:       profile.Job,
	}
	return profileDto
}

func ProfilesModelToProfilesDto(profilesModel []models.Profile) []ProfileDto {
	var profiles []ProfileDto
	for _, profileModel := range profilesModel {
		profile := ProfileModelToProfileDto(profileModel)
		profiles = append(profiles, profile)
	}
	return profiles
}
