package menu

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据菜单id删除菜单
// @Description 根据菜单id删除菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path int true "菜单数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/menu/{id} [delete]
func (menuHandler *MenuHandler) DeleteMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m *model.Menu
	if err := m.DeleteMenu(id); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	
	menuHandler.DeleteMenuFromRedis()
	v1.SendResponse(c, nil, nil)
}
