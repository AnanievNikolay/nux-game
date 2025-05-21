package dto

import (
	"github.com/AnanievNikolay/nux-game/domain"
)

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Phone    string `db:"phone"`
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Phone:    u.Phone,
	}
}

func NewUser(domainObject *domain.User) User {
	return User{
		ID:       domainObject.ID,
		Username: domainObject.Username,
		Phone:    domainObject.Phone,
	}
}
