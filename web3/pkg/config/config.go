package config

// Only import standard libraries into config files

import (
	"log"
)

type AppConfig struct {
	Infolog *log.Logger
}