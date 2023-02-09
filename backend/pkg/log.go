package pkg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// SetupZeroLog sets up the zerolog logger
func SetupZeroLog() {
	// setup zerolog format
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})
}
