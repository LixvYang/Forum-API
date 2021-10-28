package middleware

import (
	v1 "mixindev/api/v1"
	"mixindev/pkg/errmsg"
	"mixindev/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Parse the JWT
		if _, err := token.ParseRequest(c); err != nil {
			v1.SendResponse(c, errmsg.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
