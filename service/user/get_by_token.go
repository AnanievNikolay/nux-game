package user

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) GetUserByToken(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) (*domain.User, error) {
	logger = logger.WithFields(s.logger.Data).WithField("token", token)

	mf := utils.LogTimeSpent(logger, "GetUserByToken")
	defer mf()

	userToken, err := s.tokenService.GetToken(ctx, logger, token)
	if err != nil {
		return nil, fmt.Errorf("tokenService.GetToken: %w", err)
	}

	if userToken == nil || userToken.ExpiresAt.Before(time.Now()) {
		return nil, domain.ErrTokenInvalidOrExpired
	}

	user, err := s.repository.GetByToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("repository.GetByToken: %w", err)
	}

	if user == nil {
		return nil, domain.ErrorUserNotFound
	}

	return user, nil
}
