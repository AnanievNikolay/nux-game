package token

import (
	"context"
	"fmt"
	"time"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

func (s *Service) GetValidToken(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) (*domain.Token, error) {
	logger = logger.WithFields(s.logger.Data).WithFields(logrus.Fields{
		"token": token,
	})

	mf := utils.LogTimeSpent(logger, "GetValidToken")
	defer mf()

	userToken, err := s.repository.GetToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("repository.GetToken: %w", err)
	}

	if userToken == nil || userToken.ExpiresAt.Before(time.Now()) {
		return nil, domain.ErrTokenInvalidOrExpired
	}

	return userToken, nil
}
