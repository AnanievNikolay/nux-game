package user

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Service interface {
	CreateUser(
		ctx context.Context,
		logger *logrus.Entry,
		username, phone string,
	) (*domain.User, error)
	GetUserByToken(
		ctx context.Context,
		logger *logrus.Entry,
		token string,
	) (*domain.User, error)
}

type Handler struct {
	logger *logrus.Entry

	service Service
}

func NewHandler(logger *logrus.Entry, service Service) *Handler {
	return &Handler{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "handler",
			"handler": "user",
		}),

		service: service,
	}
}
