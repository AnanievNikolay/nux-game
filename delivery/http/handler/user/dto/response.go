package dto

import "github.com/AnanievNikolay/nux-game/domain"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}

func NewUser(domainObject *domain.User) *User {
	if domainObject == nil {
		return nil
	}

	return &User{
		ID:       domainObject.ID,
		Username: domainObject.Username,
		Token:    domainObject.Token,
		Phone:    domainObject.Phone,
	}
}
