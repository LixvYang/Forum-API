package tag

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteTag(id); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
