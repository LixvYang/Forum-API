package menu

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name   string `form:"name" json:"name" xml:"name"  binding:"required"`
	Path   string `form:"path" json:"path" xml:"path" binding:"required"`
	Method string `form:"method" json:"method" xml:"method" binding:"required"`
}

type CreateResponse struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

// @Summary 创建菜单
// @Description 创建菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param menu body menu.CreateRequest true "创建新标签"
// @Success 200 {object} menu.CreateResponse "{"code":0,"message":"OK","data":{"tag_name":".."}}"
// @Router /v1/menu/add [post]
func (menuHandler *MenuHandler) AddMenu(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}

	u := model.Menu{
		Name:   r.Name,
		Method: r.Method,
		Path:   r.Path,
	}

	if err := u.CreateMenu(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	rep := CreateResponse(r)
	menuHandler.DeleteMenuFromRedis()
	v1.SendResponse(c, nil, rep)
}
