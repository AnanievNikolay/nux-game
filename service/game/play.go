package game

import (
	"context"
	"errors"
	"fmt"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

func (s *Service) Play(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) (*domain.Game, error) {
	logger = logger.WithFields(s.logger.Data).WithField("token", token)

	mf := utils.LogTimeSpent(logger, "Play")
	defer mf()

	userToken, err := s.tokenService.GetValidToken(ctx, logger, token)
	if err != nil {
		if errors.Is(err, domain.ErrTokenInvalidOrExpired) {
			return nil, err
		}
	}

	gameResult := s.getGameResult(userToken.UserID, s.getNumber())

	id, err := s.repository.Save(ctx, gameResult)
	if err != nil {
		return nil, fmt.Errorf("repository.Save: %w", err)
	}

	gameResult.ID = id

	return gameResult, nil
}
