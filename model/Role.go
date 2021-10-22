package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string  `gorm:"type:varchar(50);not null" json:"name"`
}