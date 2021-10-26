package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	if err := c.Bind(&user); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	user.ID = uint(userId)

	if err := user.Encrypt(); err != nil {
		v1.SendResponse(c, errmsg.ErrEncrypt, nil)
		return
	}

	if err := user.UpdateUser(); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
