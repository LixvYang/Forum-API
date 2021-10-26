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
