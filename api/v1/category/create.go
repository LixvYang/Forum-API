package category

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	CategoryName string `json:"category_name"`
}

type CreateResponse struct {
	CategoryName string `json:"category_name"`
}

func AddCategory(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	u := model.Category{
		CategoryName: r.CategoryName,
	}

	if err := u.CreateCatrgory(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}


	//convert type
	rsp := CreateResponse(r)

	v1.SendResponse(c, nil, rsp)
}
