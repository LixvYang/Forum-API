package article

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
	TotalCount int64                `json:"totalCount"`
	List       []model.ArticleInfo `json:"list"`
}

type ArticleHandler struct {
	ctx context.Context
	redisClient *redis.Client
}


func NewArticleHandler(ctx context.Context, redisClient *redis.Client) *ArticleHandler {
	return &ArticleHandler{
		ctx: ctx,
		redisClient: redisClient,
	}
}

func (articleHandler *ArticleHandler) DeleteArticleFromRedis()  {
	log.Println("Delete article from Redis")
	articleHandler.redisClient.Del("article")
}

// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags article
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} article.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[]}"
// @Router /v1/article [get]
func (articleHandler *ArticleHandler) GetArticlesList(c *gin.Context) {
	val, err := articleHandler.redisClient.Get("articles").Result()
	if err == redis.Nil {
		var r ListRequest
		if err := c.Bind(&r); err != nil {
			v1.SendResponse(c, errmsg.ErrBind, nil)
			return
		}
		articleList, count, err := service.ListArticles(r.Offset, r.Limit)
		if err != nil {
			v1.SendResponse(c, err, nil)
			return
		}
		data, _ := json.Marshal(articleList)
		articleHandler.redisClient.Set("articles", data, 0)
		v1.SendResponse(c, nil, ListResponse{
			TotalCount: count,
			List:       articleList,
		})
	} else if err != nil {
		v1.SendResponse(c, err, nil)
		return 
	} else {
		log.Println("Get articles from Redis")
		ArticleResponse := make([]model.ArticleInfo, 0)
		if err := json.Unmarshal([]byte(val), &ArticleResponse); err != nil {
			v1.SendResponse(c, err, nil)
			return 
		}
		v1.SendResponse(c, nil, ArticleResponse)
	}
}
