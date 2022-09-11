package controller

import "github.com/gin-gonic/gin"

func addUserAPI(rg *gin.RouterGroup) {
	userAPI := rg.Group("users")

	userAPI.GET("/:user_id")
}
