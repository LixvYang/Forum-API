package article

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用文章ID获取单页信息
func GetArticleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := model.GetArticleById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrArticleNotFound, nil)
	}
	user, uErr := model.GetUserById(article.UserId)
	category, cErr := model.GetCategoryById(article.CategoryId)
	tag, tErr := model.GetTagById(article.TagId)
	if uErr != nil || cErr != nil || tErr != nil {
		v1.SendResponse(c, errmsg.ErrArticleNotFound, nil)
		return
	}
	articleInfo := &model.ArticleInfo{
		Id:           int(article.ID),
		Title:        article.Title,
		Content:      article.Content,
		CategoryId:   article.CategoryId,
		CategoryName: category.CategoryName,
		TagId:        article.TagId,
		TagName:      tag.TagName,
		UserId:       article.UserId,
		UserName:     user.Username,
		Avatar:       user.Avatar,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}

	v1.SendResponse(c, nil, articleInfo)
}
