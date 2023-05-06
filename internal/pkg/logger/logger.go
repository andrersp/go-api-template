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

func Get() zerolog.Logger {
	once.Do(func() {
		log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	})
	return log
}
