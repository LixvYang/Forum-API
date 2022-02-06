package model

import (
	"mixindev/pkg/constvar"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(255) not null" json:"category_name"`
}

type CategoryInfo struct {
	Id           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
}

//Create category
func (c *Category) CreateCatrgory() error {
	return db.Create(&c).Error
}

//list categories
func (c *Category) ListCategories(offset,limit int) ([]Category, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categories := make([]Category,0)
	var count int64

	if err := db.Model(&c).Count(&count).Error; err != nil {
		return categories, count, err
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&categories).Error; err != nil {
		return categories, count, err
	}

	return categories, count, nil
}

// 根据标签id获取标签数据.
func (c *Category) GetCategoryById(id int) (*Category, error) {
	d := db.Find(&c, id)
	return c, d.Error
}

// Delete category by id
func (c *Category) DeleteCategory(id int) error {
	return db.Where("id = ?", id).Delete(&c).Error
}


