package controller

import (
	"github.com/flametest/go-gin-server-example/internal/logic"
	"github.com/flametest/go-gin-server-example/pkg/dto"
	"github.com/flametest/go-gin-server-example/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addGroupController(rg *gin.RouterGroup) {
	groupController := rg.Group("groups")

	groupController.GET("/:group_id", getGroupById)
}

func getGroupById(c *gin.Context) {
	req := &dto.GetGroupByIdReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
	log.Debug().Msgf("get group by id req: %v", req)
	groupLogic := logic.NewGroupLogicImpl()
	group, err := groupLogic.GetGroupById(c, req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, group)
}
