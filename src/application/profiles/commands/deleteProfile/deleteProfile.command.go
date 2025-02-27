package deleteProfile

type DeleteProfileCommand struct {
	Id string `validate:"required"`
}
