package user

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
	TotalCount int              `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}

func ListUsers(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUsers(r.Offset, r.Limit)
	if err != nil {
		v1.SendResponse(c, err, nil)
		return
	}
	v1.SendResponse(c, nil, ListResponse{
		TotalCount: int(count),
		UserList:   infos,
	})
}
