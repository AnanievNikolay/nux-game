package token

import (
	"context"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetToken(ctx context.Context, token string) (*domain.Token, error)
	Deactivate(ctx context.Context, token string) error
}

type UnitOfWork interface {
	UpdateToken(
		ctx context.Context,
		oldToken string,
		newToken *domain.Token,
	) error
}

type Service struct {
	logger *logrus.Entry

	repository Repository
	uow        UnitOfWork

	ttl int
}

func NewService(
	logger *logrus.Entry,
	cfg *config.Config,
	repository Repository,
	uow UnitOfWork,
) *Service {
	return &Service{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "service",
			"service": "token",
		}),

		repository: repository,
		uow:        uow,

		ttl: cfg.Service.Token.TTL,
	}
}
