package game

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) GetHistoryByToken(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) ([]domain.Game, error) {
	logger = logger.WithFields(s.logger.Data).WithField("token", token)

	mf := utils.LogTimeSpent(logger, "GetHistoryByToken")
	defer mf()

	userToken, err := s.tokenService.GetValidToken(ctx, logger, token)
	if err != nil {
		if errors.Is(err, domain.ErrTokenInvalidOrExpired) {
			return nil, err
		}
	}

	history, err := s.repository.GetHistoryByToken(ctx, userToken.Token)
	if err != nil {
		return nil, fmt.Errorf("repository.GetHistoryByToken: %w", err)
	}

	return history, nil
}
