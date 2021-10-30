package article

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"mixindev/service"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount int64                `json:"totalCount"`
	List       []model.ArticleInfo `json:"list"`
}

// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags article
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} article.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[]}"
// @Router /v1/article [get]
func GetArticlesList(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}
	articleList, count, err := service.ListArticles(r.Offset, r.Limit)

	if err != nil {
		v1.SendResponse(c, err, nil)
		return
	}
	v1.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       articleList,
	})
}
