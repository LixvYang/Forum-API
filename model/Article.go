package model

import (
	"mixindev/pkg/constvar"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	CategoryId int    `gorm:"column:category_id"`
	TagId      int    `gorm:"column:tag_id"`
	UserId     int    `gorm:"column:user_id"`
}

type ArticleInfo struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	TagId        int       `json:"tag_id"`
	TagName      string    `json:"tag_name"`
	UserId       int       `json:"user_id"`
	UserName     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Create new article
func (a *Article) CreateArticle() error {
	return db.Create(&a).Error
}

// Get article by id
func GetArticleById(id int) (Article, error) {
	var article Article
	d := db.First(&article, id)
	return article, d.Error
}

// Delete article by id
func DeleteArticle(id int) error {
	var article Article
	d := db.Where("id = ?", id).Delete(&article)
	return d.Error
}

// Get articles list
func (a *Article) ListArticles(offset, limit int) ([]Article, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	articles := make([]Article, 0)
	var count int64
	if err := db.Model(&a).Count(&count).Error; err != nil {
		return articles, count, err
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&articles).Error; err != nil {
		return articles, count, err
	}
	return articles, count, nil
}

// 验证创建字段
