package model

import (
	"log"
	"mixindev/pkg/errmsg"

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

// type ArticleInfo struct {
// 	Id           uint64    `json:"id"`
// 	Title        string    `json:"title"`
// 	Content      string    `json:"content"`
// 	CategoryId   uint64    `json:"category_id"`
// 	CategoryName string    `json:"category_name"`
// 	TagId        uint64    `json:"tag_id"`
// 	TagName      string    `json:"tag_name"`
// 	UserId       uint64    `json:"user_id"`
// 	UserName     string    `json:"username"`
// 	Avatar       string    `json:"avatar"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"update_at"`
// }

// Create new article
func (a *Article) CreateArticle() error {
	return db.Create(&a).Error
}

// Get article by id
func GetArticleById(id int) (Article, error) {
	var a Article
	log.Println("开始查询")
	err = db.First(&a, id).Error
	log.Println("查询失败")
	if err != nil {
		log.Println("查询失败")
	}
	return a, errmsg.OK
}

// Delete article by id
func DeleteArticle(id int) error {
	var art Article
	log.Println("开始删除")
	err = db.Where("id = ?", id).Delete(&art).Error
	log.Println("完成删除")
	if err != nil {
		return err
	}
	return err
}

// Get articles list
func (a *Article) ListArticles() {

}

// 验证创建字段
