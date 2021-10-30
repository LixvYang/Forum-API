package article

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 用文章id获取单篇信息
// @Description 用文章id获取单篇文章信息
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.ArticleInfo "{"code":0,"message":"OK","data":{}}"
// @Router /v1/article/{id} [get]
func GetArticleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var art *model.Article
	var cate *model.Category
	var tag *model.Tag
	var user *model.User
	article, err := art.GetArticleById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrArticleNotFound, nil)
	}
	user, uErr := user.GetUserById(article.UserId)
	category, cErr := cate.GetCategoryById(article.CategoryId)
	tag, tErr := tag.GetTagById(article.TagId)
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
