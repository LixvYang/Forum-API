package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(255) not null" json:"tag_name"`
}