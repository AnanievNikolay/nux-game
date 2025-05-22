package domain

import "time"

type Game struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	IsWin     bool      `json:"is_win"`
	Number    int       `json:"number"`
	Prize     float32   `json:"prize"`
	CreatedAt time.Time `json:"created_at"`
}
