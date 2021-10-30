package user

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/auth"
	"mixindev/pkg/errmsg"
	"mixindev/pkg/token"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
}

// @Summary Login generates the authentication token
// @Description 登录
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body user.LoginRequest true "login"
// @Success 200 {object} user.LoginResponse "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /v1/login [post]
func Login(c *gin.Context) {
	var u LoginRequest
	var user model.User
	if err := c.Bind(&u); err != nil {
		fmt.Println("err", err)
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	d, err := user.GetUserByName(u.Username)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		v1.SendResponse(c, errmsg.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: uint64(d.ID), Username: d.Username}, "")
	if err != nil {
		v1.SendResponse(c, errmsg.ErrToken, nil)
		return
	}

	v1.SendResponse(c, nil, model.Token{Token: t})
}
