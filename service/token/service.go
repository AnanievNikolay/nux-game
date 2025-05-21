package token

import (
	"context"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetToken(ctx context.Context, token string) (*domain.Token, error)
}

type Service struct {
	logger *logrus.Entry

	repository Repository

	ttl int
}

func NewService(logger *logrus.Entry, cfg *config.Config, repository Repository) *Service {
	return &Service{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "service",
			"service": "token",
		}),

		repository: repository,

		ttl: cfg.Service.Token.TTL,
	}
}
