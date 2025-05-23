package user

import (
	"context"
	"errors"
	"fmt"

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

	userToken, err := s.tokenService.GetValidToken(ctx, logger, token)
	if err != nil {
		if errors.Is(err, domain.ErrTokenInvalidOrExpired) {
			return nil, err
		}
		return nil, fmt.Errorf("tokenService.GetToken: %w", err)
	}

	user, err := s.repository.GetByID(ctx, userToken.UserID)
	if err != nil {
		return nil, fmt.Errorf("repository.GetByID: %w", err)
	}

	if user == nil {
		return nil, domain.ErrorUserNotFound
	}

	user.Token = userToken.Token

	return user, nil
}
