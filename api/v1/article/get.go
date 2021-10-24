package article

import (
	// v1 "mixindev/api/v1"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mixindev/model"
	"mixindev/pkg/errmsg"
)

func GetArticleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := model.GetArticleById(uint64(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  err,
		"data":    data,
		"message": errmsg.ErrArticleNotFound,
	})
}

func TableName(c *gin.Context) {
	tablename := model.TableName()
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    tablename,
		"message": errmsg.ErrArticleNotFound,
	})
}
