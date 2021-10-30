package category

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 用分类id获取单个分类信息
// @Description 用分类id获取单个分类信息
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.CategoryInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":".."}}"
// @Router /v1/category/{id} [get]
func GetCategoryById(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	var category *model.Category
	category, err := category.GetCategoryById(categoryId)
	if err != nil {
		v1.SendResponse(c, errmsg.ErrCategoryNotFound, nil)
		return
	}

	v1.SendResponse(c,
		nil,
		&model.CategoryInfo{
			Id:           uint64(category.ID),
			CategoryName: category.CategoryName})
}
