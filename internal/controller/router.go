package controller

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userAPI := v1.Group("users")
		{
			userAPI.GET("/:user_id")
		}
		groupAPI := v1.Group("groups")
		{
			groupAPI.GET("/:group_id")
		}
	}
	return router
}
