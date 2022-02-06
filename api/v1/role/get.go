package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取单个角色
// @Description 获取单个角色
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path int true "角色数据的数据库id"
// @Success 200 {object} model.RoleInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"..."}}"
// @Router /v1/role/{id} [get]
func (roleHandler *RoleHandler) GetRoleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role *model.Role
	l, err := role.GetRoleById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrRoleGet, nil)
		return
	}
	v1.SendResponse(c, nil, &model.RoleInfo{Id: uint64(l.ID), Name: l.Name})
}
