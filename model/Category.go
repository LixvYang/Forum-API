package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(255) not null" json:"category_name"`
}

// 根据标签id获取标签数据.
func GetCategoryById(id int) (Category, error) {
	var category Category
	d := db.First(&category,id)
	return category, d.Error
}