package article

import (
	"log"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticleById(c *gin.Context) {
	log.Println("启动GetArticleById服务")
	id, _ := strconv.Atoi(c.Param("id"))
	art, err := model.GetArticleById(id)
	log.Println("传入参数id完成")
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.ErrArticleNotFound,
		"data":    art,
		"message": err,
	})
}
