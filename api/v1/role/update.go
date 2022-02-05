package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 更新角色数据
// @Description 更新角色数据
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path int true "角色数据的数据库id"
// @Param role body model.Role true "role"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/role/{id} [put]
func (roleHandler *RoleHandler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	role.ID = uint(id)
	if err := role.UpdateRole(); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	// Delete role from redis
	roleHandler.DeleteRoleFromRedis()
	v1.SendResponse(c, nil, nil)
}
