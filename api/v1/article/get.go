package article

import (
	v1 "mixindev/api/v1"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"

	"mixindev/pkg/errmsg"
	"mixindev/model"
)

func GetArticleById(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := model.GetArticleById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrArticleNotFound,nil)
		return
	}

	articleInfo := model.ArticleInfo{
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
	v1.SendResponse(c, nil, articleInfo)
}

func TableName(c *gin.Context)  {
	tablename := model.TableName()
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    tablename,
		"message": errmsg.ErrArticleNotFound,
	})
}