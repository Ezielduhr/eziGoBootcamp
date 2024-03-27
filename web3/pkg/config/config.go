package config

import "log"

// Only import standard libraries into config files.
// ^ lol? *confused noises*

import (
	scs "github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	Infolog   *log.Logger
	Session   *scs.SessionManager
	CSRFToken string
}
