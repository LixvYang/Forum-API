package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "The user's database id index num"
// @Param user body model.User true "The user info"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/user/{id} [put]
func UpdateUserById(c *gin.Context) {
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
