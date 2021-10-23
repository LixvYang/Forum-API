package model

import (
	"time"
	"mixindev/pkg/errmsg"

)

type ArticleModel struct {
	BaseModel
	Title       string	`gorm:"type:varchar(255);not null" json:"title"`
	Content     string	`gorm:"type:varchar(255);not null" json:"content"`
	CategoryId int	`gorm:"type:int;not null" json:"category_id"`
	TagId      int	`gorm:"type:int;not null" json:"tag_id"`
	UserId     int	`gorm:"type:int;not null" json:"user_id"`
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
func (a *ArticleModel) TableName() string {
	return "article"
}

// Create new article
func (a *ArticleModel) CreateArticle(data *ArticleModel)  {
	
}

// Get article by id
func GetArticleById(id uint64) (ArticleModel, error) {
	var article ArticleModel
	err := db.First(&article, id)
	if err != nil {
		return article, errmsg.ErrArticleNotFound
	}
	return 	article, errmsg.OK
}

// Delete article by id
func (a *ArticleModel) DeleteArticle(id uint64)  {
	
}

// Get articles list
func (a *ArticleModel) ListArticles()  {
	
}

// 验证创建字段
