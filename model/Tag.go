package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(255) not null" json:"tag_name"`
}

// 根据标签id获取标签数据.
func (t Tag) GetTagById(id uint64) (Tag, error) {
	d := db.First(&t, id)
	return t, d.Error
}