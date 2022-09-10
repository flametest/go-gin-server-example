package server

import (
	"github.com/flametest/go-gin-server-example/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GinServer struct {
	http.Server
}

func NewServer(router *gin.Engine) *GinServer {
	server := http.Server{
		Addr:         config.Cfg.HttpHost,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &GinServer{
		server,
	}
}
