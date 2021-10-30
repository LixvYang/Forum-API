package menu

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 更改菜单
// @Description 更改菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path int true "菜单数据的数据库id"
// @Param menu body model.Menu true "The user info"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/menu/{id} [put]
func UpdateMenuById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
		return
	}
	menu.ID = uint(id)

	if err := menu.UpdateMenu(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
