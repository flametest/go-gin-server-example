package db

import (
	"context"
	"gorm.io/gorm"
)

const UserDOTableName = "user"

type UserDO struct {
	Base
	FirstName string `gorm:"type:varchar(32);column:first_name"`
	LastName  string `gorm:"type:varchar(32);column:last_name"`
}

func (UserDO) TableName() string {
	return UserDOTableName
}

var UserDao IUserDao

type IUserDao interface {
	SaveUser(ctx context.Context, do *UserDO) (uint64, error)
	GetUserById(ctx context.Context, userId uint64) (*UserDO, error)
}

type userDaoImpl struct {
	*gorm.DB
}

func newUserDaoImpl(DB *gorm.DB) *userDaoImpl {
	return &userDaoImpl{DB: DB}
}

func (u *userDaoImpl) SaveUser(ctx context.Context, do *UserDO) (uint64, error) {
	if err := u.WithContext(ctx).Save(do).Error; err != nil {
		return 0, err
	}
	return do.Id, nil
}

func (u *userDaoImpl) GetUserById(ctx context.Context, userId uint64) (*UserDO, error) {
	userDO := &UserDO{}
	if err := u.WithContext(ctx).Where("id=?", userId).First(userDO).Error; err != nil {
		return nil, err
	}
	return userDO, nil
}
