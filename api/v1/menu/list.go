package menu

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
	Offset int `form:"offset" json:"offset" xml:"offset"`
	Limit  int `form:"limit" json:"limit" xml:"limit"`
}

type ListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	List       []model.MenuInfo `json:"list"`
}

type MenuHandler struct {
	ctx context.Context
	redisClient *redis.Client
}

func NewMenuHandler(ctx context.Context, redisClient *redis.Client) *MenuHandler {
	return &MenuHandler{
		ctx: 			 ctx,
		redisClient: redisClient,
	}
}

func (menuHandler *MenuHandler) DeleteMenuFromRedis()  {
	log.Println("Delete menu from Redis")
	menuHandler.redisClient.Del("menus")
}

// @Summary 获取菜单列表
// @Description 获取菜单列表
// @Tags menu
// @Accept  json
// @Produce  json
// @Param   offset      query    int     false     "Offset"
// @Param   limit      query    int     false      "Limit"
// @Success 200 {object} menu.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"tag_name":"..."}]}}"
// @Router /v1/menu [get]
func (menuHandler *MenuHandler) ListMenus(c *gin.Context) {
	val, err := menuHandler.redisClient.Get("menus").Result()
	if err == redis.Nil {
		var r ListRequest
		if err := c.ShouldBindQuery(&r); err != nil {
			v1.SendResponse(c, errmsg.ErrBind, nil)
			return
		}

		if r.Limit == 0 {
			r.Limit = 20
		}

		infos, count, err := service.ListMenus(r.Offset, r.Limit)
		if err != nil {
			v1.SendResponse(c, err, nil)
			return
		}

		data, _ := json.Marshal(infos)
		menuHandler.redisClient.Set("menus", data, 0)

		v1.SendResponse(c, nil, ListResponse{
			TotalCount: uint64(count),
			List:       infos,
		}) 
	} else if err != nil {
		v1.SendResponse(c, err, nil)
		return
	} else {
		log.Println("Request to Redis for menus")
		MenuResponse := make([]model.MenuInfo, 0)
		if err := json.Unmarshal([]byte(val), &MenuResponse); err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		v1.SendResponse(c, nil, MenuResponse)
	}
}
