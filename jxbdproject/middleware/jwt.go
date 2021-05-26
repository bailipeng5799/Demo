package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jxbdproject/common"
	"jxbdproject/dao"
	"net/http"
	"strings"
)

func JwtAuthMiddleware()gin.HandlerFunc {
	return func(c *gin.Context){
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusOK,gin.H{
				"msg":"请求头中auth为空",
			})
			c.Abort()
			return
		}
		//按照空格分割
		parts := strings.SplitN(authHeader," ",2)
		if !(len(parts)==2 && parts[0] == "Bearer"){
			c.JSON(http.StatusOK,gin.H{
				"msg":"请求头中auth格式有错误",
			})
			c.Abort()
			return
		}
		//正确处理token
		stringtoken,claims,err := common.ParseToken(parts[1])
		if err!=nil || !stringtoken.Valid{
			c.JSON(http.StatusOK,gin.H{
				"msg":"无效的token",
			})
			c.Abort()
			return
		}
		userid := claims.Userid
		fmt.Println(userid)
		user := dao.CheckUserById(userid)
		if user.Id == 0{
			c.JSON(http.StatusOK,gin.H{
				"msg":"用户不存在在",
			})
			c.Abort()
			return
		}
		c.Set("userid",userid)
		c.Next()

	}
}
