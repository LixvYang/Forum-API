package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	l, err := model.GetRoleById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrRoleGet, nil)
		return
	}
	v1.SendResponse(c, nil, &model.RoleInfo{Id: uint64(l.ID), Name: l.Name})
}
