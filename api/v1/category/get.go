package category

import (
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryById(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))

	category, err := model.GetCategoryById(categoryId)
	if err != nil {
		v1.SendResponse(c,errmsg.ErrCategoryNotFound,nil)
		return
	}

	v1.SendResponse(c,
		nil,
		&model.CategoryInfo{
			Id: uint64(category.ID), 
			CategoryName: category.CategoryName})
}