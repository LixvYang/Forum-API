package model

import (
	"gorm.io/gorm"
)

type CategoryModel struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(255) not null" json:"category_name"`
}