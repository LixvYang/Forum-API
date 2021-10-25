package article

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteArticle(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
	}
	v1.SendResponse(c, nil, nil)
}
