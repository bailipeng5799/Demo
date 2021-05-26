package middleware

import (
	"github.com/gin-gonic/gin"
	"jxbdproject/common"
	"net/http"
)

func RateLimitMiddleware()gin.HandlerFunc{
	limiter := common.NewConnLimiter(2)
	return func(c *gin.Context) {
		//如果没拿到token
		if !limiter.GetConn(){
			//common.SendErrorResponse(c,http.StatusTooManyRequests,gin.H{"Error":"Too many requests"})
			c.String(http.StatusTooManyRequests,"Too many requests")
			c.Abort()
			return
		}
		c.Next()
		//关闭这个请求
		defer limiter.ReleaseConn()
	}
}
