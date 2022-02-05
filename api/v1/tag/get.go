package tag

import (
	v1 "mixindev/api/v1"
	"mixindev/pkg/errmsg"
	"mixindev/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 用标签id获取单个标签信息
// @Description 用标签id获取单个标签信息
// @Tags tag
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.TagInfo "{"code":0,"message":"OK","data":{"id":0,"tag_name":"..."}}"
// @Router /v1/tag/{id} [get]
func (tagHandler *TagHandler) GetTagById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t *model.Tag
	tag, err := t.GetTagById(id)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrUserNotFound, nil)
		return
	}
	v1.SendResponse(c, nil, model.TagInfo{Id: int(tag.ID), TagName: tag.TagName})
}