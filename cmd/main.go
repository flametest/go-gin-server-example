package main

import (
	"github.com/flametest/go-gin-server-example/internal/config"
	"github.com/flametest/go-gin-server-example/pkg/log"
	"github.com/rs/zerolog"
)

func main() {
	config.InitConfig()
	log.Initialize("go-gin-server-example", zerolog.DebugLevel)
	log.Debug().Str("xx", "yy").Send()
}
