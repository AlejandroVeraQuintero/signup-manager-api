package dtos

import "github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"

type ProfileDto struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int64  `json:"age"`
	Job       string `json:"job"`
	State     string `json:"state"`
}

func ProfileModelToProfileDto(profile models.Profile) ProfileDto {
	var profileDto ProfileDto = ProfileDto{
		Id:        profile.Id,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Age:       profile.Age,
		Job:       profile.Job,
		State:     profile.State,
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
