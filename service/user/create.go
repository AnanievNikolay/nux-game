package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/AnanievNikolay/nux-game/common/utils"
	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) CreateUser(
	ctx context.Context,
	logger *logrus.Entry,
	username, phone string,
) (*domain.User, error) {
	logger = logger.WithFields(s.logger.Data).WithFields(logrus.Fields{
		"username": username,
		"phone":    phone,
	})

	mf := utils.LogTimeSpent(logger, "CreateUser")
	defer mf()

	id, err := s.repository.GetIDByUsernameAndPhone(ctx, username, phone)
	if err != nil {
		return nil, fmt.Errorf("repository.GetIDByUsernameAndPhone: %w", err)
	}

	if id != "" {
		return nil, domain.ErrorUsernameWithThisPhoneNotUnique
	}

	user := &domain.User{
		ID:       uuid.NewString(),
		Username: username,
		Phone:    phone,
	}

	tokenLogger := logrus.NewEntry(logger.Logger)

	token := s.tokenService.Issue(ctx, tokenLogger, user.ID)

	if err := s.unitOfWork.Save(ctx, user, token); err != nil {
		return nil, fmt.Errorf("unitOfWork.Save: %w", err)
	}

	return user, nil
}
