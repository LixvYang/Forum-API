package category

import (
	"encoding/json"
	"log"
	v1 "mixindev/api/v1"
	"mixindev/model"
	"mixindev/pkg/errmsg"
	"mixindev/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/net/context"
)

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	List       []model.CategoryInfo `json:"list"`
}

type CategoryHandler struct {
	ctx context.Context
	redisClient *redis.Client
}

func NewCategoryHandler(ctx context.Context, redisClient *redis.Client) *CategoryHandler {
	return &CategoryHandler{
		ctx: ctx,
		redisClient: redisClient,
	}
}

func (categoryHandler *CategoryHandler) DeleteCategoryFromRedis()  {
	log.Println("Delete category from Redis")
	categoryHandler.redisClient.Del("category")
}

// @Summary 获取分类列表
// @Description 获取分类列表
// @Tags category
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} category.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"category_name":"前端"}]}}"
// @Router /v1/category [get]
func (categoryHandler *CategoryHandler) ListCategories(c *gin.Context) {
	val, err := categoryHandler.redisClient.Get("category").Result()
	if err == redis.Nil {
		var r ListRequest
		if err := c.Bind(&r); err != nil {
			v1.SendResponse(c, errmsg.ErrBind, nil)
			return
		}

		infos, count, err := service.ListCategories(r.Offset, r.Limit)
		if err != nil {
			v1.SendResponse(c, err, nil)
			return
		}

		data, _ := json.Marshal(infos)
		categoryHandler.redisClient.Set("category", data, 0)
		v1.SendResponse(c, nil, ListResponse{
			TotalCount: uint64(count),
			List:       infos,
		})
	} else if err != nil {
		v1.SendResponse(c, err, nil)
		return
	} else {
		log.Println("Get category from Redis")
		CategoryResponse := make([]model.CategoryInfo, 0)
		if err := json.Unmarshal([]byte(val), &CategoryResponse); err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		v1.SendResponse(c, nil, CategoryResponse)
	}
}
