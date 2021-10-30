package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "The user's database id index num"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var u *model.User
	err := u.DeleteUser(userId)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	v1.SendResponse(c, nil, nil)
}
