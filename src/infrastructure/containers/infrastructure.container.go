package containers

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/ports"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/adapters/repository"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/configuration"
)

func GetProfileRepository() ports.IProfileRepository {
	return &repository.ProfileRepository{
		Connection: configuration.GetDatabaseInstance(),
	}
}
