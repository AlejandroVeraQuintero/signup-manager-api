package containers

import "github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/services"

func GetAddProfileService() *services.AddProfileService {
	return &services.AddProfileService{
		ProfileRepository: GetProfileRepository(),
	}
}

func GetDeleteProfileService() *services.DeleteProfileService {
	return &services.DeleteProfileService{
		ProfileRepository: GetProfileRepository(),
	}
}
