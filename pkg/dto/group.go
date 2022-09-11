package dto

type GetGroupByIdReq struct {
	GroupId uint64 `uri:"group_id" binding:"required"`
}

type GroupDTO struct {
}
