package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(255) not null" json:"tag_name"`
}

// 根据标签id获取标签数据.
func GetTagById(id int) (Tag, error) {
	var tag Tag
	d := db.First(&tag, id)
	return tag, d.Error
}
