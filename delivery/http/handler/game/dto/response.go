package dto

import (
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Game struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	IsWin     bool      `json:"is_win"`
	Number    int       `json:"number"`
	Prize     float32   `json:"prize"`
	CreatedAt time.Time `json:"created_at"`
}

func (g *Game) ToDomain() *domain.Game {
	return &domain.Game{
		ID:        g.ID,
		UserID:    g.UserID,
		IsWin:     g.IsWin,
		Number:    g.Number,
		Prize:     g.Prize,
		CreatedAt: g.CreatedAt,
	}
}

func NewGame(domainObject *domain.Game) Game {
	return Game{
		ID:        domainObject.ID,
		UserID:    domainObject.UserID,
		IsWin:     domainObject.IsWin,
		Number:    domainObject.Number,
		Prize:     domainObject.Prize,
		CreatedAt: domainObject.CreatedAt,
	}
}

type Games []Game

func NewGames(domainObjects []domain.Game) Games {
	retVal := make(Games, len(domainObjects))

	for i := range domainObjects {
		retVal[i] = NewGame(&domainObjects[i])
	}

	return retVal
}
