package category

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据标签id删除分类
// @Description 根据标签id删除分类
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path int true "标签数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/category/{id} [delete]
func (categoryHandler *CategoryHandler) DeleteCategory(c *gin.Context) {
	CategoryId, _ := strconv.Atoi(c.Param("id"))
	var category *model.Category
	if err := category.DeleteCategory(CategoryId); err != nil {
		v1.SendResponse(c, errmsg.ErrDatabase, nil)
		return
	}
	categoryHandler.DeleteCategoryFromRedis()
	v1.SendResponse(c, nil, nil)
}
