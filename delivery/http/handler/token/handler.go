package token

import (
	"context"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

type Service interface {
	UpdateToken(ctx context.Context, logger *logrus.Entry, token string) (*domain.Token, error)
	DeactivateToken(ctx context.Context, logger *logrus.Entry, token string) error
}

type Handler struct {
	logger *logrus.Entry

	service Service
}

func NewHandler(logger *logrus.Entry, service Service) *Handler {
	return &Handler{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "handler",
			"handler": "token",
		}),

		service: service,
	}
}
