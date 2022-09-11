package logic

import (
	"context"
	"github.com/flametest/go-gin-server-example/pkg/dto"
)

type UserLogic interface {
	GetUserById(ctx context.Context, req *dto.GetUserByIdReq) (*dto.UserDTO, error)
}

type userLogicImpl struct {
}

func NewUserLogicImpl() *userLogicImpl {
	return &userLogicImpl{}
}

func (u userLogicImpl) GetUserById(ctx context.Context, req *dto.GetUserByIdReq) (*dto.UserDTO, error) {
	return &dto.UserDTO{}, nil
}
