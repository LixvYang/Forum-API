package menu

import (
	"fmt"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
