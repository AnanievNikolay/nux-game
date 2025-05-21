package user

import (
	"context"

	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

type UnitOfWork interface {
	Save(
		ctx context.Context,
		user *domain.User,
		token *domain.Token,
	) error
}

type TokenService interface {
	Issue(ctx context.Context, logger *logrus.Entry, userID string) *domain.Token
	GetToken(
		ctx context.Context,
		logger *logrus.Entry,
		token string,
	) (*domain.Token, error)
}

type Repository interface {
	GetIDByUsernameAndPhone(
		ctx context.Context,
		username, phone string,
	) (string, error)
	GetByToken(
		ctx context.Context,
		token string,
	) (*domain.User, error)
}

type Service struct {
	logger *logrus.Entry

	unitOfWork UnitOfWork

	tokenService TokenService

	repository Repository
}

func NewService(
	logger *logrus.Entry,

	unitOfWork UnitOfWork,

	tokenService TokenService,

	repository Repository,
) *Service {
	return &Service{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "service",
			"service": "user",
		}),

		unitOfWork: unitOfWork,

		tokenService: tokenService,

		repository: repository,
	}
}
