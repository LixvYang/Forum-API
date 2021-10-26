package user

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

func AddUser(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	u := model.User{
		Username: r.Username,
		Password: r.Password,
		Avatar:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmb_j0g0GkHwuF42L8gNpJtSrT10kRQyEMvQ&usqp=CAU",
	}

	if err := u.Encrypt(); err != nil {
		v1.SendResponse(c, errmsg.ErrEncrypt, nil)
		return
	}

	if err := u.AddUser(); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	v1.SendResponse(c, nil, rsp)
}
