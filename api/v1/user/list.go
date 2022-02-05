package user

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
	TotalCount int              `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}

type UserHandler struct {
	ctx context.Context
	redisClient *redis.Client
}

func NewUserHandler(ctx context.Context, redisClient *redis.Client) *UserHandler {
	return &UserHandler{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

// @Summary List the users in the database
// @Security ApiKeyAuth
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} user.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /v1/user [get]
func (userHandler *UserHandler) ListUsers(c *gin.Context) {
	val, err := userHandler.redisClient.Get("users").Result()
	if err == redis.Nil {
		var r ListRequest
		if err := c.Bind(&r); err != nil {
			v1.SendResponse(c, errmsg.ErrBind, nil)
			return
		}
	
		infos, count, err := service.ListUsers(r.Offset, r.Limit)
		if err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		data, _ := json.Marshal(infos)
		userHandler.redisClient.Set("users", string(data), 0)
		v1.SendResponse(c, nil, ListResponse{
			TotalCount: int(count),
			UserList:   infos,
		})
	} else if err != nil {
		v1.SendResponse(c, err, nil)
		return
	} else {
		log.Printf("Request to Redis")
		UserResponse := make([]model.UserInfo, 0)
		json.Unmarshal([]byte(val), &UserResponse)
		v1.SendResponse(c, nil, UserResponse)
	}

}
