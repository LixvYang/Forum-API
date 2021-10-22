package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name string `gorm:"type:varchar(50) not null" json:"name"`
	Path string `gorm:"type:varchar(50) not null" json:"path"`
	Method string `gorm:"type:varchar(50) not null" json:"method"`
}