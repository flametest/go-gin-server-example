package dto

type GetUserByIdReq struct {
	UserId uint64 `uri:"user_id" binding:"required"`
}

type UserDTO struct {
}
