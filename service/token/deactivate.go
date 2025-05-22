package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
	"github.com/sirupsen/logrus"
)

func (s *Service) DeactivateToken(ctx context.Context, logger *logrus.Entry, token string) error {
	logger = logger.WithFields(s.logger.Data).WithField("token", token)

	mf := utils.LogTimeSpent(logger, "DeactivateToken")
	defer mf()

	userToken, err := s.GetValidToken(ctx, logger, token)
	if err != nil {
		if errors.Is(err, domain.ErrTokenInvalidOrExpired) {
			return err
		}

		return fmt.Errorf("service.GetValidToken: %w", err)
	}

	if err := s.repository.Deactivate(ctx, userToken.Token); err != nil {
		return fmt.Errorf("repository.Deactivate: %w", err)
	}

	return nil
}
