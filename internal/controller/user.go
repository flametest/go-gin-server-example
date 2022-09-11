package controller

import (
	"github.com/flametest/go-gin-server-example/internal/logic"
	"github.com/flametest/go-gin-server-example/pkg/dto"
	"github.com/flametest/go-gin-server-example/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserController(rg *gin.RouterGroup) {
	userController := rg.Group("users")

	userController.GET("/:user_id", getUserById)
}

func getUserById(c *gin.Context) {
	req := &dto.GetUserByIdReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		return
	}
	log.Debug().Msgf("get user by id req: %v", req)
	userLogic := logic.NewUserLogicImpl()
	user, err := userLogic.GetUserById(c, req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}
