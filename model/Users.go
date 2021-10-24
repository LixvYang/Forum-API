package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Avatar   string `gorm:"type:varchar(255);not null" json:"avatar"  label:"角色头像"`
}

// GetUserById gets an user by the user id.
func (u User) GetUserById(id uint64) (User, error) {
	d := db.Where("id = ?", id).First(&u)
	fmt.Println("GetUser-d", d)
	return u, d.Error
}
