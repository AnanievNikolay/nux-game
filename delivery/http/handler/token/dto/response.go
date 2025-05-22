package dto

import (
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Token struct {
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewToken(domainObject *domain.Token) *Token {
	if domainObject == nil {
		return nil
	}

	return &Token{
		UserID:    domainObject.UserID,
		Token:     domainObject.Token,
		ExpiresAt: domainObject.ExpiresAt,
	}
}
