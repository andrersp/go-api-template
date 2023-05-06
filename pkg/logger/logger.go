package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	log  zerolog.Logger
	once sync.Once
)

func GetLogger() zerolog.Logger {
	once.Do(func() {
		log = zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	})
	return log
}
