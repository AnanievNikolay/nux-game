package config

import (
	"github.com/natefinch/lumberjack"
)

type Config struct {
	Logger   *lumberjack.Logger `json:"logger"`
	Delivery *Delivery          `json:"delivery"`
	DB       *DB                `json:"db"`
	Service  *Service           `json:"service"`
	Game     *Game              `json:"game"`
}

type Delivery struct {
	HTTP *HTTP `json:"http"`
}

type HTTP struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DB struct {
	SQLite *SQLite `json:"sqlite"`
}

type SQLite struct {
	MigratePath string `json:"migrate_path"`
	FileName    string `json:"file_name"`
	FileFolder  string `json:"file_folder"`
	TTL         int    `json:"ttl"`
}

type Service struct {
	Token TokenService `json:"token"`
}

type TokenService struct {
	TTL int `json:"ttl"`
}

type Game struct {
	MaxGameNumber int `json:"max_game_number"`
}
