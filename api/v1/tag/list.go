package tag

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"mixindev/service"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Offset int `form:"offset" json:"offset" xml:"offset"`
	Limit  int `form:"limit" json:"limit" xml:"limit"`
}

type ListResponse struct {
	TotalCount uint64          `json:"totalCount"`
	List       []model.TagInfo `json:"list"`
}




// @Summary 获取标签列表
// @Description 获取标签列表
// @Tags tag
// @Accept  json
// @Produce  json
// @Param   offset      query    int     false     "Offset"
// @Param   limit      query    int     false      "Limit"
// @Success 200 {object} tag.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"tag_name":"..."}]}}"
// @Router /v1/tag [get]
func ListTags(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	infos, count, err := service.ListTags(r.Offset, r.Limit)
	if err != nil {
		v1.SendResponse(c, err, nil)
		return
	}

	v1.SendResponse(c, nil, ListResponse{
		TotalCount: uint64(count),
		List:       infos,
	})
}
