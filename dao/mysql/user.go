package mysql

import (
	"chatGPT/model"
	"context"
	"gorm.io/gorm"
)

type User struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *User {
	return &User{}
}

// UserNameExist 用户名是否已存在
func (u *User) UserNameExist(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = u.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = u.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (u *User) CreateUser(user *model.User) error {
	err := u.DB.Model(&model.User{}).Create(&user).Error
	return err
}
