package menu

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取单个菜单
// @Description 获取单个菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path int true "菜单数据的数据库id"
// @Success 200 {object} model.MenuInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"..."}}"
// @Router /v1/menu/{id} [get]
func (menuHandler *MenuHandler) GetMenuById(c *gin.Context) {
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
