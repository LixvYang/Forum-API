package tag

import (
	v1 "mixindev/api/v1"
	"mixindev/pkg/errmsg"
	"mixindev/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTagById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, err := model.GetTagById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrUserNotFound, nil)
		return
	}
	v1.SendResponse(c, nil, model.TagInfo{Id: int(tag.ID), TagName: tag.TagName})
}