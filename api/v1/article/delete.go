package article

import (
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.OK,
		"message": err,
	})
}
