package menu

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m *model.Menu
	if err := m.DeleteMenu(id); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
