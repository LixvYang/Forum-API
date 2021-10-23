package article

import (
	. "mixindev/api/v1"
	"github.com/gin-gonic/gin"
	"strconv"

	"mixindev/pkg/errmsg"
	"mixindev/model"
)

func GetArticleById(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := model.GetArticleById(uint64(id))
	if err != nil {
		SendResponse(c, errmsg.ErrArticleNotFound,nil)
		return
	}

	articleInfo := &model.ArticleInfo{
		Id:           article.Id,
		Title:        article.Title,
		Content:      article.Content,
		CategoryId:   1,//article.CategoryId,
		CategoryName: "1",//category.CategoryName,
		TagId:        1,//article.TagId,
		TagName:      "1",//tag.TagName,
		UserId:       1,//article.UserId,
		UserName:     "1",//user.Username,
		Avatar:       "1",//user.Avatar,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}
	SendResponse(c, nil, articleInfo)
}
