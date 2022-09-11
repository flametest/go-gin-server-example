package controller

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	addUserAPI(v1)
	addGroupAPI(v1)
	return router
}
