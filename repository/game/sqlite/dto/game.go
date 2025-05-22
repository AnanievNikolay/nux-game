package dto

import (
	"time"

	"github.com/AnanievNikolay/nux-game/domain"
)

type Game struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	IsWin     bool      `db:"is_win"`
	Number    int       `db:"number"`
	Prize     float32   `db:"prize"`
	CreatedAt time.Time `db:"created_at"`
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
		UserID:    domainObject.UserID,
		IsWin:     domainObject.IsWin,
		Number:    domainObject.Number,
		Prize:     domainObject.Prize,
		CreatedAt: domainObject.CreatedAt,
	}
}

type Games []Game

func (g Games) ToDomain() []domain.Game {
	retVal := make([]domain.Game, len(g))

	for i := range g {
		retVal[i] = *g[i].ToDomain()
	}

	return retVal
}
