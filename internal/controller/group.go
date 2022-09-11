package controller

import "github.com/gin-gonic/gin"

func addGroupAPI(rg *gin.RouterGroup) {
	groupAPI := rg.Group("groups")

	groupAPI.GET("/:group_id")
}
