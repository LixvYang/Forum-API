package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"mixindev/pkg/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var u *model.User

	user, err := u.GetUserById(userId)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrUserNotFound, nil)
		return
	}
	res := &model.UserInfo{
		Id:       int(user.ID),
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	v1.SendResponse(c, nil, res)
}

func GetUseInfo(c *gin.Context) {
	res, _ := token.ParseRequest(c)
	var u *model.User
	user, err := u.GetUserById(int(res.ID))
	if err != nil {
		v1.SendResponse(c, errmsg.ErrUserNotFound, nil)
		return
	}

	resUser := &model.UserInfo{Id: int(user.ID), Username: user.Username, Avatar: user.Avatar}
	v1.SendResponse(c, nil, resUser)
}
