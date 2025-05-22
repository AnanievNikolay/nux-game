package dto

type UserCreateRequest struct {
	Username string `validate:"required,min=3,max=32"`
	Phone    string `validate:"required,e164"`
}
