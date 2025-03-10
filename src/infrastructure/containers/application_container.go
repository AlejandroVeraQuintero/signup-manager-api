package containers

import (
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/addProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/commands/deleteProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/queries/getByIdProfile"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/application/profiles/queries/getListProfiles"
	"github.com/mehdihadeli/go-mediatr"
)

func RegisterMediators() {
	mediatr.RegisterRequestHandler(addProfile.NewAddProfileHandler(GetAddProfileService()))
	mediatr.RegisterRequestHandler(deleteProfile.NewDeleteProfileHandler(GetDeleteProfileService()))
	mediatr.RegisterRequestHandler(getByIdProfile.NewGetByIdProfileHandler(GetProfileRepository()))
	mediatr.RegisterRequestHandler(getListProfiles.NewGetListProfilesHandler(GetProfileRepository()))
}
