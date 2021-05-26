package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jxbdproject/common"
	"jxbdproject/dao"
	"jxbdproject/model"
	"net/http"
)
func UserLogin(c *gin.Context) {
	//获取需要输入的账号和密码
	var temp model.User
	if err := c.ShouldBind(&temp); err != nil {
		common.SendErrorResponse(c, model.ErrorRequestBodyParseFailed.HttpSC, gin.H{"Error": model.ErrorRequestBodyParseFailed.Error})
		return
	}
	user, _ := dao.CheckUserNameAndPassword(temp.Username, temp.Password)
	//查询账号密码是否正确
	if user.Id > 0 {
		topics, _ := dao.MasterTopics()
		mytoken, _ := common.Createmv(user)
		common.SendNormalResPonse(c, http.StatusOK, gin.H{"code": 100, "topics": topics, "token": mytoken, "userid": user.Id})

	} else {
		//代表账号密码错误
		c.Header("Content-Type", "application/json")
		//返回101代表账号或者密码错误
		common.SendErrorResponse(c, model.ErrorNotAuthUser.HttpSC, gin.H{"code": 101, "Err": model.ErrorNotAuthUser.Error})
		//发送code002代表用户验证失败
	}
}
func UserRegist(c *gin.Context){
	var temp model.User
	if err := c.ShouldBind(&temp);err!=nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	//首先查询此账号是否被注册过
	fmt.Println(temp)
	user := dao.CheckUser(temp.Username)
	if user.Id > 0{
		//返回一个code103代表此账号已经被注册，不能再次被使用
		//Errorcode 004
		common.SendErrorResponse(c,model.ErrorRegistFailed.HttpSC,gin.H{"Error":model.ErrorRegistFailed.Error})
	}else{
		//此用户名可以使用将temp中的所有信息添加到数据库中
		status := dao.AddUser(&temp)
		//status为bool类型代表是否注册成功
		if !status{
			common.SendErrorResponse(c,model.ErrorRegistFailed.HttpSC,gin.H{"Error":model.ErrorRegistFailed.Error})
		}else{
			user = dao.CheckUser(temp.Username)
			mytoken,_ :=common.Createmv(user)
			common.SendNormalResPonse(c,http.StatusCreated,gin.H{"code":105,"token":mytoken})
		}
	}
}
