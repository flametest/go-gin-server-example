package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var logger zerolog.Logger

func Initialize(app string, logLevel zerolog.Level) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger = zerolog.
		New(output).
		Level(logLevel).
		With().Timestamp().Caller().
		Logger().
		Hook(appHook{App: app})
}

type appHook struct {
	App string `json:"app"`
}

func (a appHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	e.Str("app", a.App)
}
