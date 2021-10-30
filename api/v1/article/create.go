package article

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId int    `json:"category_id"`
	TagId      int    `json:"tag_id"`
	UserId     int    `json:"user_id"`
}

// @Summary 创建文章
// @Description 创建文章
// @Tags article
// @Accept  json
// @Produce  json
// @Param article body article.CreateRequest true "创建文章"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/article/add [post]
func AddArticle(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}
	u := model.Article{
		Title:      r.Title,
		Content:    r.Content,
		CategoryId: r.CategoryId,
		TagId:      r.TagId,
		UserId:     r.UserId,
	}

	if err := u.Validate(); err != nil {
		v1.SendResponse(c, errmsg.ErrValidation, nil)
		return
	}
	// 检验字段的合法性
	if valid := ValidateCreateArticle(u.UserId, u.CategoryId, u.TagId); !valid {
		v1.SendResponse(c, errmsg.ErrValidation, nil)
		return
	}

	// 写入数据库
	if err := u.CreateArticle(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}



	v1.SendResponse(c, nil, nil)
}
