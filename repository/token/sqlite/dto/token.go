package dto

import (
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Token struct {
	UserID    string    `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}

func NewToken(domainObject *domain.Token) Token {
	return Token{
		UserID:    domainObject.UserID,
		Token:     domainObject.Token,
		ExpiresAt: domainObject.ExpiresAt,
	}
}
