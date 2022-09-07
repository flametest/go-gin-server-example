package log

import (
	"github.com/rs/zerolog"
	"os"
)

var logger zerolog.Logger

func Initialize(app string, logLevel zerolog.Level) {
	logger = zerolog.
		New(os.Stderr).
		Level(logLevel).
		With().Timestamp().
		Logger().
		Hook(appHook{App: app})
}

type appHook struct {
	App string `json:"app"`
}

func (a appHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	e.Str("app", a.App)
}
