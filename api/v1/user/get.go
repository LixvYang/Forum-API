package user

import (
	// "encoding/json"
	// "log"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"mixindev/pkg/token"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis"
)

// @Summary Get an user by the user id
// @Description Get an user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.UserInfo "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user/{id} [get]
func (userHandler *UserHandler) GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var u *model.User
	// val, err := userHandler.redisClient.Get("UserId:").Result()
	// if err == redis.Nil {
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
	// data, _ := json.Marshal(res)
	// userHandler.redisClient.Set("UserId:", string(data), 0)
	v1.SendResponse(c, nil, res)
	// } else if err != nil {
	// 	v1.SendResponse(c, err, nil)
	// 	return
	// } else {
	// 	log.Println("Request to Redis for usersId")
	// 	UserResponse := make([]model.UserInfo, 0)
	// 	json.Unmarshal([]byte(val), &UserResponse)
	// 	v1.SendResponse(c, nil, UserResponse)
	// }
	
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
