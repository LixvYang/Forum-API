package model

import (
	"time"
)

type BaseModel struct {
	Id        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt time.Time `gorm:"index"`
}
