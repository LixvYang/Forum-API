package menu

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMenuById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m *model.Menu
	menu, err := m.GetMenuById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrMenuGet, nil)
		return
	}
	v1.SendResponse(c, 
					nil, 
					model.MenuInfo{
						Id: int(menu.ID), 
						Name: menu.Name, 
						Path: menu.Path, 
						Method: menu.Method,
					})
}
