package main

import (
	"context"
	"github.com/flametest/go-gin-server-example/internal/config"
	"github.com/flametest/go-gin-server-example/internal/controller"
	"github.com/flametest/go-gin-server-example/pkg/log"
	"github.com/flametest/go-gin-server-example/pkg/server"
	"github.com/rs/zerolog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	defer stop()

	log.Initialize("go-gin-server-example", zerolog.DebugLevel)
	log.Debug().Msg("log initialization complete.")

	config.InitConfig()
	log.Debug().Msg("config initialization complete.")

	router := controller.NewRouter()
	srv := server.NewServer(router)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()

	stop()
	log.Info().Msg("shutting down gracefully...")

	// The context is used to inform the server it has 5 seconds to finish, the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown: ", err)
	}

	log.Info().Msg("Server exiting")
}
