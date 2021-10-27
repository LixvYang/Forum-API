package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRole(c *gin.Context) {
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

	v1.SendResponse(c, nil, nil)
}
