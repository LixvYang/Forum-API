package menu

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
	TotalCount uint64           `json:"totalCount"`
	List       []model.MenuInfo `json:"list"`
}

func ListMenus(c *gin.Context) {
	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	if r.Limit == 0 {
		r.Limit = 20
	}

	infos, count, err := service.ListMenus(r.Offset, r.Limit)
	if err != nil {
		v1.SendResponse(c, err, nil)
		return
	}
	v1.SendResponse(c, nil, ListResponse{
		TotalCount: uint64(count),
		List:       infos,
	})
}
