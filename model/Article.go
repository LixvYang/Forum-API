package model

import (
	"mixindev/pkg/errmsg"
	"time"
)

type Article struct {
	BaseModel
	Title      string `gorm:"column:title;type:varchar(255);not null" json:"title" binding:"required"`
	Content    string `gorm:"column:content;type:varchar(255);not null" json:"content" binding:"required"`
	CategoryId uint64 `gorm:"column:category_id;bigint(20);not null" json:"category_id" binding:"required"`
	TagId      uint64 `gorm:"column:tag_id;bigint(20);not null" json:"tag_id" binding:"required"`
	UserId     uint64 `gorm:"column:user_id;bigint(20);not null" json:"user_id" binding:"required"`
}

type ArticleInfo struct {
	Id           uint64    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryId   uint64    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	TagId        uint64    `json:"tag_id"`
	TagName      string    `json:"tag_name"`
	UserId       uint64    `json:"user_id"`
	UserName     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
}

// Return Table Name
func TableName() string {
	return "article"
}

// Create new article
func (a *Article) CreateArticle() error {
	return db.Create(&a).Error
}

// Get article by id
func GetArticleById(id uint64) (Article, error) {
	var article Article
	err := db.First(&article).Error
	if err != nil {
		return article, errmsg.ErrArticleNotFound
	}
	return article, errmsg.OK
}

// Delete article by id
func (a *Article) DeleteArticle(id uint64) {

}

// Get articles list
func (a *Article) ListArticles() {

}

// 验证创建字段
