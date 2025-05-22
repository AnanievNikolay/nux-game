package game

import (
	"math"
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
)

func (s *Service) getGameResult(userID string, number int) *domain.Game {
	return &domain.Game{
		UserID:    userID,
		IsWin:     number%2 == 0,
		Number:    number,
		Prize:     s.calculatePrize(number),
		CreatedAt: time.Now(),
	}
}

func (s *Service) calculatePrize(number int) float32 {
	var result float64

	switch {
	case number > 900:
		result = float64(number) * 0.7
	case number > 600:
		result = float64(number) * 0.5
	case number > 300:
		result = float64(number) * 0.3
	default:
		result = float64(number) * 0.1
	}

	return float32(math.Round(result*100) / 100)
}
