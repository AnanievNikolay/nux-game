package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) UpdateToken(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) (*domain.Token, error) {
	logger = logger.WithFields(s.logger.Data).WithField("token", token)

	mf := utils.LogTimeSpent(logger, "UpdateToken")
	defer mf()

	userToken, err := s.GetValidToken(ctx, logger, token)
	if err != nil {
		if errors.Is(err, domain.ErrTokenInvalidOrExpired) {
			return nil, err
		}

		return nil, fmt.Errorf("service.GetValidToken: %w", err)
	}

	logger = logger.WithField("userID", userToken.UserID)

	newToken := s.Issue(ctx, logger, userToken.UserID)

	if err := s.uow.UpdateToken(ctx, userToken.Token, newToken); err != nil {
		return nil, fmt.Errorf("uow.UpdateToken: %w", err)
	}

	return newToken, nil
}
