package role

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
	List       []model.RoleInfo `json:"list"`
}

type RoleHandler struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewRoleHandler(ctx context.Context, redisClient *redis.Client) *RoleHandler {
	return &RoleHandler{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (roleHandler *RoleHandler) DeleteRoleFromRedis() {
	log.Println("Remove data from Redis")
	roleHandler.redisClient.Del("roles")
}

// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags role
// @Accept  json
// @Produce  json
// @Param   offset      query    int     false     "Offset"
// @Param   limit      query    int     false     "Limit"
// @Success 200 {object} role.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"tag_name":"..."}]}}"
// @Router /v1/role [get]
func (roleHandler *RoleHandler) ListRoles(c *gin.Context) {
	val, err := roleHandler.redisClient.Get("roles").Result()
	if err == redis.Nil {
		var r ListRequest
		if err := c.ShouldBindQuery(&r); err != nil {
			v1.SendResponse(c, errmsg.ErrBind, nil)
			return
		}

		if r.Limit == 0 {
			r.Limit = 10
		}

		infos, count, err := service.ListRoles(r.Offset, r.Limit)
		if err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		data, _ := json.Marshal(infos)
		roleHandler.redisClient.Set("roles", string(data), 0)

		v1.SendResponse(c, nil, ListResponse{
			TotalCount: uint64(count),
			List:       infos,
		})
	} else if err != nil {
		v1.SendResponse(c, err, nil)
		return 
	} else {
		log.Println("Request to Redis for roles")
		RoleResponse := make([]model.RoleInfo, 0)
		if err := json.Unmarshal([]byte(val), &RoleResponse); err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		v1.SendResponse(c, nil, RoleResponse)
	}
}
