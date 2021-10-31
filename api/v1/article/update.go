package article

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 更改文章
// @Description 更改文章
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path int true "菜单数据的数据库id"
// @Param menu body model.Article true "The article info"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/article/{id} [put]
func UpdateArticleById(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		v1.SendResponse(c ,errmsg.ErrBind, nil)
		return
	}
	article.ID = uint(id)
	if err := article.UpdateArticleById(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}