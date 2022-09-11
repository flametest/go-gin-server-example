package logic

import (
	"context"
	"github.com/flametest/go-gin-server-example/pkg/dto"
)

type GroupLogic interface {
	GetGroupById(ctx context.Context, req *dto.GetGroupByIdReq) (*dto.GroupDTO, error)
}

type groupLogicImpl struct {
}

func NewGroupLogicImpl() *groupLogicImpl {
	return &groupLogicImpl{}
}

func (g groupLogicImpl) GetGroupById(ctx context.Context, req *dto.GetGroupByIdReq) (*dto.GroupDTO, error) {
	return &dto.GroupDTO{}, nil
}
