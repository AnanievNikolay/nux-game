package game

import (
	"context"

	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(ctx context.Context, game *domain.Game) (int, error)
	GetHistoryByToken(ctx context.Context, token string) ([]domain.Game, error)
}

type TokenService interface {
	GetValidToken(
		ctx context.Context,
		logger *logrus.Entry,
		token string,
	) (*domain.Token, error)
}

type Service struct {
	logger *logrus.Entry

	maxGameNumber int

	repository Repository

	tokenService TokenService
}

func NewService(
	logger *logrus.Entry,

	cfg *config.Config,

	repository Repository,

	tokenService TokenService,
) *Service {
	return &Service{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "service",
			"service": "game",
		}),

		maxGameNumber: cfg.Game.MaxGameNumber,

		repository: repository,

		tokenService: tokenService,
	}
}
