package article

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据标签id删除文章
// @Description 根据标签id删除文章
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path int true "文章的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/article/{id} [delete]
func (articleHandler *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var a *model.Article
	err := a.DeleteArticle(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
	}
	articleHandler.DeleteArticleFromRedis()
	v1.SendResponse(c, nil, nil)
}
