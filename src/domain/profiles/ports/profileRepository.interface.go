package ports

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/domain/profiles/models"
)

type IProfileRepository interface {
	FindAll() ([]models.Profile, error)
	FindById(id string) (models.Profile, error)
	Add(profile models.Profile) (string, error)
	Delete(id string) error
}
