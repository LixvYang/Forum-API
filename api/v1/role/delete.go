package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据角色id删除角色
// @Description 根据角色id删除角色
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path int true "角色数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/role/{id} [delete]
func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role *model.Role
	if err := role.DeleteRole(id); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
