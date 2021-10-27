package model

import (
	"mixindev/pkg/constvar"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(255) not null" json:"tag_name"`
}

type TagInfo struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

// 根据标签id获取标签数据.
func GetTagById(id int) (Tag, error) {
	var tag Tag
	d := db.First(&tag, id)
	return tag, d.Error
}

// 创建新标签
func (t *Tag) CreateTag() error {
	return db.Create(&t).Error
}

// 根据标签id删除标签
func DeleteTag(id int) error {
	var tag Tag
	return db.Where("id = ?", id).Delete(&tag).Error
}

//list tags
func (t *Tag) ListTags(offset,limit int) ([]Tag, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	tags := make([]Tag,0)
	var count int64

	if err := db.Model(&t).Count(&count).Error; err != nil {
		return tags, count, err
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&tags).Error; err != nil {
		return tags, count, err
	}

	return tags, count, nil
}
