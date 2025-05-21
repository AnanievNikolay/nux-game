package token

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) GetToken(
	ctx context.Context,
	logger *logrus.Entry,
	token string,
) (*domain.Token, error) {
	logger = logger.WithFields(s.logger.Data).WithFields(logrus.Fields{
		"token": token,
	})

	mf := utils.LogTimeSpent(logger, "GetToken")
	defer mf()

	userToken, err := s.repository.GetToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("repository.GetToken: %w", err)
	}

	return userToken, nil
}
