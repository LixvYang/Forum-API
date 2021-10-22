package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string	`gorm:"type:varchar(255);not null" json:""`
	Content     string	`gorm:"type: ;not null" json:""`
	Category_id int	`gorm:"type: ;not null" json:""`
	Tag_id      int	`gorm:"type: ;not null" json:""`
	User_id     int	`gorm:"type: ;not null" json:""`
}
