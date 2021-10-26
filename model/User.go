package model

import (
	"mixindev/pkg/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Avatar   string `gorm:"type:varchar(255);not null" json:"avatar"  label:"角色头像"`
}

type UserInfo struct {
	Id       int
	Username string
	Avatar   string
}

// GetUserById gets an user by the user id.
func GetUserById(id int) (User, error) {
	var user User
	d := db.First(&user, id)
	return user, d.Error
}

// Create user
func (u *User) AddUser() error {
	return db.Create(&u).Error
}

//Delete user
func DeleteUser(id int) error {
	var user User
	return db.Delete(&user, id).Error
}

// GetUserByName gets an user by the username.
func (u *User) GetUserByName(username string)  {
	
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

//Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
