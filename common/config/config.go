package config

import (
	"github.com/natefinch/lumberjack"
)

type Config struct {
	Logger   *lumberjack.Logger `json:"logger"`
	Delivery *Delivery          `json:"delivery"`
	DB       *DB                `json:"db"`
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
}
