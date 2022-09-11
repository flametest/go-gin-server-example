package server

import (
	"github.com/flametest/go-gin-server-example/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GinHttpServer struct {
	http.Server
}

func NewHttpServer(router *gin.Engine) *GinHttpServer {
	server := http.Server{
		Addr:         config.Cfg.HttpHost,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &GinHttpServer{
		server,
	}
}
