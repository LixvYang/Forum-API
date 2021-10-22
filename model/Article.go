package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string	`gorm:"type:varchar(255);not null" json:"title"`
	Content     string	`gorm:"type:varchar(255);not null" json:"content"`
	CategoryId int	`gorm:"type:int;not null" json:"catefory_id"`
	TagId      int	`gorm:"type:int;not null" json:"tag_id"`
	UserId     int	`gorm:"type:int;not null" json:"user_id"`
}
