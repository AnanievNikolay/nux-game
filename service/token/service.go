package token

import (
	"github.com/AnanievNikolay/nux-game/common/config"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Entry

	ttl int
}

func NewService(logger *logrus.Entry, cfg *config.Config) *Service {
	return &Service{
		logger: logger.WithFields(logrus.Fields{
			"layer":   "service",
			"service": "token",
		}),

		ttl: cfg.Service.Token.TTL,
	}
}
