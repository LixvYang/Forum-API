package model

import (
	"mixindev/pkg/constvar"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50) not null" json:"name"`
	Path   string `gorm:"type:varchar(50) not null" json:"path"`
	Method string `gorm:"type:varchar(50) not null" json:"method"`
}

type MenuInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

//Check menu by id
func (m *Menu) GetMenuById(id int) (*Menu, error) {
	d := db.First(&m, id)
	return m, d.Error
}

// create menu
func (m *Menu) CreateMenu() error {
	return db.Create(&m).Error
}

//delete menu
func (m *Menu) DeleteMenu(id int) error {
	return db.Delete(&m, id).Error
}

//update menu
func (m *Menu) UpdateMenu() error {
	return db.Omit("created_at").Save(m).Error
}

//list menu
func (m *Menu) ListMenus(offset, limit int) ([]Menu, int64, error) {
	t := Menu{}
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	list := make([]Menu, 0)
	var count int64
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return list, count, err
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}

	return list, count, nil

}
