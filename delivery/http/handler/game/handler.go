package game

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Service interface {
	Play(
		ctx context.Context,
		logger *logrus.Entry,
		token string,
	) (*domain.Game, error)
	GetHistoryByToken(
		ctx context.Context,
		logger *logrus.Entry,
		token string,
	) ([]domain.Game, error)
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
