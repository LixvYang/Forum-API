package model

import (
	"fmt"
	"mixindev/pkg/auth"
	"mixindev/pkg/constvar"

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

// Create user
func (u *User) AddUser() error {
	return db.Create(&u).Error
}

// GetUserById gets an user by the user id.
func GetUserById(id int) (User, error) {
	var user User
	d := db.First(&user, id)
	return user, d.Error
}

//Delete user
func DeleteUser(id int) error {
	var user User
	return db.Delete(&user, id).Error
}

func (u *User) UpdateUser() error {
	return db.Save(&u).Error
}

// GetUserByName gets an user by the username.
func (u *User) GetUserByName(username string) (*User, error) {
	d := db.Where("username = ?", username).First(&u)
	fmt.Println("GetUser-d", d)
	return u, d.Error
}

//List users
func (u *User) ListUser(offset, limit int) ([]User, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]User, 0)
	var count int64
	if err := db.Model(&u).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, err

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
