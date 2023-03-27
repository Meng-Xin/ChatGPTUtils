package mysql

import (
	"chatGPT/dao"
	"chatGPT/model"
	"context"
	"gorm.io/gorm"
)

type Chat struct {
	*gorm.DB
}

func NewChatDB(ctx context.Context) *Chat {
	return &Chat{dao.NewDao(ctx)}
}

// SaveScenes 保存场景
func (c *Chat) SaveScenes(uid uint, chatScenes *model.Chat) (err error) {
	var count int64
	err = c.DB.Model(model.User{}).Where("id=?", uid).Count(&count).Error
	if count == 0 {
		return err
	}
	chatScenes.UserId = uid
	// 否则保存本次场景信息
	err = c.DB.Model(&model.Chat{}).Create(&chatScenes).Error
	return err
}
