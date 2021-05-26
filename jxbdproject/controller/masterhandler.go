package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jxbdproject/common"
	"jxbdproject/dao"
	"jxbdproject/model"
	"log"
	"net/http"
	"os"
)
//管理者查询详情
func  MasterTopicDetail(c *gin.Context){
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(id)
	topic := dao.MasterDetailTopic(id.Id)
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topic":topic})

}
//管理者添加题目
func MasterTopicAdd(c *gin.Context){
	var addtopic model.Topic
	if err := c.BindJSON(&addtopic);err !=nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterAddTopic(addtopic);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	//增加成功后要返回所有的题目信息
	topics,_:= dao.MasterTopics()
	//返回110代表增加成功
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Code":110,"Topics":topics})
}
func MasterTopicPhotoAdd(c *gin.Context){
	//在前端的html页面上  <input type="file" name="f1">
	file,err := c.FormFile("f1")
	if err != nil{
		common.SendErrorResponse(c,http.StatusInternalServerError,gin.H{"Message":"文件添加失败"})
		return
	}
	fmt.Println(file.Filename)
	dst := model.Video_Dir+file.Filename
	if err := c.SaveUploadedFile(file,dst);err != nil{
		common.SendErrorResponse(c,http.StatusInternalServerError,gin.H{"Message":"服务器保存图片失败"})
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Url":dst,"Message":"图片上传成功"})
	fmt.Println(dst)
}


//管理者删除题目
func MasterTopicDelete(c *gin.Context){
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterDeleteTopic(id.Id);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	topics,_ := dao.MasterTopics()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Code":111,"Topics":topics})
}
//根据题目信息查询有此关键词的题目
func MasterTopicCheck(c *gin.Context){
	var topic model.Topic
	if err := c.ShouldBind(&topic);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	topics,err:= dao.MasterCheckTopic(topic.Question)
	if err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}

//按照题目分类查找题目
func MasterTopicCheckByKind(c *gin.Context){
	var kind model.Topic
	if err := c.ShouldBind(&kind);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	topics,err := dao.MasterCheckByKindTopic(kind)
	if err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}

//管理者修改题目
func MasterTopicUpdate(c *gin.Context){
	var topic model.Topic
	if err := c.BindJSON(&topic);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterUpdateTopic(topic);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	topics,_ := dao.MasterTopics()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}

func MasterAllSchool(c *gin.Context){
	allschool,_ := dao.MasterAllSchool()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Allschool":allschool})
}
func MasterSchoolAdd(c *gin.Context){
	var school model.DrvingSchool
	if err := c.BindJSON(&school);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterAddDrvingSchool(school);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
		return
	}
	allschool,_:=dao.MasterAllSchool()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Allschool":allschool})
}
func MasterSchoolDelete(c *gin.Context){
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterDeleteDrvingSchool(id.Id);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
	}
	allschool,_:=dao.MasterAllSchool()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Allschool":allschool})
}

func MasterSchoolUpadate(c *gin.Context){
	var school model.DrvingSchool
	if err := c.ShouldBind(&school);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.MasterUpdateDrvingSchool(school);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err":model.ErrorDBError.Error})
	}

	allschool,_ := dao.MasterAllSchool()
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Allschool":allschool})
}
func MasterSchoolCheck(c *gin.Context){
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	school := dao.MasterCheckDrvingSchool(id.Id)
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"School":school})
}
func MasterUploadVideos(c *gin.Context){
	var video model.Video
	if err := c.ShouldBind(&video.PracticeName);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	file,err := c.FormFile("videofile")
	if err != nil{
		log.Printf("Error when try to open file: %v",err)
		common.SendErrorResponse(c,model.ErrorInternalFaults.HttpSC,gin.H{"Error":model.ErrorInternalFaults.Error})
	}
	fmt.Println(file.Filename)
	video.VideoName=file.Filename

	dst := model.Video_Dir+file.Filename
	if err := c.SaveUploadedFile(file,dst);err != nil{
		log.Printf("Video file save error :%v",err)
		common.SendErrorResponse(c,model.ErrorInternalFaults.HttpSC,gin.H{"Error":model.ErrorInternalFaults.Error})
		return
	}
	err = dao.MasterAddVideo(video)
	if err != nil{
		log.Printf("db err: %v",err)
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Error":model.ErrorDBError})
		return
	}
	//如果保存文件正常返回文件路径将请求头设置为201
	common.SendNormalResPonse(c,http.StatusCreated,gin.H{"message":"upload successfully!","url":dst})
}
//删除视频
func MasterDeleteVideo(c *gin.Context){
	var video model.Video
	if err := c.ShouldBind(&video.PracticeName);err != nil{
		log.Printf("PracticeName for failure err:%v",err)
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	//在数据库中查询videoName
	video.VideoName = dao.CheckVideoNameByPracticeName(video.PracticeName)
	//如果删除失败
	if !dao.DeleteVideoByPracticeName(video.PracticeName){
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Error":model.ErrorDBError})
		return
	}
	//remove函数当文件不存在的时候也会报错
	//所以IsNotExist会判断这个错误是不是因为文件不存在引起的
	//需要不是这个错误引起才可以执行
	err :=os.Remove(model.Video_Dir+video.VideoName)
	if err != nil && !os.IsNotExist(err){
		log.Printf("Deleting video error: %v",err)
		common.SendErrorResponse(c,model.ErrorInternalFaults.HttpSC,gin.H{"Error":model.ErrorInternalFaults.Error})
		return
	}
	//删除成功
	common.SendNormalResPonse(c,http.StatusCreated,gin.H{"message":"Delete successfully!"})
}
