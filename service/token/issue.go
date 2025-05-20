package token

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) Issue(ctx context.Context, logger *logrus.Entry, userID string) *domain.Token {
	logger = logger.WithFields(s.logger.Data)

	mf := utils.LogTimeSpent(logger, "Issue")
	defer mf()

	return &domain.Token{
		UserID:    userID,
		Token:     uuid.NewString(),
		ExpiresAt: time.Now().Add(time.Duration(s.ttl) * time.Second),
	}
}
