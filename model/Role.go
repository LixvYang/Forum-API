package model

import (
	"gorm.io/gorm"
)

type RoleModel struct {
	gorm.Model
	Name string  `gorm:"type:varchar(50);not null" json:"name"`
}