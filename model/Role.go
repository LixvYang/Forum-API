package model

import (
	"mixindev/pkg/constvar"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}

type RoleInfo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

//get role by id
func GetRoleById(id int) (Role, error) {
	var role Role
	d := db.First(&role, id)
	return role, d.Error
}

//delete role
func DeleteRole(id int) error {
	var role Role
	return db.Delete(&role, id).Error
}

//create role
func (r *Role) CreateRole() error {
	return db.Create(&r).Error
}

// update role
func (r *Role) UpdateRole() error {
	return db.Omit("created_at").Save(r).Error
}

// list role
func (r *Role) ListRole(offset, limit int) ([]Role, int64, error) {
	t := Role{}
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	list := make([]Role, 0)
	var count int64
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return list, count, err
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}
	return list, count, nil

}
