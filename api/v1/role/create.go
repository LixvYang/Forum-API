package role

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

type CreateResponse struct {
	Name string `json:"name"`
}

func AddRole(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	u := model.Role{
		Name: r.Name,
	}

	if err := u.CreateRole(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse(r)

	// Show the user information.
	v1.SendResponse(c, nil, rsp)
}
