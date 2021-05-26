package common

import (
	"github.com/gin-gonic/gin"
)
func SendErrorResponse(c *gin.Context,httpstatus int,data gin.H){
	c.JSON(httpstatus,data)
}
func SendNormalResPonse(c *gin.Context,httpstatus int,data gin.H){
	c.JSON(httpstatus,data)
}