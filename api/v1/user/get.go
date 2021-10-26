package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := model.GetUserById(userId)
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

}
