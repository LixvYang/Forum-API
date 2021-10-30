package tag

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据标签id删除标签
// @Description 根据标签id删除标签
// @Tags tag
// @Accept  json
// @Produce  json
// @Param id path int true "标签数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/tag/{id} [delete]
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tag *model.Tag
	if err := tag.DeleteTag(id); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	v1.SendResponse(c, nil, nil)
}
