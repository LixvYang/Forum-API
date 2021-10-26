package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteUser(userId)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	v1.SendResponse(c, nil, nil)
}
