package addprofile

type AddProfileCommand struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Age       int64  `json:"age" validate:"required"`
	Job       string `json:"job" validate:"required"`
}
