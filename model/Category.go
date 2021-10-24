package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(255) not null" json:"category_name"`
}

// 根据标签id获取标签数据.
func (c Category) GetCategoryById(id uint64) (Category, error) {
	d := db.Where("id = ?", id).First(&c)
	return c, d.Error
}