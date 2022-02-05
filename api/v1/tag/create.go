package tag

import (
	"fmt"
	"log"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	TagName string `form:"tag_name" json:"tag_name" xml:"tag_name" binding:"required"`
}

type CreateResponse struct {
	TagName string `json:"tag_name"`
}

// @Summary 创建标签
// @Description 创建标签
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tags body tag.CreateRequest true "创建新标签"
// @Success 200 {object} tag.CreateResponse "{"code":0,"message":"OK","data":{"tag_name":"..."}}"
// @Router /v1/tag/add [post]
func (tagHandler *TagHandler) AddTag(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ErrBind, nil)
	}
	u := model.Tag{
		TagName: r.TagName,
	}

	if err := u.CreateTag(); err != nil {
		fmt.Println(err)
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}

	//convert r to CreateResponse
	rsp := CreateResponse(r)
	log.Println("Remove data from Redis")
	tagHandler.redisClient.Del("tags")

	// Show the user information.
	v1.SendResponse(c, nil, rsp)

}
