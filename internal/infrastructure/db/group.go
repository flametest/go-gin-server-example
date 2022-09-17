package db

import (
	"context"
	"gorm.io/gorm"
)

const GroupDOTableName = "group"

type GroupDO struct {
	Base
	Name string `gorm:"type:varchar(16);column:name"`
}

func (GroupDO) TableName() string {
	return GroupDOTableName
}

var GroupDao IGroupDao

type IGroupDao interface {
	SaveGroup(ctx context.Context, do *GroupDO) (uint64, error)
	GetGroupById(ctx context.Context, groupId uint64) (*GroupDO, error)
}

type groupDaoImpl struct {
	*gorm.DB
}

func (g *groupDaoImpl) SaveGroup(ctx context.Context, do *GroupDO) (uint64, error) {
	if err := g.WithContext(ctx).Save(do).Error; err != nil {
		return 0, err
	}
	return do.Id, nil
}

func (g *groupDaoImpl) GetGroupById(ctx context.Context, groupId uint64) (*GroupDO, error) {
	groupDO := &GroupDO{}
	if err := g.WithContext(ctx).Where("id=?", groupId).First(groupDO).Error; err != nil {
		return nil, err
	}
	return groupDO, nil
}
